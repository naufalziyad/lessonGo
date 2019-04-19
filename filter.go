package main

import (
	"fmt"
	"strings"
)

func main() {
	var data = []string{"Dono", "Kasino", "Indro", "Waluya"}
	var dataTerdapatO = filter(data, func(each string) bool {
		return strings.Contains(each, "o")
	})
	var dataLength5 = filter(data, func(each string) bool {
		return len(each) == 5
	})

	fmt.Println("data asli \t\t:", data)

	fmt.Println("filter ada hurut O nya :", dataTerdapatO)

	fmt.Println("filter jumlah huruf ada 5:", dataLength5)
}

func filter(data []string, callback func(string) bool) []string {
	var result []string
	for _, each := range data {
		if filtered := callback(each); filtered {
			result = append(result, each)
		}
	}
	return result
}
