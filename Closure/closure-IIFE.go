package main

import "fmt"

func main() {
	fmt.Println("ini merupakan penerapan closure Immediately Invoked Function Expression (IIFE)")
	//berikut contoh penerapan IIFE untuk filtering data berupa array
	numbers := []int{5, 1, 35, 6, 3, 3, 1, 312}
	newNumbers := func(min int) []int {
		var r []int
		for _, e := range numbers {
			if e < min {
				continue
			}
			r = append(r, e)
		}
		return r
	}(3)
	fmt.Println("Ori Number =", numbers)
	fmt.Println("New Number =", newNumbers)
}
