package main

// author : Naufal Ziyad Luthfiansyah //

import ("fmt"
		"net/http")

func index_handler (a http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(a, `<h1>Hallo my name Naufal Ziyad Luthfiansyah</h1>
		<p>this is my first time learn Go Language </p>
		<p> Go is awesome, fast and simple !`)
}

func index_about (a http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(a, `<h1>this is about page`)
}

func main(){
	http.HandleFunc("/", index_handler)
	http.HandleFunc("/about", index_about)
	http.ListenAndServe(":8000", nil) //port yang digunakan
}