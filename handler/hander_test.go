package handler

import (
	"book-formater/model"
	"log"
	"regexp"
	"testing"
)

func TestTrimContent(t *testing.T) {

	TrimContent()
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

	chapters := ReadFile("../book/shengzu.json")
	for k, v := range chapters {
		if len(v.Content) < 50 {
			log.Println(k, v.Title, v.Content)
		} else {
			log.Println(k, v.Title)
		}
	}
}
