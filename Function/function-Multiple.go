package main

// author : Naufal Ziyad Luthfiansyah //

import "fmt"

func multiple(a, b string) (string, string) {
	return a, b
}

func main() {
	a, b := "hello", "my name Naufal"

	fmt.Println(multiple(a, b))
}
