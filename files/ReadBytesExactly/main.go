package main

import (
	"os"
	"log"
	"io"
)

func main() {
	// Open file for reading
	file, err := os.Open("test.txt")
	if err != nil {
		log.Fatal(err)
	}

	// The file.Read() function will happily read a tiny file in to a large
	// byte slice, but io.ReadFull() will return an
	// error if the file is smaller than the byte slice.
	byteSlice := make([]byte, 102)
	numBytesRead, err := io.ReadFull(file, byteSlice) // 如果文件内容长度小于byteSlice长度，io.ReadFull返回unexpected EOF错误
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Number of bytes read: %d\n", numBytesRead)
	log.Printf("Data read: %s\n", byteSlice)
}