package main

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"time"
)

func recordMetrics() {
	for {
		ops.Inc()

		time.Sleep(2 * time.Second)
	}
}

var (
	ops = promauto.NewCounter(prometheus.CounterOpts{
		Name: "mxshop_test",
		Help: "just for test",
	})
)

func main() {
	go recordMetrics()
	r := gin.Default()
	r.GET("/metrics", gin.WrapH(promhttp.Handler()))
	r.Run(":8050")
}
