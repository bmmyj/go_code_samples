package main

import (
	"log"
	"os"
)

func main() {
	err := os.Remove("test.txt") // 如果目标文件不存在，会返回no such file or directory错误
	if err != nil {
		log.Fatal(err)
	}
}