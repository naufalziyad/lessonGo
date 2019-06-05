package main

import (
	"fmt"
	"regexp"
)

func main() {
	fmt.Println("Split")
	text := "Hai, saya naufal"

	regex, _ := regexp.Compile(`[a-z]+`)

	str := regex.Split(text, -1)
	fmt.Printf("%#v \n", str)

}
