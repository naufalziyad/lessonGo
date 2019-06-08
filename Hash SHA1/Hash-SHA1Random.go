package main

import (
	"crypto/sha1"
	"fmt"
	"time"
)

func randomHash(data string) (string, string) {
	randomPattern := fmt.Sprintf("%d", time.Now().UnixNano())
	randomText := fmt.Sprintf("Text = '%s', Pattern = '%s'", data, randomPattern)
	fmt.Println(randomText)
	sha := sha1.New()
	sha.Write([]byte(randomText))
	encrpyted := sha.Sum(nil)

	return fmt.Sprintf("%x", encrpyted), randomPattern
}

func main() {
	fmt.Println("Penerapan Enkripsi Hash SHA1 dengan random key")

	data := "ini adalah kata-kata rahasia"
	fmt.Printf("Original Text : %s\n\n", data)

	hashed1, randomPattern1 := randomHash(data)
	fmt.Printf("hashed 1 : %s \n \n", hashed1)

	hashed2, randomPattern2 := randomHash(data)
	fmt.Printf("hashed 2 : %s \n \n", hashed2)

	hashed3, randomPattern3 := randomHash(data)
	fmt.Printf("hashed 3 : %s \n \n", hashed3)

	hashed4, randomPattern4 := randomHash(data)
	fmt.Printf("hashed 4 : %s \n \n", hashed4)

	_, _, _, _ = randomPattern1, randomPattern2, randomPattern3, randomPattern4

}
