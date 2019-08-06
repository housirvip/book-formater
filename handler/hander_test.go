package handler

import (
	"book-formater/model"
	"log"
	"regexp"
	"strconv"
	"strings"
	"testing"
)

func TestChapter(t *testing.T) {

	orm := model.Orm()
	err := orm.Where("book_id = ?", 1).Iterate(new(model.Chapter), func(i int, bean interface{}) error {

		c := bean.(*model.Chapter)

		err := TrimContent(c)
		if err != nil {
			log.Println(err)
		}

		err = TrimNumAndTitle(c)
		if err != nil {
			log.Println(err)
		}

		err = c.UpdateCols("title", "content", "num")
		if err != nil {
			log.Println(err)
		}

		log.Println(c.Id, c.Title, c.Num)

		return nil
	})
	if err != nil {
		log.Println(err)
	}
}

func TestTrimNumAndTitleTitle(t *testing.T) {

	orm := model.Orm()
	err := orm.Where("book_id = ?", 1).Iterate(new(model.Chapter), func(i int, bean interface{}) error {

		c := bean.(*model.Chapter)
		err := TrimNumAndTitle(c)
		if err != nil {
			log.Println(err)
		}

		err = c.UpdateCols("title", "num")
		if err != nil {
			log.Println(err)
		}

		log.Println(c.Num, c.Title)

		return nil
	})
	if err != nil {
		log.Println(err)
	}
}

func TestTrimTitle(t *testing.T) {

	orm := model.Orm()
	err := orm.Where("book_id = ?", 1).Iterate(new(model.Chapter), func(i int, bean interface{}) error {

		c := bean.(*model.Chapter)
		err := TrimTitle(c)
		if err != nil {
			log.Println(err)
		}
		log.Println(c.Title)

		return nil
	})
	if err != nil {
		log.Println(err)
	}
}

func TestTrimContentTemp(t *testing.T) {

	orm := model.Orm()
	err := orm.Where("book_id = ?", 1).Iterate(new(model.Chapter), func(i int, bean interface{}) error {

		c := bean.(*model.Chapter)
		re, err := regexp.Compile(`https://.+.com`)
		if err != nil {
			log.Println(err)
		}
		c.Content = re.ReplaceAllString(c.Content, "")

		err = c.UpdateCols("content")
		if err != nil {
			log.Println(err)
		}
		//log.Println(c.Content)

		return nil
	})
	if err != nil {
		log.Println(err)
	}
}

func TestTrimContentEnter(t *testing.T) {

	orm := model.Orm()
	err := orm.Where("book_id = ?", 1).Iterate(new(model.Chapter), func(i int, bean interface{}) error {

		c := bean.(*model.Chapter)
		re, err := regexp.Compile(`[ ]{3,}`)
		if err != nil {
			log.Println(err)
		}
		c.Content = re.ReplaceAllString(c.Content, "\n")
		log.Println(c.Content)

		return nil
	})
	if err != nil {
		log.Println(err)
	}
}

func TestTrimUrlAndNum(t *testing.T) {

	orm := model.Orm()
	err := orm.Where("book_id = ?", 1).Iterate(new(model.Chapter), func(i int, bean interface{}) error {

		c := bean.(*model.Chapter)

		c.Url = strings.Replace(c.Url, "https://www.yangguiweihuo.com/2/2754/", "", -1)
		c.Url = strings.Replace(c.Url, ".html", "", -1)

		var err error
		c.Num, err = strconv.Atoi(c.Url)
		if err != nil {
			return err
		}
		log.Println(c.Num)

		err = c.UpdateCols("num")
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		log.Println(err)
	}
}

func TestNumCount(t *testing.T) {

	orm := model.Orm()
	total, err := orm.Count(new(model.Chapter))
	if err != nil {
		log.Println(err)
	}

	for i := 0; int64(i) < total; i++ {
		has, _ := orm.Exist(&model.Chapter{Num: i})
		if !has {
			log.Println(i)
		}
	}
}

func TestCreateBook(t *testing.T) {

	CreateBook(1)
}

func TestChapterList(t *testing.T) {

	chapter := &model.Chapter{
		BookId: 1,
	}
	cs, err := chapter.AllByBookId()
	if err != nil {
		log.Fatal(err)
	}
	log.Println(len(cs))
}

func Test1(t *testing.T) {

	s := "        第一五零零章庞家灭门真相          神令再现的事情，罗烈并不知道。          此刻的他也没有远离冰城，而是打招呼通知南宫天王不要再战"
	re1, err := regexp.Compile(`[ ]{3,}`)
	if err != nil {
		log.Println(err)
	}
	s = re1.ReplaceAllString(s, "\n")
	log.Println(s)
}

func TestReadFile(t *testing.T) {

	book := &model.Book{
		Name:   "无疆",
		Author: "小刀锋利",
	}
	err := book.Create()
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("book id:", book.Id)

	chapters := ReadFile("../book/wj.json")
	for k, v := range chapters {
		text := strings.Replace(v.Content, "\r\r", "\n", -1)
		v.Content = strings.Replace(text, "\u00a0", "", -1)
		v.BookId = book.Id
		if err != nil {
			log.Println(err)
		}
		err = v.Create()
		log.Println(k, v.Title)
	}
}
