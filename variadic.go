package main

import "fmt"

func main() {
	var avg = calculate(2, 4, 3, 6, 4, 7)
	var msg = fmt.Sprintf("Rata rata : %.2f", avg)
	println(msg)
}

func calculate(numbers ...int) float64 {
	var total int = 0
	for _, number := range numbers {
		total += number
	}

	var avg = float64(total) / float64(len(numbers))
	return avg
}
