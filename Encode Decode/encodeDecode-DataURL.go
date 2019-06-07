package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	data := "https://detik.com/"
	encodedString := base64.URLEncoding.EncodeToString([]byte(data))
	fmt.Println("Hasil Encode : ", encodedString)

	decodedByte, _ := base64.URLEncoding.DecodeString(encodedString)

	decodedString := string(decodedByte)
	fmt.Println("Hasil Decoded : ", decodedString)
}
