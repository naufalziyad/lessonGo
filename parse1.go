package main

// author : Naufal Ziyad Luthfiansyah //

import ("fmt"
"net/http"
"io/ioutil")

func main(){
	resp, _ := http.Get("https://www.washingtonpost.com/news-sitemap-index.xml")
	bytes, _ := ioutil.ReadAll(resp.Body)
	string_body := string(bytes)
	fmt.Println(string_body)
	resp.Body.Close()
}