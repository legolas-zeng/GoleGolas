//Author: zwa
//Date: 2020/7/9

package main

import (
	"database/sql"
	"encoding/json"
	"errors"
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

func init() {
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
}

type Db interface {
	Query() (map[int]map[string]string, error)
	Update() error
	Insert() error
	Delete() error
}

type Oracle struct {
	Username, Password, Host, Port, Db, Sql string
}

func (m Oracle) Query() (map[int]map[string]string, error) {
	//通过数据库连接池db
	rows, err := db.Query(m.Sql)
	if err != nil {
		return nil, errors.New("\n查询" + m.Sql + "失败,原因:\n" + err.Error())
	}
	defer rows.Close()
	//字典类型
	//构造scanArgs、values两个数组，scanArgs的每个值指向values相应值的地址
	columns, _ := rows.Columns()
	scanArgs := make([]interface{}, len(columns))
	values := make([]sql.RawBytes, len(columns))
	for i := range values {
		scanArgs[i] = &values[i]
	}

	//最后得到的map
	results := make(map[int]map[string]string)
	i := 0
	for rows.Next() {
		//将行数据保存到record字典
		err = rows.Scan(scanArgs...)
		if err != nil {
			return nil, errors.New("结果组装失败,原因:\n" + err.Error())
		}
		row := make(map[string]string)
		for k, v := range values {
			key := columns[k]
			row[key] = string(v)
		}
		results[i] = row
		i++

	}
	mjson, _ := json.Marshal(results)
	mString := string(mjson)
	fmt.Println(mString)
	return results, nil
}

func main() {
	m1 := Oracle{Sql: "SELECT * FROM CREDIT.ACCOUNT_CANCEL WHERE ROWNUM < 10"}
	fmt.Println(m1.Query())
}
