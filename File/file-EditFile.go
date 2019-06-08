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

func writeFile() {
	//open file dan berikan hak akses RW
	file, err := os.OpenFile(path, os.O_RDWR, 0644)
	if isError(err) {
		return
	}
	defer file.Close()

	_, err = file.WriteString("Ini percobaan untuk menulis kan data didalam file: \n")
	if isError(err) {
		return
	}

	_, err = file.WriteString("i love Go Language")
	if isError(err) {
		return
	}

	//save update
	err = file.Sync()
	if isError(err) {
		return
	}

	fmt.Println("File Successfully Update -->", path)
}

func main() {
	writeFile()
}
