package main

// author : Naufal Ziyad Luthfiansyah //

import "fmt"

func main(){
	grades := make(map[string]float32)

	grades["Naufal"] = 29
	grades["Wulan"] = 16
	grades["Timmy"] = 10

	/*TimsGrade := grades["Timmy"]
	fmt.Println(TimsGrade)

	delete(grades, "Timmy")*/
	fmt.Println(grades)

	for k, v :=range grades {
		fmt.Println(k, ":",v)
	}
}