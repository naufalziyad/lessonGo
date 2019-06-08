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

func deleteFile() {
	err := os.Remove(path)

	if isError(err) {
		return
	}

	fmt.Println("File Successfully Deleted -->", path)
}

func main() {
	deleteFile()
}
