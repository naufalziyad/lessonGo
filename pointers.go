package main

// author : Naufal Ziyad Luthfiansyah //

import "fmt"

func main(){
	x :=29
	a := &x //memory address

	fmt.Println(x)
	fmt.Println(a)
	*a = 5
	fmt.Println(x)
	*a = *a**a
	fmt.Println(x)
	fmt.Println(*a)

}