package controllers

import (
	"github.com/astaxie/beego"
	"fmt"
	"github.com/astaxie/beego/orm"
	"quickstart/models"
	"encoding/json"
	"github.com/astaxie/beego/logs"
)

type UserController struct{
	beego.Controller
}
/*
 db insert
 */
func (c *UserController) UserInsert(){
	o := orm.NewOrm()
	//orm.Debug = true sql
	user := models.User{Name:"slene"}

	o.Begin()//事务
	id, err := o.Insert(&user)
	if err == nil {
		o.Commit()
	} else {
		fmt.Printf("ID: %d, ERR: %v\n", id, err)
		o.Rollback()
	}

	c.Ctx.WriteString("insert ok")
}

/*
db select
 */
func (c *UserController) UserGet() {
	var id int
	c.Ctx.Input.Bind(&id, "id") //获取get形式的参数
	o := orm.NewOrm()
	u := models.User{Id: id}
	err := o.Read(&u)
	fmt.Printf("ERR:%v\n", err)
	fmt.Println(u.Name)
	c.Ctx.WriteString(u.Name)
}

/*
db update
 */
func (c *UserController) Update() {
	var name string
	var ob models.User
	json.Unmarshal(c.Ctx.Input.RequestBody, &ob) //post获取参数方式 注意配置支持
	logs.Info("参数:")
	logs.Info(ob)
	name = ob.Name
	logs.Info("name" + name)
	o := orm.NewOrm()
	user := models.User{Id:2, Name:name}
	o.Begin()
	num, err := o.Update(&user)
	if err == nil{
		fmt.Printf("update, affected num :%d", num)
		o.Commit()
	} else {
		fmt.Printf("update, ERR: %v\n", err)
		o.Rollback()
	}
	c.Ctx.WriteString("update ok")
}