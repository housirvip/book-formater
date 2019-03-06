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
		s := c.Content
		re1, err := regexp.Compile(`[ ]{3,}`)
		if err != nil {
			log.Println(err)
		}
		s = re1.ReplaceAllString(s, "\n")

		re2, err := regexp.Compile(`^第[零一二三四五六七八九十]+章.+\n`)
		if err != nil {
			log.Println(err)
		}
		s = re2.ReplaceAllString(s, "")

		re3, err := regexp.Compile(`http.?://[a-zA-Z0-9./]+`)
		if err != nil {
			log.Println(err)
		}
		s = re3.ReplaceAllString(s, "")

		c.Content = s

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
