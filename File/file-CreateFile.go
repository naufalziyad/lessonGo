package main

import (
	"fmt"
	"os"
)

var path = "/Users/detikcom/Documents/test_usingGo.txt"

func isError(err error) bool {
	if err != nil {
		fmt.Println(err.Error())
	}
	return (err != nil)
}

func createFile() {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		file, err := os.Create(path)
		if isError(err) {
			return
		}
		defer file.Close()
	}
	fmt.Println("File Successfully Created in -->", path)
}

func main() {
	createFile()
}
