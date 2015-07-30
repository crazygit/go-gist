package main

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	. "quickstart/models"
)

func init() {
	orm.RegisterDriver("mysql", orm.DR_MySQL)
	orm.RegisterDataBase("default", "mysql", "root:linliang@/beego?charset=utf8")
	orm.RunSyncdb("default", false, true)
}

func main() {
	orm.Debug = true
	o := orm.NewOrm()
	o.Using("default") // 默认使用 default，你可以指定为其他数据库

	profile := new(Profile)
	profile.Age = 30

	user := new(User)
	user.Profile = profile
	user.Name = "slene"

	fmt.Println(o.Insert(profile))
	fmt.Println(o.Insert(user))
}
