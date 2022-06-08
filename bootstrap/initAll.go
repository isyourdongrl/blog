package bootstrap

import (
	"fmt"
	"go-blog/controller"
	"go-blog/global"
	"net/http"
)

// InitWeb 初始化web
func InitWeb(){
	fmt.Println("start....")
	//server:=http.Server{
	//	Addr: "localhost:8080",
	//}
	global.Server.Addr="localhost:8080"
}

// InitRouter 初始化路由
func InitRouter(){
	fmt.Println("init router")
	// 默认跳到首页
	http.HandleFunc("/",controller.Index)
}
