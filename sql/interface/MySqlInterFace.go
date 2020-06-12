package main

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"
	// "fmt"
	_ "github.com/go-sql-driver/mysql"
)

//数据库连接DSN
const (
	DbType = "mysql"
	// DefaultConn = "root:system@tcp(172.16.4.51:3306)/monitor?charset=utf8"
)

//mysql动态连接池
var db *sql.DB
var DefaultConn string

//连接数据库基本数据字典
type Info struct {
	username string
	password string
	host     string
	port     string
	dbname   string
	charset  string
}

func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}

func init() {
	//获取配置文件
	info := Info{}
	info.username = "root"
	info.password = "qq1005521"
	info.host = "localhost"
	info.port = "3306"
	info.dbname = "movie"
	info.charset = "utf8"
	//组装MYSQL连接DSN连接
	DefaultConn = info.username + ":" + info.password + "@tcp(" + info.host + ":" + info.port + ")/" + info.dbname + "?charset=" + info.charset
	// "初始化数据库连接池"
	db, _ = sql.Open(DbType, DefaultConn)
	db.SetMaxOpenConns(2000)
	db.SetMaxIdleConns(1000)
	db.Ping()
}

//定义一个接口，只要实现了这个接口的方法就可以调用这个接口
type Db interface {
	//http://127.0.0.1/mysql/query?sql=select host,type,table_schema,table_name,table_rows,table_size_m,index_size_m,sysdates from monitorss order by index_size_m desc limit 100
	Query() (map[int]map[string]string, error)
	//http://172.16.4.51/mysql/update?table=tttt&set=name:lxp&where=id:1
	Update() error
	// http://127.0.0.1/mysql/insert?table=ttt&data=name:lixueping,id:1
	Insert() error
	// http://127.0.0.1/mysql/delete?table=ttt&where=id:1
	Delete() error
}

//定义mysql结构，条件全部满足后就可以实现接口方法
type Mysql struct {
	Sql string
	//table->map[string]string  data->map[string]map[string]string
	Data map[string]interface{}
}

type Oracle struct {
	Username, Password, Host, Port, Db, Sql string
}

type Mongodb struct {
	Username, Password, Host, Port, Db, Sql string
}

//实现mysql的db接口
//"root:password@tcp(127.0.0.1:3306)/monitor?charset=utf8"
func (m Mysql) Query() (map[int]map[string]string, error) {
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
	return results, nil
}

func (m Mysql) Insert() error {
	//m.Data由table和data构成
	//初始化一个结果集
	args := make([]interface{}, len(m.Data["data"].(map[string]string)))
	hosts := "("
	//统计次数
	count := 0
	//组装结果集和sql
	for k, v := range m.Data["data"].(map[string]string) {
		args[count] = v
		count++
		if count != len(m.Data["data"].(map[string]string)) {
			hosts += k + ","
		} else {
			hosts += k + ")"
		}

	}
	sqlScript := "INSERT INTO " + m.Data["table"].(string) + " " + hosts + " VALUES (?" + strings.Repeat(",?", len(m.Data["data"].(map[string]string))-1) + ")"
	// println(sqlScript)
	/*
	   println(sqlScript, args)
	   _, err := db.Query(sqlScript, args...)
	   由于自带的insert比如写死参数，无法动态调配，所以换成组装SQL的形式，以下代码留着学习
	*/

	stmt, err := db.Prepare(sqlScript)
	defer stmt.Close()
	if err != nil {
		return errors.New(sqlScript + " 字段不存在或者数据不正确！")
	}
	_, err = stmt.Exec(args...)
	CheckErr(err)

	return err
}

func (m Mysql) Update() error {
	//m.Data由table、set和where构成
	//初始化一个结果集
	args := make([]interface{}, len(m.Data["set"].(map[string]string))+len(m.Data["where"].(map[string]string)))
	//统计次数
	count := 0
	c_where := 0
	//组装结果集set和where
	set := " SET "
	where := " WHERE "
	for k, v := range m.Data["set"].(map[string]string) {
		args[count] = v
		count++
		if count != len(m.Data["set"].(map[string]string)) {
			set += k + "=?, "
		} else {
			set += k + "=? "
		}
	}

	for kk, vv := range m.Data["where"].(map[string]string) {
		args[count] = vv
		count++
		c_where++
		if len(m.Data["where"].(map[string]string)) == 1 {
			where += kk + "=?"
		} else if c_where != len(m.Data["where"].(map[string]string)) && len(m.Data["where"].(map[string]string)) != 1 {
			where += kk + "=? and "
		} else if c_where == len(m.Data["where"].(map[string]string)) {
			where += kk + "=? "
		}
	}

	sqlScript := "UPDATE " + m.Data["table"].(string) + set + where
	// println(sqlScript)
	/*
	   println(sqlScript, args)
	   _, err := db.Query(sqlScript, args...)
	   由于自带的insert比如写死参数，无法动态调配，所以换成组装SQL的形式，以下代码留着学习
	*/

	stmt, err := db.Prepare(sqlScript)
	defer stmt.Close()
	if err != nil {
		return errors.New(sqlScript + " 字段不存在或者数据不正确！")
	}
	_, err = stmt.Exec(args...)
	CheckErr(err)

	return err
}

func (m Mysql) Delete() error {
	//m.Data由table和where构成
	//初始化一个结果集
	args := make([]interface{}, len(m.Data["where"].(map[string]string)))
	//统计次数
	count := 0
	//组装结果集set和where
	where := " WHERE "

	for kk, vv := range m.Data["where"].(map[string]string) {
		args[count] = vv
		count++
		if len(m.Data["where"].(map[string]string)) == 1 {
			where += kk + "=?"
		} else if count != len(m.Data["where"].(map[string]string)) && len(m.Data["where"].(map[string]string)) != 1 {
			where += kk + "=? and "
		} else if count == len(m.Data["where"].(map[string]string)) {
			where += kk + "=? "
		}
	}

	sqlScript := "DELETE FROM " + m.Data["table"].(string) + where
	// println(sqlScript)
	/*
	   println(sqlScript, args)
	   _, err := db.Query(sqlScript, args...)
	   由于自带的insert比如写死参数，无法动态调配，所以换成组装SQL的形式，以下代码留着学习
	*/

	stmt, err := db.Prepare(sqlScript)
	defer stmt.Close()
	if err != nil {
		return errors.New(sqlScript + " 字段不存在或者数据不正确！")
	}
	_, err = stmt.Exec(args...)
	CheckErr(err)

	return err

}

func main() {
	m1 := Mysql{Sql: "select * from jiulianshan2 limit 10;"}
	fmt.Println(m1.Query())
	m2 := Mysql{
		Sql:  nil,
		Data: Info{},
	}
	m2.Insert()

}
