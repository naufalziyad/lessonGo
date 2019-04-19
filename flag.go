package main

import ("flag"
	"fmt")

func main(){
	var nama = flag.String("nama", "anonymous", "ketik nama anda")
	var umur = flag.Int64("umur", 25, "ketik nama usia")

	flag.Parse()
	fmt.Printf("name\t: %s\n", *nama)
	fmt.Printf("age\t: %d\n", *umur)
}