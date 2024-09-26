package main

import (
	"log"
	"net/http"
)

func main() {
	// 设置静态文件目录
	fs := http.FileServer(http.Dir("web_resource/dist/"))

	// 设置路由，将所有 HTTP 请求转发给文件服务器
	http.Handle("/", fs)

	// 监听并在 8089 端口启动服务器
	log.Println("Listening on :8089...")
	err := http.ListenAndServe(":8089", nil)
	if err != nil {
		log.Fatal(err)
	}
}
