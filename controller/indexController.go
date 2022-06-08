package controller

import (
	"net/http"
	"text/template"
)


// Index 首页
func Index(w http.ResponseWriter,r *http.Request){
	w.Header().Set("content-type","text/html;charset=utf-8")
	index:=template.New("index.html")
	index,_=index.ParseFiles("./template/index.html")
	index.Execute(w,nil)
}

