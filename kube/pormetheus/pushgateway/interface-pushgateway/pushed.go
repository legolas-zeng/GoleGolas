//Author: zwa
//Date: 2020/6/12

package main

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/push"
)

//var(
//	records = prometheus.NewGauge(prometheus.GaugeOpts{
//		Name: "db_backup_records_processed",
//		Help: "在上次数据库备份中处理的记录数",
//	})
//)

var (
	pushurl = "http://192.168.3.16:9091"
)

type Pushs interface {
	sent() error
	pushinfo()
}

type Info struct {
	Name  string
	Help  string
	value int
	job   string
}

func (p Info) sent() error {
	registry := prometheus.NewRegistry()
	records := prometheus.NewGauge(prometheus.GaugeOpts{Name: p.Name, Help: p.Help})
	registry.MustRegister(records)
	pusher := push.New(pushurl, p.job).Gatherer(registry)
	records.Set(float64(p.value))

	if err := pusher.Add(); err != nil {
		fmt.Println("不能推送数据到pushgateway:", err)
		return err
	} else {
		return nil
	}
}

func (p Info) pushinfo() {
	fmt.Println("123")

}

func Work(i Pushs) {
	i.sent()
}

func test() {
	a := make(map[string]interface{})
	a["测试"] = 123
	a["结果"] = 1.321
	a["过程"] = "abcdef"
}

func main() {

	info := Info{job: "db_backup", Name: "db_backup_records_processed", Help: "在上次数据库备份中处理的记录数", value: 123456}

	Work(info)
}

//func main() {
//	p1:=Info{job:"db_backup",Name:"db_backup_records_processed",Help:"在上次数据库备份中处理的记录数",value:123456}
//	p1.sent()
//}
