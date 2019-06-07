package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	data := "percobaan penerapan fungsi encode dan decode"

	//func encoded
	encoded := make([]byte, base64.StdEncoding.EncodedLen(len(data)))
	base64.StdEncoding.Encode(encoded, []byte(data))
	encodedString := string(encoded)
	fmt.Println("Hasil Encoded : ", encodedString)

	//func decoded
	decoded := make([]byte, base64.StdEncoding.DecodedLen(len(encoded)))
	_, err := base64.StdEncoding.Decode(decoded, encoded)
	if err != nil {
		fmt.Println(err.Error())
	}
	decodedString := string(decoded)
	fmt.Println("Hasil Decode : ", decodedString)
}
