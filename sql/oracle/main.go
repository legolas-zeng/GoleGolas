//Author: zwa
//Date: 2020/6/30

package main

import (
	"database/sql"
	"fmt"
	_ "github.com/godror/godror"
)

var db *sql.DB

//数据库配置
const (
	host        = "192.168.3.203"
	port        = 1521
	user        = "hrstatic"
	sqlpassword = "efgawetJG339va"
	dbname      = "ceszro"
)

func initsql() {
	// 用户名/密码@IP:端口/实例名
	osqlInfo := fmt.Sprintf("%s/%s@%s:%d/%s", user, sqlpassword, host, port, dbname)
	DB, err := sql.Open("godror", osqlInfo)
	if err != nil {
		panic(err)
	}
	err = DB.Ping()
	if err != nil {
		panic(err)
	}
	db = DB
	rows, err := db.Query("SELECT * FROM CREDIT.ACCOUNT_CANCEL WHERE ROWNUM < 10")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(rows)

	//rows.Close()

}

func main() {
	initsql()
	defer db.Close()
}
