package main

import (
	"os"
	"log"
	"io"
)

// Copy a file
func main() {
	// Open original file
	originalFile, err := os.Open("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer originalFile.Close()

	// Create new file
	newFile, err := os.Create("test_copy.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer newFile.Close()

	// Copy the bytes to destination from source
	bytesWritten, err := io.Copy(newFile, originalFile)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Copied %d bytes.", bytesWritten)

	// Commit the file contents
	// Flushes memory to disk
	err = newFile.Sync() // 相当与C语言的flush，将缓冲区中的内容写到持久存储中
	if err != nil {
		log.Fatal(err)
	}
}