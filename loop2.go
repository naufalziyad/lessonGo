package main

// author : Naufal Ziyad Luthfiansyah //

import "fmt"

func main(){
	a :=5
	for {
		fmt.Println("Learn about Go Language", a)
		a+=3
		if a > 25{
			break
		}
	}
}
