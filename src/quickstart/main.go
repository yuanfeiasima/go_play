package main

import (
	_ "quickstart/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"quickstart/models"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	orm.RegisterDataBase("default", "mysql", "root:123456@/lulu?charset=utf8", 30)
	orm.RegisterModel(new(models.User))
	//orm.RunSyncdb("default", false, true)
}

func main() {
	beego.Run()
}

