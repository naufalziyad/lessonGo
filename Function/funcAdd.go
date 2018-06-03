package main

// author : Naufal Ziyad Luthfiansyah //

import ("fmt")

func add(x,y float32) float32 {
	return x+y
}

func main(){
	var num1, num2 float32 = 5.6, 9.5

	fmt.Println(add(num1,num2))
}
