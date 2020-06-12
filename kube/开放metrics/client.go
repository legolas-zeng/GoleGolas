package main

import (
	"fmt"
	"github.com/Unknwon/goconfig"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type Cmd struct {
	ReqType  int
	FileName string
}

func recordMetrics() {
	go func() {
		for {
			opsProcessed.Inc()
			time.Sleep(2 * time.Second)
		}
	}()
}

var (
	cpuTemp = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "cpu_temperature_celsius",
			Help: "CPU当前温度.",
		})
	hdFailures = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "hd_errors_total",
			Help: "硬盘错误计数.",
		},
		[]string{"device"},
	)
	http_request_total = promauto.NewCounter(
		prometheus.CounterOpts{
			Name: "http_request_total",
			Help: "已处理的http请求的总数",
		})
	opsProcessed = promauto.NewCounter(
		prometheus.CounterOpts{
			Name: "myapp_processed_ops_total",
			Help: "已处理事件的总数",
		})
)

func init() {
	// Metrics have to be registered to be exposed:
	prometheus.MustRegister(cpuTemp)
	prometheus.MustRegister(hdFailures)
}

func main() {
	recordMetrics()
	cfg, err := goconfig.LoadConfigFile("D:\\GoleGolas\\kube\\开放metrics\\conf.ini")
	//cfg, err := goconfig.LoadConfigFile("/root/conf.ini")
	fmt.Println(cfg)
	if err != nil {
		log.Println("读取配置文件失败:", err)
		return
	}

	url, err := cfg.GetValue("DEFAULT", "url")
	//fmt.Println(value)
	if err != nil {
		log.Println("读取值失败:", err)
		return
	}
	//url := "http://127.0.0.1:8000/handle"
	cpuTemp.Set(65.3)
	hdFailures.With(prometheus.Labels{"device": "/dev/sda"}).Inc()
	http.Handle("/metrics", promhttp.Handler())
	//log.Fatal(http.ListenAndServe(":8888", nil))

	http.HandleFunc("/", func(w http.ResponseWriter, _ *http.Request) {
		http_request_total.Inc()
		resp, err := http.Get(url)

		if err != nil {
			log.Println("Post failed:", err)
			return
		}

		defer resp.Body.Close()

		content, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Println("Read failed:", err)
			return
		}

		log.Println("返回值:", string(content))
		fmt.Fprint(w, "返回值:", string(content))
	})
	log.Fatal(http.ListenAndServe(":8081", nil))

}
