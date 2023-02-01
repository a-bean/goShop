package main

import (
	hpb "GoStart/api"
	"GoStart/gin_grpc"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	helloSrv := server.NewHelloServer()

	engine := gin.Default()
	hpb.RegisterGreeterServerHTTPServer(helloSrv, engine)

	server := &http.Server{
		Addr:    ":8082",
		Handler: engine,
	}
	//支持自动生成端口以及定义ip和端口
	_ = engine.SetTrustedProxies(nil)

	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		panic(err)
	}
}
