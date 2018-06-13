package main

import (
	"fmt"
	"net/http"
	"time"
)

func write(w http.ResponseWriter, r *http.Request) {
	//
	fmt.Fprintf(w, "Hello World! %s", time.Now())
}

func main() {
	http.HandleFunc("/", write)
	fmt.Println(time.Now().Month() + 3)
	http.ListenAndServe(":8080", nil)
}
