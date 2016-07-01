package main

import (
	"log"
	"io/ioutil"
)

func main() {
	// Read file to byte slice
	data, err := ioutil.ReadFile("test.txt") // 返回未知长度的byte[]
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Data read: %s\n", data)
}