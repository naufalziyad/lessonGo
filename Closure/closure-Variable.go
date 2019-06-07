package main

import "fmt"

func main() {
	fmt.Println("Ini adalah penerapan closure yang disimpan dalam bentuk variable")

	getMinMax := func(n []int) (int, int) {
		var min, max int
		for i, e := range n {
			switch {
			case i == 0:
				max, min = e, e
			case e > max:
				max = e
			case e < min:
				min = e
			}
		}
		return min, max
	}

	numbers := []int{5, 12, 42, 2, 12, 11}
	var min, max = getMinMax(numbers)
	fmt.Printf("Angka tersedia : %v\n angka terkecil : %v\n angka terbesar :%v\n ", numbers, min, max)
}
