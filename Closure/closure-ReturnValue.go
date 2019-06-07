package main

import "fmt"

func findMax(number []int, max int) (int, func() []int) {
	var res []int
	for _, e := range number {
		if e <= max {
			res = append(res, e)
		}
	}
	//berikut merupakan fungsi tanpa nama yang dikembalikan dengan skema ' func() []int ' .
	return len(res), func() []int {
		return res
	}
}

func main() {
	fmt.Println("Penerapan Closure dengan return value")
	//Studi kasus ini adalah mencari berapa banyak angka dibawah angka maksimal dan menampilkan angka-angka yang berada di bawah maksimal.

	angkaMaksimal := 5
	dataAngka := []int{2, 6, 3, 3, 36, 2, 7, 3, 4, 2}
	howMany, getNumber := findMax(dataAngka, angkaMaksimal)
	hasilAngka := getNumber()

	fmt.Println("Angka yang dimiliki : ", dataAngka)
	fmt.Println("Ditemukan sebanyak : ", howMany, "angka dibawah angka maksimal")
	fmt.Println("Angka dibawah maksimal adalah : ", hasilAngka)

}
