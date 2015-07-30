package main

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Id   int
	Name string `or:"size(100)"`
}

func init() {
	orm.RegisterDataBase("default", "mysql", "root:linliang@/baidu_top?charset=utf8", 30)
	orm.RegisterModel(new(User))
	orm.RunSyncdb("default", false, true)
}

func main() {
	orm.Debug = true
	o := orm.NewOrm()

	user := User{Name: "slene"}

	//insert
	id, err := o.Insert(&user)
	fmt.Printf("ID: %d, ERR: %v\n", id, err)

	// update
	user.Name = "astaxie"
	num, err := o.Update(&user)
	fmt.Printf("NUM: %d, ERR: %v\n", num, err)

	//read
	u := User{Id: user.Id}
	err = o.Read(&u)
	fmt.Printf("ERR: %v\n", err)

	//delete
	num, err = o.Delete(&u)
	fmt.Printf("ID: %d, ERR: %v\n", num, err)

}
