package main

import (
	"net/http"
	"log"
	"fmt"
	"strings"
)
func sayHello(w http.ResponseWriter, r *http.Request) {
	//w.Write([]byte("hello 世界！"))
	fmt.Println("------------------")
	r.ParseForm()
	fmt.Println(r.Form)
	fmt.Println("path：", r.URL.Path)
	fmt.Println("scheme：", r.URL.Scheme)
	fmt.Println(r.Form["url_ long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "Hello GoWeb!")
}
func main() {

	http.HandleFunc("/", sayHello)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
