package main

import (
	"log"
	"os"
)

func main() {
	// Truncate a file to 100 bytes. If file
	// is less than 100 bytes the original contents will remain
	// at the beginning, and the rest of the space is
	// filled will null bytes. If it is over 100 bytes,
	// Everything past 100 bytes will be lost. Either way
	// we will end up with exactly 100 bytes.
	// Pass in 0 to truncate to a completely empty file

	//如果文件内容超过100字节，后续多出100的被截取丢弃，前面100个字节被保留
	// 如果文件内容不超过100字节，文件内容会被保留，且不足100字节的部分会被用0值填充至100字节
	//第二个参数为0,就是将文件内容全部清空
	err := os.Truncate("test.txt", 100) //
	if err != nil {
		log.Fatal(err)
	}
}