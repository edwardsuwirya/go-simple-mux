package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func homeRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hello World")
}
func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", homeRoute)
	log.Println("Running on localhost:6969")
	err := http.ListenAndServe("localhost:6969", r)
	if err != nil {
		log.Fatalln(err)
	}
}
