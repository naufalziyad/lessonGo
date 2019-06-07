package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	data := "Hallo Nama saya Naufal Ziyad "
	encodedString := base64.StdEncoding.EncodeToString([]byte(data))
	fmt.Println("Hasil Encode : ", encodedString)

	decodedByte, _ := base64.StdEncoding.DecodeString(encodedString)
	fmt.Println("Proses Decoded Byte : ", decodedByte)
	decodedString := string(decodedByte)
	fmt.Println("Hasil Decoded : ", decodedString)
}
