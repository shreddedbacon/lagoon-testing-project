package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Request at %v\n", time.Now())
	for k, v := range r.Header {
		fmt.Printf("%v: %v\n", k, v)
		w.Write([]byte(fmt.Sprintf("%v: %v\n", k, v)))
	}
}
