package handler

import (
	"book-formater/model"
	"log"
	"regexp"
)

func TrimContent() {

	orm := model.Orm()
	err := orm.Where("id > ?", 0).Iterate(new(model.Chapter), func(i int, bean interface{}) error {

		c := bean.(*model.Chapter)
		re, err := regexp.Compile(`^第[零一二三四五六七八九十]+章.+\n`)
		if err != nil {
			log.Println(err)
		}
		c.Content = re.ReplaceAllString(c.Content, "")

		err = c.UpdateCols("content")
		if err != nil {
			log.Println(err)
		}

		log.Println(c.Id, c.Title)

		return nil
	})
	if err != nil {
		log.Println(err)
	}
}
