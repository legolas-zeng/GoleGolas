package main

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/prometheus/common/log"
	"net/http"
)

type fooCollector struct {
	fooMetric *prometheus.Desc
	barMetric *prometheus.Desc
}

func newFooCollector() *fooCollector {
	m1 := make(map[string]string)
	m1["env"] = "测试pod"
	v := []string{"hostname"}
	return &fooCollector{
		fooMetric: prometheus.NewDesc("fff_metrics", "Show metrics a for mysql", nil, nil),
		barMetric: prometheus.NewDesc("bbb_metrics", "Show metrics a bar occu", v, m1),
	}
}

// 检测cpu温度
var cpuTemp = prometheus.NewGauge(prometheus.GaugeOpts{
	Name: "cpu_temperature_celsius",
	Help: "CPU当前温度.",
})

var hdFailures = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "hd_errors_total",
		Help: "硬盘错误计数.",
	},
	[]string{"device"},
)

var http_request_total = promauto.NewCounter(
	prometheus.CounterOpts{
		Name: "http_request_total",
		Help: "已处理的http请求的总数",
	},
)

func (collect *fooCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- collect.barMetric
	ch <- collect.fooMetric

}

func (collect *fooCollector) Collect(ch chan<- prometheus.Metric) {
	var metricValue float64

	if 1 == 1 {
		metricValue = 1
	}
	ch <- prometheus.MustNewConstMetric(collect.fooMetric, prometheus.GaugeValue, metricValue)
	ch <- prometheus.MustNewConstMetric(collect.barMetric, prometheus.CounterValue, metricValue, "本地测试机")
}

func init() {
	foo := newFooCollector()
	//注册指标
	prometheus.MustRegister(foo)
	prometheus.MustRegister(cpuTemp)
	prometheus.MustRegister(hdFailures)
}

func main() {
	//测试用，设置温度值
	cpuTemp.Set(65.3)
	// Set和With().Inc 去改变指标的数据内容
	hdFailures.With(prometheus.Labels{"device": "/dev/sda"}).Inc()
	http.HandleFunc("/", func(http.ResponseWriter, *http.Request) {
		http_request_total.Inc()
	})
	log.Info("beging to server on Port: 18080")
	http.Handle("/metrics", promhttp.Handler())
	log.Fatal(http.ListenAndServe(":18080", nil))
}
