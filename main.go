package main

import "gin-test/router"

func main() {
	// gin.SetMode(gin.ReleaseMode)
	r := router.GetRouter()
	r.LoadHTMLGlob("templates/*")
	r.Static("/static", "static")
	r.Run(":80") // 监听并在 0.0.0.0:8080 上启动服务
}
