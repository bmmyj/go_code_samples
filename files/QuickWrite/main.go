// The ioutil package has a useful function called WriteFile() that will handle creating/opening,
// writing a slice of bytes, and closing. It is useful if you just need a quick way to dump
// a slice of bytes to a file.
package main

import (
	"io/ioutil"
	"log"
)

func main() {
	err := ioutil.WriteFile("test.txt", []byte("Hi\n"), 0666)
	if err != nil {
		log.Fatal(err)
	}
}