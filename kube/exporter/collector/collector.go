package collector

import (
	"github.com/prometheus/client_golang/prometheus"
	"sync"
)

type Metrics struct {
	metrics map[string]*prometheus.Desc
	mutex   sync.Mutex
}

type pod struct {
	podName   string
	namespace string
}

func newGlobalMetric(namespace string, metricName string, docString string, labels []string) *prometheus.Desc {
	return prometheus.NewDesc(namespace+"_"+metricName, docString, labels, nil)
}

func NewMetrics(namespace string) *Metrics {
	return &Metrics{
		metrics: map[string]*prometheus.Desc{
			"pod_network_receive_bytes": newGlobalMetric(namespace, "pod_network_receive_bytes", "network receive bytes of pod ", []string{"pod_name", "namespace", "node_name"}),
		},
	}
}

/**
 * 接口：Describe
 * 功能：传递结构体中的指标描述符到channel
 */
func (c *Metrics) Describe(ch chan<- *prometheus.Desc) {
	for _, m := range c.metrics {
		ch <- m
	}
}

/**
 * 接口：Collect
 * 功能：抓取最新的数据，传递给channel
 */
func (c *Metrics) Collect(ch chan<- prometheus.Metric) {
	var pod = pod{}
	c.mutex.Lock() // 加锁
	defer c.mutex.Unlock()
	// do something ,add metric to channel
	ch <- prometheus.MustNewConstMetric(c.metrics["pod_network_receive_bytes"], prometheus.CounterValue, 0, "node-exporter-gl84b", "default", "192.168.3.17")
}
