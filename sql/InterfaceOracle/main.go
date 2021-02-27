//Author: zwa
//Date: 2020/7/9

package main

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	_ "github.com/godror/godror"
	"reflect"
	"strconv"
	"time"
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
	fmt.Println(reflect.TypeOf(rows))
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

func Tests(Sql string) map[int]map[string]string {
	rows, _ := db.Query(Sql)
	defer rows.Close()
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
		err := rows.Scan(scanArgs...)
		if err != nil {
			fmt.Println(err)
		}
		row := make(map[string]string)
		for k, v := range values {
			key := columns[k]
			row[key] = string(v)
		}
		results[i] = row
		i++
	}
	return results
}

func handledata(data map[int]map[string]string) {
	xlsx := excelize.NewFile()
	index := xlsx.NewSheet("Sheet1")
	xlsx.SetActiveSheet(index)
	//表头
	var rows = []string{"门店名称", "原签约门店", "合同编号", "客户名称", "审批金额", "签约金额", "签约期限", "还款方式", "放款日期", "初笔还歀日", "末笔还款日期", "第一期计划应还", "每月计划应还", "应还日期", "当期计划应还", "剩余未还本金", "已还期数", "是否结清", "当期逾期天数", "逾期天数调整", "逾期特殊调整为C", "合同尾号", "是否为追加贷", "合并合同编号（剔除追加贷）", "最大逾期天数（合并合同编号）", "期初状态（合并合同编号）"}
	//var rows = []string{"APPROVED_AMOUNT","BORROW_MONEY","CONTRACT_DURATION","REPAY_WAY","LEND_DATE" ,"FISRT_PERIOD_REPAY","LAST_PERIOD_REPAY","FIRST_JH_MONEY","MONTH_JH_MONEY","REPAY_DATE","JH_MONEY","SY_REPAY_PRINCIPAL","SETTLE_PERIOD","IS_SETTLE","YQ_DAYS"}
	//xlsx.SetCellValue("Sheet1", cell, )
	xlsx.InsertRow("Sheet1", 3)
	for k, v := range rows {
		//fmt.Println(k,v)
		col := GetLetterByNum(k + 1)
		cell := col + "1"
		xlsx.SetCellValue("Sheet1", cell, v)
	}
	//表内容，行数k
	for k, v := range data {
		fmt.Println("-----", k, v)
		fmt.Printf("%T \n", v["IS_SETTLE"])
		//剔除已结清
		if v["IS_SETTLE"] == "0" {
			a := 0
			for _, j := range []string{"STORE_NAME", "ORIGINAL_STRUCT_NAME", "CONTRACT_NO", "CUSTOMER_NAME", "APPROVED_AMOUNT", "BORROW_MONEY", "CONTRACT_DURATION", "REPAY_WAY", "LEND_DATE", "FISRT_PERIOD_REPAY", "LAST_PERIOD_REPAY", "FIRST_JH_MONEY", "MONTH_JH_MONEY", "REPAY_DATE", "JH_MONEY", "SY_REPAY_PRINCIPAL", "SETTLE_PERIOD", "IS_SETTLE", "YQ_DAYS"} {
				col := GetLetterByNum(a + 1)
				fmt.Println("数据：", col+strconv.Itoa(k+2), v[j])
				xlsx.SetCellValue("Sheet1", col+strconv.Itoa(k+2), v[j])
				a++
			}
			//除0之外的逾期天数调整+1
			if v["YQ_DAYS"] != "0" {
				int, _ := strconv.Atoi(v["YQ_DAYS"])
				xlsx.SetCellValue("Sheet1", "T"+strconv.Itoa(k+2), int+1)
			} else {
				xlsx.SetCellValue("Sheet1", "T"+strconv.Itoa(k+2), 0)
			}
			//获取合同号最后三个字符
			xlsx.SetCellValue("Sheet1", "V"+strconv.Itoa(k+2), v["CONTRACT_NO"][len(v["CONTRACT_NO"])-3:])
			//判断是否为追加贷
			if v["CONTRACT_NO"][len(v["CONTRACT_NO"])-3:] == "ZJ1" {
				xlsx.SetCellValue("Sheet1", "W"+strconv.Itoa(k+2), "是")
				xlsx.SetCellValue("Sheet1", "X"+strconv.Itoa(k+2), v["CONTRACT_NO"])

			} else {
				xlsx.SetCellValue("Sheet1", "W"+strconv.Itoa(k+2), "否")
				xlsx.SetCellValue("Sheet1", "X"+strconv.Itoa(k+2), juegestr(v["CONTRACT_NO"]))
			}

		}
	}

	err := xlsx.SaveAs("C:\\Users\\Administrator.000\\Desktop\\Workbook.xlsx")
	if err != nil {
		fmt.Println(err)
	}
}

func juegestr(str string) string {
	if len(str) <= 18 {
		return str
	} else {
		return str[:18]
	}
}

//二十六进制转换
func GetLetterByNum(num int) (s string) {
	if num <= 0 {
		panic("num参数取值范围为大于零的整数")
	}
	var temp []rune
	yu := 0
	shang := num
	for {
		yu = shang % 26
		shang = shang / 26
		if yu == 0 {
			yu = 26
			shang--
		}
		temp = append(temp, rune(yu+'A'-1))
		if shang == 0 {
			break
		}
	}
	for i := len(temp) - 1; i >= 0; i-- {
		s += string(temp[i])
	}
	return
}

func handledata2(data map[int]map[string]string) {
	xlsx := excelize.NewFile()
	index := xlsx.NewSheet("Sheet1")
	xlsx.SetActiveSheet(index)
	//表头
	var rows = []string{"门店名称", "原签约门店", "合同编号", "客户名称", "审批金额", "签约金额", "签约期限", "还款方式", "放款日期", "初笔还歀日", "末笔还款日期", "第一期计划应还", "每月计划应还", "应还日期", "当期计划应还", "剩余未还本金", "已还期数", "是否结清", "当期逾期天数", "逾期天数调整", "合同尾号", "是否为追加贷", "合并合同编号（剔除追加贷）", "最大逾期天数（合并合同编号）", "期初状态（合并合同编号）", "逾期特殊调整为C"}
	//var rows = []string{"APPROVED_AMOUNT","BORROW_MONEY","CONTRACT_DURATION","REPAY_WAY","LEND_DATE" ,"FISRT_PERIOD_REPAY","LAST_PERIOD_REPAY","FIRST_JH_MONEY","MONTH_JH_MONEY","REPAY_DATE","JH_MONEY","SY_REPAY_PRINCIPAL","SETTLE_PERIOD","IS_SETTLE","YQ_DAYS"}
	//xlsx.SetCellValue("Sheet1", cell, )
	xlsx.InsertRow("Sheet1", 3)
	for k, v := range rows {
		//fmt.Println(k,v)
		col := GetLetterByNum(k + 1)
		cell := col + "1"
		xlsx.SetCellValue("Sheet1", cell, v)
	}
	//表内容，行数k
	for k, v := range data {
		fmt.Println("-----", k, v)
		a := 0
		for _, j := range []string{"STORE_NAME", "ORIGINAL_STRUCT_NAME", "CONTRACT_NO", "CUSTOMER_NAME", "APPROVED_AMOUNT", "BORROW_MONEY", "CONTRACT_DURATION", "REPAY_WAY", "LEND_DATE", "FISRT_PERIOD_REPAY", "LAST_PERIOD_REPAY", "FIRST_JH_MONEY", "MONTH_JH_MONEY", "REPAY_DATE", "JH_MONEY", "SY_REPAY_PRINCIPAL", "SETTLE_PERIOD", "IS_SETTLE", "YQ_DAYS", "YQ_DAYS_TZ", "CONTRACT_NO_WH", "IS_ZJ", "CONTRACT_NO_HB", "MAX_YQ_DAYS", "CE_DE"} {
			col := GetLetterByNum(a + 1)
			fmt.Println("数据：", col+strconv.Itoa(k+2), v[j])
			xlsx.SetCellValue("Sheet1", col+strconv.Itoa(k+2), v[j])
			a++
		}
	}
	err := xlsx.SaveAs("C:\\Users\\Administrator.000\\Desktop\\Workbook.xlsx")
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	t1 := time.Now()
	// WHERE ROWNUM <= 500
	//sql:= "select STORE_NAME,ORIGINAL_STRUCT_NAME,CONTRACT_NO,CUSTOMER_NAME,APPROVED_AMOUNT,BORROW_MONEY,CONTRACT_DURATION,REPAY_WAY,LEND_DATE,FISRT_PERIOD_REPAY,LAST_PERIOD_REPAY,FIRST_JH_MONEY,MONTH_JH_MONEY,REPAY_DATE,JH_MONEY,SY_REPAY_PRINCIPAL,SETTLE_PERIOD,IS_SETTLE,YQ_DAYS from (select rownum rn, t.* from (select * from CREDIT.ACCOUNT_LOAN_INTEREST WHERE store_id  in (select id from CREDIT.company_struct cs where status = '0' and type = '1' start with id = (select  struct_id from CREDIT.user_info where user_id = 'yunwei') connect by p_struct_id = prior id) AND REPORT_DATE_STR='202005'  order by id,repay_day asc) t WHERE ROWNUM <= 500) t1 where t1.rn > 0"
	sql := `
	SELECT
	 x.CONTRACT_NO as CONTRACT_NO_1, x.YQ_DAYS, 
	 max(x.YQ_DAYS) over (partition by x.CONTRACT_NO) as MAX_YQ_DAYS,
	 case when x.YQ_DAYS != 0 THEN (x.YQ_DAYS + 1) ELSE 0  end AS YQ_DAYS_TZ,
	 SUBSTR(x.CONTRACT_NO,-3) AS CONTRACT_NO_WH,
	 CASE WHEN SUBSTR(x.CONTRACT_NO,-3,2) = 'ZJ' THEN 1 ELSE 0 end AS IS_ZJ,
	 CASE WHEN SUBSTR(x.CONTRACT_NO,-3,2) = 'ZJ' THEN x.CONTRACT_NO ELSE  SUBSTR(x.CONTRACT_NO,1,18) end AS CONTRACT_NO_HB,
	 DECODE(CEIL(max(x.YQ_DAYS) over (partition by x.CONTRACT_NO) /30 ) 
		,0,'M0' 
		,1,'M1'
		,2,'M2'
		,3,'M3'
		,4,'M4'
		,5,'M5'
		,6,'M6'
		,'M7' )AS CE_de,
	 x.STORE_NAME,
		x.ORIGINAL_STRUCT_NAME,
		x.CONTRACT_NO,
		x.CUSTOMER_NAME,
		x.APPROVED_AMOUNT,
		x.BORROW_MONEY,
		x.CONTRACT_DURATION,
		x.REPAY_WAY,
		x.LEND_DATE,
		x.FISRT_PERIOD_REPAY,
		x.LAST_PERIOD_REPAY,
		x.FIRST_JH_MONEY,
		x.MONTH_JH_MONEY,
		x.REPAY_DATE,
		x.JH_MONEY,
		x.SY_REPAY_PRINCIPAL,
		x.SETTLE_PERIOD,
		x.IS_SETTLE,
		x.YQ_DAYS
FROM
	(
		SELECT
			ROWNUM rn,
			T .*
		FROM
			(
				SELECT
					*
				FROM
					CREDIT.ACCOUNT_LOAN_INTEREST 
				WHERE
					store_id IN (
						SELECT
							ID
						FROM
							CREDIT.company_struct cs
						WHERE
							status = '0'
						AND TYPE = '1' START WITH ID = (
							SELECT
								struct_id
							FROM
								CREDIT.user_info
							WHERE
								user_id = 'yunwei'
						) CONNECT BY p_struct_id = PRIOR ID
					)
				AND REPORT_DATE_STR = '202005'
				ORDER BY
					ID,
					repay_day ASC
			) T
		WHERE
			ROWNUM <= 500
	) x
WHERE
	x.rn > 0
`
	data := Tests(sql)
	handledata2(data)
	elapsed := time.Since(t1)
	fmt.Println("总共用时: ", elapsed)

}
