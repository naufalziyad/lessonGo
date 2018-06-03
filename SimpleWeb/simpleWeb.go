package main

// author : Naufal Ziyad Luthfiansyah //

import ("fmt"
		"net/http")

func index_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Amazing Go Language")
}

func about_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "-Naufal Ziyad L")
}

func main(){
	http.HandleFunc("/", index_handler)
	http.HandleFunc("/about/", about_handler)
	http.ListenAndServe(":8000",nil)
}