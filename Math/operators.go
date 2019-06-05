package main

import (
	"fmt"
	"math"
)

//pembulatan
func round(num float64) float64 {
	if num < 0 {
		return math.Ceil(num - 0.5)
	}
	return math.Floor(num + 0.5)
}

func main() {
	var meal_cost float64
	var tip_percent int
	var tax_percent int

	fmt.Scan(&meal_cost)
	fmt.Scan(&tip_percent)
	fmt.Scan(&tax_percent)

	tip := meal_cost * (float64(tip_percent) / 100)
	tax := meal_cost * (float64(tax_percent) / 100)
	totalCost := meal_cost + tip + tax
	var t = int(round(totalCost))
	fmt.Printf("Total %d ", t)

}
