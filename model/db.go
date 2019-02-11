package model

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"log"
)

var orm *xorm.Engine

func init() {

	orm = conn()
}

func conn() *xorm.Engine {

	engine, err := xorm.NewEngine("mysql", "root:housirvip@/spider?charset=utf8")
	if err != nil {
		log.Fatal(err)
	}

	return engine
}

func Orm() *xorm.Engine {

	if orm == nil || orm.DB().Ping() != nil {
		orm = conn()
	}

	return orm
}
