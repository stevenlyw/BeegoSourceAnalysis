package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type Model struct {
	beego.Controller
}

type Score struct {
	Id int
	Score string
}

func init()  {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default","mysql","root:@tcp(127.0.0.1:3306)/test?charset=utf8mb4",30)
	orm.RegisterModel(new(Score))
}

func (c *Model) Get() {
	o :=orm.NewOrm()
	// read one
	score := Score{Id: 1}
	err := o.Read(&score)
	if err !=nil {
		fmt.Printf("ERR: %v\n", err)
	}
	fmt.Println(score)
	c.Data["json"] = map[string]interface{}{"success": 0, "message": "111","data":score}
	c.ServeJSON()
}