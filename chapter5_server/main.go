// chapter5_server project main.go
package main

import (
	"io"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello,World")
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "<h1>This is just for test</h1>")
}

func main() {

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/hello", helloHandler)
	err := http.ListenAndServe(":8088", nil)

	if err != nil {
		log.Fatal("ListenAndServe:", err.Error())
	} else {
		log.Println("ListenAndServe success")
	}

}
