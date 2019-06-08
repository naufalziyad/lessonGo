package main

import (
	"crypto/sha1"
	"fmt"
)

func main() {
	fmt.Println("Penerapan algoritma enkripsi dengan metode Hash SHA1")
	data := "Ini adalah kode rahasia kita 74812738123"
	sha := sha1.New()
	sha.Write([]byte(data))

	enkripsi := sha.Sum(nil)
	enkripsiString := fmt.Sprintf("%x", enkripsi)

	fmt.Println(enkripsiString)
}

/* SHA1 = Secure Hash Algorithm 1
SHA1 merupakan enkripsi bersifat satu arah (one way encription) */
