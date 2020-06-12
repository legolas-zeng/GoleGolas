//Author: zwa
//Date: 2020/6/11

package main

//
//import (
//	"fmt"
//	"github.com/prometheus/client_golang/prometheus"
//	"github.com/prometheus/client_golang/prometheus/push"
//)
//
//var (
//	pushUrl = "http://192.168.3.16:9091/"
//)
//func ExamplePusher_Push() {
//	completionTime := prometheus.NewGauge(prometheus.GaugeOpts{
//		Name: "db_backup_last_completion_timestamp_seconds",
//		Help: "The timestamp of the last successful completion of a DB backup.",
//	})
//	completionTime.SetToCurrentTime()
//	if err := push.New(pushUrl, "db_backup").Collector(completionTime).Grouping("db", "customers").Grouping("instance","测试数据").Push();
//	err != nil {
//		fmt.Println("Could not push completion time to Pushgateway:", err)
//	}
//}
//func main() {
//	ExamplePusher_Push()
//}
import (
	"fmt"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/push"
)

var (
	completionTime = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "db_backup_last_completion_timestamp_seconds",
		Help: "数据库备份的最后一次完成的时间戳，成功或者失败",
	})
	successTime = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "db_backup_last_success_timestamp_seconds",
		Help: "最后一次成功完成数据库备份的时间戳",
	})
	duration = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "db_backup_duration_seconds",
		Help: "最后一次数据库备份的持续时间，以秒为单位",
	})
	records = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "db_backup_records_processed",
		Help: "在上次数据库备份中处理的记录数",
	})
)

func performBackup() (int, error) {
	// Perform the backup and return the number of backed up records and any
	// applicable error.
	// ...
	return 42, nil
}

func ExamplePusher_Add() {
	registry := prometheus.NewRegistry()
	registry.MustRegister(completionTime, duration, records)
	pusher := push.New("http://192.168.3.16:9091", "db_backup").Gatherer(registry)
	start := time.Now()
	n, err := performBackup()
	records.Set(float64(n))
	duration.Set(time.Since(start).Seconds())
	completionTime.SetToCurrentTime()
	if err != nil {
		fmt.Println("DB backup failed:", err)
	} else {
		pusher.Collector(successTime)
		successTime.SetToCurrentTime()
	}
	if err := pusher.Add(); err != nil {
		fmt.Println("Could not push to Pushgateway:", err)
	}
}

func main() {
	ExamplePusher_Add()
}
