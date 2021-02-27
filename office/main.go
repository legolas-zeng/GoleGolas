//Author: zwa
//Date: 2020/7/27

package main

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/tealeg/xlsx"
)

var (
	ExeclFile = "C:\\Users\\Administrator\\Desktop\\特殊销账.xlsx"
)

func main() {
	xlFile, err := xlsx.OpenFile(ExeclFile)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	sheet := xlFile.Sheets[0]
	fmt.Println("工作表名: ", sheet.Name)
	map1 := make(map[string]interface{})
	for _, row := range sheet.Rows[1:] {
		number := row.Cells[0]
		beego.Info(number)
		xiaozhang := fmt.Sprintf("%s", number)
		fmt.Printf("%T \n", xiaozhang)
		map1[xiaozhang] = "1"
	}
	fmt.Println(map1)
}
