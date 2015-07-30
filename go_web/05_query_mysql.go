package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
func main() {

	db, err := sql.Open("mysql", "develop:Aa123456@tcp(db-dev.cd.tclclouds.com:3308)/test?charset=utf8")
	checkErr(err)
	// 插入数据
	stmt, err := db.Prepare("INSERT userinfo SET username=?,departname=?,created=?")
	checkErr(err)

	res, err := stmt.Exec("astaxie", "研发部门", "2012-12-09")
	id, err := res.LastInsertId()
	checkErr(err)
	fmt.Println(id)

	// 更新数据
	stmt, err = db.Prepare("UPDATE userinfo set username=? where uid=?")
	checkErr(err)
	res, err = stmt.Exec("crazygit", id)
	checkErr(err)
	affect, err := res.RowsAffected()
	checkErr(err)
	fmt.Println(affect)

	// 查询数据
	rows, err := db.Query("SELECT * FROM userinfo")
	checkErr(err)
	for rows.Next() {
		var uid int
		var username string
		var department string
		var created string
		err = rows.Scan(&uid, &username, &department, &created)
		checkErr(err)
		fmt.Println(uid)
		fmt.Println(username)
		fmt.Println(department)
		fmt.Println(created)
	}

	// 删除数据
	stmt, err = db.Prepare("DELETE FROM userinfo where uid=?")
	checkErr(err)
	res, err = stmt.Exec(1)
	checkErr(err)
	affect, err = res.RowsAffected()
	checkErr(err)
	fmt.Println(affect)
	db.Close()
}
