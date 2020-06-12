//Author: zwa
//Date: 2020/6/11

package main

import (
	"time"

	"GoleGolas/kube/pormetheus/pushgateway/bin-pushgateway/prometheus"
	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.New()
	app.Use(gin.Logger())
	prom := prometheus.NewPrometheus("gin")
	prom.UseWithAuth(app, gin.Accounts{"gin": "gonic"}, "/metrics")

	app.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"package": "pushgateway测试"})
	})
	app.GET("/version", func(c *gin.Context) {
		c.JSON(200, gin.H{"version": "1.0.0", "time": time.Now()})
	})
	app.POST("/data", func(c *gin.Context) {
		c.JSON(200, gin.H{"success": "成功！！！"})
	})

	if err := app.Run(":80"); err != nil {
		panic(err)
	}
}
