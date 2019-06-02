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

	res1 := regex.FindAllString(text, 2)
	fmt.Printf("%#v \n", res1)

	res2 := regex.FindAllString(text, -1)
	fmt.Printf("%#v \n", res2)

}
