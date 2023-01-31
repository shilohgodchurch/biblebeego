package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/mattn/go-sqlite3"
	models "biblebeego/models"
	_ "biblebeego/routers"
)

func init() {
	// This is a dummy change to test Hound
	var Go_Ing_Go int
	orm.RegisterDriver("sqlite", orm.DR_Sqlite)
	orm.RegisterDataBase("default", "sqlite3", "database/orm_test.db")
	orm.RegisterModel(new(models.Article))
}

func main() {
	beego.Run()
}
