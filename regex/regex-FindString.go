package main

import (
	"fmt"
	"regexp"
)

func main() {
	text := "Nama Saya Naufal Ziyad Luthfiansyah"
	regex, err := regexp.Compile(`[a-z]+`)

	if err != nil {
		fmt.Println("Error nih, cek lagi ya")
	}

	fmt.Println("------------------------- \nRegex Find String")
	//Regex Find String
	pola := regex.FindString(text)
	fmt.Println("Pola bernilai : ", pola) //fungsi ini hanya mengembalikan 1 buah hasil saja, akan dikembalikan yang pertama saja.

	fmt.Println("------------------------- \nRegex Find String Index")
	//Regex Find String Index
	idx := regex.FindStringIndex(text)
	fmt.Println(idx)

	str := text[1:4] //nilai index 1:4 ini didapatkan dari hasil idx atas. (masih manual dulu y)
	fmt.Println(str)

	fmt.Println("------------------------- \nRegex Find All String")
	//Regex Find All String
	str1 := regex.FindAllString(text, -1) //jika dimasukkan nilai -1 , makan akan ditampilkan semua data yang sesuai dengan pola regex.
	fmt.Println(str1)

	str2 := regex.FindAllString(text, 1) //akan ditampilkan 1 data yang sesuai dengan pola regex, biasanya yang pertama.
	fmt.Println(str2)

	fmt.Println("------------------------- \nRegex Replace All String")
	//Regex Replace All String
	strRep := regex.ReplaceAllString(text, "asik") //ketika terpenuhi, makan akan di ganti dengan "asik".
	fmt.Println(strRep)

	fmt.Println("------------------------- \nRegex Replace All String Func")
	//Regex Replace All String Funx
	strFunc := regex.ReplaceAllStringFunc(text, func(each string) string {
		if each == "Saya" {
			return "Ahai"
		}
		return each
	})
	fmt.Println(strFunc)

}
