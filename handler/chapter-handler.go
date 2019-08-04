package handler

import (
	"book-formater/helper"
	"book-formater/model"
	"errors"
	"log"
	"regexp"
	"strconv"
	"strings"
)

func TrimContent(c *model.Chapter) error {

	s := c.Content

	re1, err := regexp.Compile(`[a-zA-Z0-9./]+.*[a-zA-Z0-9./]+`)
	if err != nil {
		return err
	}
	s = re1.ReplaceAllString(s, "")

	re2, err := regexp.Compile(`^第[零一二三四五六七八九十百千0-9]+章.+\n`)
	if err != nil {
		return err
	}
	s = re2.ReplaceAllString(s, "")

	c.Content = s

	return nil
}

func TrimTitle(c *model.Chapter) error {

	src := c.Title
	if src == "" {
		return errors.New("title src 为空")
	}
	re, err := regexp.Compile(`第[零一二三四五六七八九十百千0-9]+章[ \s]?`)
	if err != nil {
		log.Fatalln(err)
	}
	c.Title = re.ReplaceAllString(src, "")

	return nil
}

func TrimNumAndTitle(c *model.Chapter) error {

	src := c.Title
	if src == "" {
		return errors.New("title src 为空")
	}
	re, err := regexp.Compile(`第[零一二三四五六七八九十百千0-9]+章[ \s]?`)
	if err != nil {
		log.Fatalln(err)
	}
	c.Title = re.ReplaceAllString(src, "")

	reMixNum, err := regexp.Compile(`[零一二三四五六七八九十百千0-9]+`)
	if err != nil {
		log.Fatalln(err)
	}
	num := reMixNum.FindString(src)
	if num == "" {
		return errors.New("num 为空")
	}

	reNum, err := regexp.Compile(`^[0-9]+$`)
	if err != nil {
		log.Fatalln(err)
	}
	if reNum.MatchString(num) {
		c.Num = NumStrToInt(num)
		return nil
	}

	reHanNum, err := regexp.Compile(`^[零一二三四五六七八九]+$`)
	if err != nil {
		log.Fatalln(err)
	}
	if reHanNum.MatchString(num) {
		hanNum := "零一二三四五六七八九"
		for _, v := range num {
			index := strings.IndexRune(hanNum, v)
			//log.Println(num, index, v, string(v))
			num = strings.Replace(num, string(v), strconv.Itoa(index/3), -1)
		}
		c.Num = NumStrToInt(num)
		return nil
	}

	numInt64, err := helper.DecodeToInt64(num)
	if err != nil {
		return err
	}
	c.Num = int(numInt64)

	return nil
}

func NumStrToInt(str string) int {
	res, err := strconv.Atoi(str)
	if err != nil {
		log.Println(err)
	}
	return res
}
