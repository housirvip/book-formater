package model

import "errors"

type Chapter struct {
	Id      int `xorm:"pk autoincr"`
	Title   string
	Content string `xorm:"text"`
	BookId  int
	Num     int
}

func (c *Chapter) OneById() error {

	ok, err := orm.ID(c.Id).Get(c)
	if err != nil {
		return err
	}
	if !ok {
		return errors.New("error: [Chapter]-OneById")
	}
	return nil
}

func (c *Chapter) AllByBookId() ([]Chapter, error) {

	var cs []Chapter
	err := orm.Where("book_id = ?", c.BookId).Find(&cs)
	if err != nil {
		return nil, err
	}
	return cs, nil
}

func (c *Chapter) UpdateCols(cols ...string) error {

	ok, err := orm.Id(c.Id).Cols(cols...).Update(c)
	if err != nil {
		return err
	}
	if ok == 0 {
		return errors.New("error: [Chapter]-UpdateCols")
	}
	return nil
}
