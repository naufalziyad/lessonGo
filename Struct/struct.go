package main

// author : Naufal Ziyad Luthfiansyah //

import "fmt"

type car struct {
	gas_pedal uint16 
	brake_pedal uint16
	steering_wheel int16
	top_speed_kmh float64
}

func main(){
	a_car :=car{gas_pedal: 2000,
		brake_pedal: 0,
		steering_wheel: 12512,
		top_speed_kmh:225.0}

	/*a_car := car{2000, 120, 111, 223}*/ //Ini adalah kondisi untuk alternatif

	fmt.Println(a_car.gas_pedal)
}