package main

import (
	"fmt"
	"regexp"
)

func main() {
	text := "Nama Saya Naufal Ziyad Luthfiansyah"
	var regex, err = regexp.Compile(`[a-z]+`)

	if err != nil {
		fmt.Println("Error nih")
	}

	fmt.Println("------------------------- \nRegex Match String")
	//Regex Match String
	var pola = regex.MatchString(text)
	fmt.Println("Pola bernilai : ", pola) //jika kondisi terpenuhi maka nilai true

}
