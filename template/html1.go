package main

import ("net/http"
		"html/template")

type NewsAggPage struct {
	Title string
	News string
}
/*
func newsAggHandler(w http.ResponseWriter, r *http.Request){
	p := NewsAggPage{Title: "News Agregator", News: "Hello lesson GO Language"}
	t, _ := template.ParseFiles("index.html")
	t.Execute(w, p)
}*/

func indexHandler(w http.ResponseWriter, r *http.Request){
	p := NewsAggPage{Title: "News Agregator", News: "Hello lesson GO Language"}
	t, _ := template.ParseFiles("index.html")
	t.Execute(w, p)
}

func main(){
	http.HandleFunc("/", indexHandler)
	/*http.HandleFunc("/agg/", newsAggHandler)*/
	http.ListenAndServe(":8000", nil)
}