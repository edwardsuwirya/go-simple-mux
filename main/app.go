package main

import (
	"flag"
	"fmt"
	"github.com/gorilla/mux"
	"gosimplemux/appHttpResponse"
	"gosimplemux/appStatus"
	"log"
	"net/http"
)

type User struct {
	Id        string `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

func homeRoute(w http.ResponseWriter, r *http.Request) {
	user1 := User{
		Id:        "123",
		FirstName: "Edi",
		LastName:  "Uchida",
	}
	responder.Data(w, appStatus.Success, appStatus.StatusText(appStatus.Success), user1)
}

var responder = appHttpResponse.NewJSONResponder()

func main() {
	hostPtr := flag.String("host", "localhost", "Listening on host")
	portPtr := flag.String("port", "6969", "Listening on port")
	flag.Parse()
	r := mux.NewRouter()
	r.HandleFunc("/user", homeRoute)
	r.HandleFunc("/user/{id}", homeRoute)
	h := fmt.Sprintf("%s:%s", *hostPtr, *portPtr)
	log.Println("Listening on", h)
	err := http.ListenAndServe(h, r)
	if err != nil {
		log.Fatalln(err)
	}
}
