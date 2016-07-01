// The example above copies the entire file in to memory. This was for convenience to pass it as a parameter
// to each of the hash functions. Another approach is to create the hash writer interface and write to it
// using Write(), WriteString(), or in this case, Copy(). The example below uses the md5 hash, but you can
// switch to use any of the others that are supported.

package main

import (
	"crypto/md5"
	"log"
	"fmt"
	"io"
	"os"
)

func main() {
	// Open file for reading
	file, err := os.Open("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Create new hasher, which is a writer interface
	hasher := md5.New()
	_, err = io.Copy(hasher, file)
	if err != nil {
		log.Fatal(err)
	}

	// Hash and print. Pass nil since
	// the data is not coming in as a slice argument
	// but is coming through the writer interface
	sum := hasher.Sum(nil)
	fmt.Printf("Md5 checksum: %x\n", sum)
}