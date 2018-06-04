package main

import "fmt"

func main(){
	b:=3
	for x:=5; b<25; x+=3{
		fmt.Println("we can do it !", x)
		b+=4
	}
}