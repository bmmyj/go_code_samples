package main

import (
	"log"
	"os"
)

var ( // 变量批量定义
	newFile *os.File
	err     error
)

func main() {
	newFile, err = os.Create("test.txt") // 如果目标文件已经存在，Creata函数会截取原有文件内容，之前内容会丢失。
	if err != nil {
		log.Fatal(err)
	}
	defer newFile.Close()
	log.Println(newFile) // log flag默认打印日期时间:LstdFlags     = Ldate | Ltime // initial values for the standard logger
}
