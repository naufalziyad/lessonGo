package main

import (
	"fmt"
	"io"
	"os"
)

var path = "/Users/detikcom/Documents/test_usingGo.txt"

func isError(err error) bool {
	if err != nil {
		fmt.Println(err.Error())
	}
	return (err != nil)
}

func readFile() {
	file, err := os.OpenFile(path, os.O_RDWR, 0644)
	if isError(err) {
		return
	}
	defer file.Close()

	data := make([]byte, 1024)
	for {
		n, err := file.Read(data)
		if err != io.EOF {
			if isError(err) {
				break
			}
		}
		if n == 0 {
			break
		}
	}
	if isError(err) {
		return
	}
	fmt.Println("File Successfull Open -->", path)
	fmt.Println(string(data))

}

func main() {
	readFile()
}
