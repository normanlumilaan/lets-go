package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)

	fmt.Println("Listening port 4000")

	err := http.ListenAndServe(":4000", mux)

	if err != nil {
		log.Fatal("Listening error")
	}

}
