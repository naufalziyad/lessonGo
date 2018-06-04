package main

// author : Naufal Ziyad Luthfiansyah //

import "fmt"

const usixteenbitmax int = 2
const kmh_multiple int = 2

type car struct {
	gas int
	pedal int
	steering int
	top_speed int
}

func (c car) kmh() int {
	return int(c.gas) * (c.top_speed/usixteenbitmax)
}

func (c car) mph() int {
	return int (c.gas) * (c.top_speed/usixteenbitmax/kmh_multiple)
}

func main(){
	a_car :=car{20, 10,2,100}

	fmt.Println(a_car.pedal)
	fmt.Println(a_car.kmh())
	fmt.Println(a_car.mph())
}