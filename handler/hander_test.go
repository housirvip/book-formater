package handler

import (
	"book-formater/model"
	"log"
	"regexp"
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

		err = TrimTitle(c)
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

func TestWriteFile(t *testing.T) {

	WriteFile("a", "你好啊臭傻逼")
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
		Name:   "圣祖",
		Author: "傲天无痕",
	}
	err := book.Create()
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("book id:", book.Id)

	chapters := ReadFile("../book/圣祖.json")
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
