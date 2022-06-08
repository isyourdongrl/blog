package main

import (
	"go-blog/bootstrap"
	"go-blog/global"
	"log"
)

func main() {
	// 初始化web
	bootstrap.InitWeb()
	// 初始化路由
	bootstrap.InitRouter()
	err:=global.Server.ListenAndServe()
	if err!=nil{
		log.Println(err)
	}
}
