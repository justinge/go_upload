package handler

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

func UploadHandler(w http.ResponseWriter, r *http.Request){
	if r.Method == "GET" {
		// 返回上传html页面
		data,err := ioutil.ReadFile("./static/view/index.html")
		if err != nil{
			io.WriteString(w,"internal server error"+err.Error())
			return
		}else {
			io.WriteString(w, string(data))
			return
		}
	} else if r.Method == "POST" {
		// 接收文件流以及储存到本地目录
		file, head, err := r.FormFile("file")
		if err!=nil{
			fmt.Printf("Fail to get data, err:%s\n",err.Error())
			return
		}
		defer file.Close()
		newFile, err:= os.Create("./"+head.Filename)
		if err!= nil {
			fmt.Printf("Fail to get data, err:%s\n",err.Error())
		}
		defer newFile.Close()
		_, err = io.Copy(newFile, file)
		if err != nil{
			fmt.Printf("Failed to save data into file, err:%s\n",err.Error())
			return
		}
		// 重定向
		http.Redirect(w, r,"/file/upload/suc",http.StatusFound)
	}

}

// UploadSucHandler:上传已完成
func UploadSucHandler(w http.ResponseWriter, r *http.Request){
	io.WriteString(w,"upload finished!")
}


