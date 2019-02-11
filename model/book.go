package model

import "errors"

type Book struct {
	Id     int `xorm:"pk autoincr"`
	Name   string
	Author string
	Cover  string
}

func (b *Book) OneById() error {

	ok, err := orm.ID(b.Id).Get(b)
	if err != nil {
		return err
	}
	if !ok {
		return errors.New("error: [Book]-OneById")
	}
	return nil
}
