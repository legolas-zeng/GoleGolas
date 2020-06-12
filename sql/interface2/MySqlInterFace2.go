package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"strconv"
	"strings"
)

var (
	key        string
	value      string
	conditions string
	str        string
)

type Model struct {
	link      *sql.DB  //存储连接对象
	tableName string   //存储表名
	field     string   //存储字段
	allFields []string //存储当前表所有字段
	where     string   //存储where条件
	order     string   //存储order条件
	limit     string   //存储limit条件
}

//构造方法
func NewModel(table string) Model {
	var this Model
	this.field = "*"
	//1.存储操作的表名
	this.tableName = table
	//2.初始化连接数据库
	this.getConnect()
	//3.获得当前表的所有字段
	this.getFields()
	return this
}

/**
 * 初始化连接数据库操作
 */
func (this *Model) getConnect() {
	//1.连接数据库
	db, err := sql.Open("mysql", "root:qq1005521@tcp(127.0.0.1:3306)/movie?charset=utf8")
	//2.判断连接
	if err != nil {
		fmt.Printf("connect mysql fail ! [%s]", err)
	}
	this.link = db
}

/**
 * 获取当前表的所有字段
 */
func (this *Model) getFields() {

	//查看表结构
	sql := "DESC " + this.tableName
	//执行并发送SQL
	result, err := this.link.Query(sql)

	if err != nil {
		fmt.Printf("sql fail ! [%s]", err)
	}

	this.allFields = make([]string, 0)

	for result.Next() {
		var field string
		var Type interface{}
		var Null string
		var Key string
		var Default interface{}
		var Extra string
		err := result.Scan(&field, &Type, &Null, &Key, &Default, &Extra)
		if err != nil {
			fmt.Printf("scan fail ! [%s]", err)
		}
		this.allFields = append(this.allFields, field)
	}

}

/**
 * 执行并发送SQL(查询)
 * @param string $sql  要查询的SQL语句
 * @return array 返回查询出来的二维数组
 */
func (this *Model) query(sql string) interface{} {

	rows2, err := this.link.Query(sql)

	if err != nil {
		return returnRes(0, ``, err)
	}
	//查询数据，取所有字段
	if err != nil {
		return returnRes(0, ``, err)
	}

	//返回所有列
	cols, err := rows2.Columns()

	if err != nil {
		return returnRes(0, ``, err)
	}

	//这里表示一行所有列的值，用[]byte表示
	vals := make([][]byte, len(cols))

	//这里表示一行填充数据
	scans := make([]interface{}, len(cols))

	//这里scans引用vals，把数据填充到[]byte里
	for k, _ := range vals {
		scans[k] = &vals[k]
	}

	i := 0
	result := make(map[int]map[string]string)

	for rows2.Next() {
		//填充数据
		rows2.Scan(scans...) //将slic地址传入
		//每行数据
		row := make(map[string]string)
		//把vals中的数据复制到row中
		for k, v := range vals {
			key := cols[k]
			//这里把[]byte数据转成string
			row[key] = string(v)
		}
		//放入结果集
		result[i] = row
		i++
	}
	return returnRes(1, result, "success")

}

/**
 * 查询多条数据
 */
func (this *Model) get() interface{} {
	sql := `select ` + this.field + ` from ` + this.tableName + ` ` + this.where + ` ` + this.order + ` ` + this.limit
	//执行并发送SQL
	result := this.query(sql)
	return result
}

/**
 * 查询一条数据
 * @param string $id 要查询的id
 * @return array  返回一条数据
 */
func (this *Model) find(user_id int) interface{} {
	//判断id是否存在
	where := ` where user_id = ` + strconv.Itoa(user_id)
	sql := `select ` + this.field + ` from ` + this.tableName + where + ` limit 1`
	//执行并发送sql
	result := this.query(sql)
	return result
}

/**
 * 设置要查询的字段信息
 * @param string $field  要查询的字段
 * @return object 返回自己，保证连贯操作
 */
func (this *Model) Field(field string) *Model {
	this.field = field
	return this
}

/**
 * order排序条件
 * @param string  $order  以此为基准进行排序
 * @return $this  返回自己，保证连贯操作
 */
func (this *Model) Order(order string) *Model {
	this.order = `order by ` + order
	return this
}

/**
 * limit条件
 * @param string $limit 输入的limit条件
 * @return $this 返回自己，保证连贯操作
 */
func (this *Model) Limit(limit int) *Model {
	this.limit = "limit " + strconv.Itoa(limit)
	return this
}

/**
 * where条件
 * @param string $where 输入的where条件
 * @return $this 返回自己，保证连贯操作
 */
func (this *Model) Where(where string) *Model {
	this.where = `where ` + where
	return this
}

/**
 * 统计总条数
 * @return int 返回总数
 */
func (this *Model) count() interface{} {
	//准备SQL语句
	sql := `select count(*) as total from ` + this.tableName + ` limit 1`
	result := this.query(sql)
	return returnRes(1, result, "success")
}

/**
 * 执行并发送SQL语句(增删改)
 * @param string $sql 要执行的SQL语句
 * @return bool|int|string 添加成功则返回上一次操作id,删除修改操作则返回true,失败则返回false
 */
func (this *Model) exec(sql string) interface{} {

	res, err := this.link.Exec(sql)

	if err != nil {
		return returnRes(0, ``, err)
	}

	result, err := res.LastInsertId()
	if err != nil {
		return returnRes(0, ``, err)
	}
	return returnRes(1, result, "success")
}

/**
 * 添加操作
 * @param array  $data 要添加的数组
 * @return bool|int|string 添加成功则返回上一次操作的id,失败则返回false
 */
func (this *Model) add(data map[string]interface{}) interface{} {

	//过滤非法字段
	for k, v := range data {
		if res := in_array(k, this.allFields); res != true {
			delete(data, k)
		} else {
			key += `,` + k
			value += `,` + `'` + v.(string) + `'`
		}
	}

	//将map中取出的键转为字符串拼接
	key = strings.TrimLeft(key, ",")
	//将map中的值转化为字符串拼接
	value = strings.TrimLeft(value, ",")
	//准备SQL语句
	sql := `insert into ` + this.tableName + ` (` + key + `) values (` + value + `)`
	// //执行并发送SQL
	result := this.exec(sql)

	return result

}

/**
 * 删除操作
 * @param string $id 要删除的id
 * @return bool  删除成功则返回true,失败则返回false
 */
func (this *Model) delete(user_id int) interface{} {

	//判断id是否存在
	if this.where == "" {
		conditions = `where user_id = ` + strconv.Itoa(user_id)
	} else {
		conditions = this.where + ` and user_id = ` + strconv.Itoa(user_id)
	}

	sql := `delete from ` + this.tableName + ` ` + conditions

	//执行并发送
	result := this.exec(sql)

	return result
}

/**
 * 修改操作
 * @param  array $data  要修改的数组
 * @return bool 修改成功返回true，失败返回false
 */
func (this *Model) update(data map[string]interface{}) interface{} {

	//过滤非法字段
	for k, v := range data {
		if res := in_array(k, this.allFields); res != true {
			delete(data, k)
		} else {
			str += k + ` = '` + v.(string) + `',`
		}
	}

	//去掉最右侧的逗号
	str = strings.TrimRight(str, ",")

	//判断是否有条件
	if this.where == "" {
		fmt.Println("没有条件")
	}

	sql := `update ` + this.tableName + ` set ` + str + ` ` + this.where

	result := this.exec(sql)
	return result
}

//是否存在数组内
func in_array(need interface{}, needArr []string) bool {
	for _, v := range needArr {
		if need == v {
			return true
		}
	}
	return false
}

//返回json
func returnRes(errCode int, res interface{}, msg interface{}) string {
	result := make(map[string]interface{})
	result["errCode"] = errCode
	result["result"] = res
	result["msg"] = msg
	data, _ := json.Marshal(result)
	return string(data)
}

func main() {

	M := NewModel("jiulianshan2")
	//查询链式操作
	res := M.Field("riqi,isbn").Order("book desc").Where("riqi = 2012").Limit(10).get()

	//添加操作
	// data := make(map[string]interface{})
	// data["ddd"] = "118284901@qq.com"
	// data["daaa"] = "118284901@qq.com"
	// data["email"] = "118284901@qq.com"
	// data["user_name"] = "张三"
	// data["add_time"] = time.Unix(time.Now().Unix(),0).Format("2006-01-02 15:04:05")
	// res:=M.add(data)

	//删除操作
	// res :=M.delete(18)

	//更新操作
	// data := make(map[string]interface{})
	// data["email"] = "118284901@qq.com"
	// data["user_name"] = "打啊"
	// data["add_time"] = time.Unix(time.Now().Unix(),0).Format("2006-01-02 15:04:05")
	//res:=M.Where("user_id = 1").update(data)
	fmt.Println(res)

}
