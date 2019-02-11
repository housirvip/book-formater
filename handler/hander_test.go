package handler

import (
	"book-formater/model"
	"log"
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
