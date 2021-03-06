package handler

import (
	"book-formater/model"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"strconv"
)

func WriteFile(file, content string) {

	fd, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	defer fd.Close()
	if err != nil {
		log.Fatal(err)
	}

	buf := []byte(content)
	_, err = fd.Write(buf)
	if err != nil {
		log.Fatal(err)
	}
}

func ReadFile(file string) []model.Chapter {

	var chapters []model.Chapter

	data, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatalln(err)
	}

	err = json.Unmarshal(data, &chapters)
	if err != nil {
		log.Fatalln(err)
	}

	return chapters
}

func CreateBook(bookId int) {

	book := &model.Book{
		Id: bookId,
	}
	err := book.OneById()
	if err != nil {
		log.Fatal(err)
	}

	chapter := &model.Chapter{
		BookId: bookId,
	}
	chapters, err := chapter.AllByBookId()
	for _, c := range chapters {
		title := "#第" + strconv.Itoa(c.Num) + "章 " + c.Title
		log.Println(title)
		//*[re:test(., "^\s*[第卷][0123456789一二三四五六七八九十零〇百千两]*[章回部节集卷].*", "i")]
		WriteFile("../book/"+book.Name+".txt", title+"\n"+c.Content+"\n")
	}
}
