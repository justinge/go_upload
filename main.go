package main

import (
	"fmt"
	"go_upload/handler"
	"net/http"
)
func main() {
	//注册路由
	http.HandleFunc("/file/upload",handler.UploadHandler)
	http.HandleFunc("/file/upload/suc",handler.UploadSucHandler)
	err:=http.ListenAndServe(":8080",nil)
	if err !=nil{
		fmt.Printf("Failed to start server,err:%s",err.Error())
	}

}
