package main

import (
	"flag"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
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
	session, _ := store.Get(r, "app-cookie")
	user1 := User{
		Id:        "123",
		FirstName: "Edi",
		LastName:  "Uchida",
	}
	if isAuth, ok := session.Values["authenticated"].(bool); !isAuth || !ok {
		http.Error(w, "", http.StatusUnauthorized)
		return
	}

	responder.Data(w, appStatus.Success, appStatus.StatusText(appStatus.Success), user1)
}

func authRoute(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "app-cookie")
	store.Options = &sessions.Options{
		MaxAge:   60 * 1,
		HttpOnly: true,
	}
	session.Values["authenticated"] = true
	err := session.Save(r, w)
	if err != nil {
		log.Fatalln(err)
	}
}
func authLogoutRoute(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "app-cookie")
	session.Values["authenticated"] = false
	session.Options.MaxAge = -1
	err := session.Save(r, w)
	fmt.Println(session.Values)
	if err != nil {
		log.Fatalln(err)
	}
}

var responder = appHttpResponse.NewJSONResponder()
var store = sessions.NewCookieStore([]byte("rahasia..."))

func main() {
	hostPtr := flag.String("host", "localhost", "Listening on host")
	portPtr := flag.String("port", "6969", "Listening on port")
	flag.Parse()

	r := mux.NewRouter()
	r.HandleFunc("/user", homeRoute)
	r.HandleFunc("/user/{id}", homeRoute)

	r.HandleFunc("/login", authRoute)
	r.HandleFunc("/logout", authLogoutRoute)

	h := fmt.Sprintf("%s:%s", *hostPtr, *portPtr)
	log.Println("Listening on", h)
	err := http.ListenAndServe(h, r)
	if err != nil {
		log.Fatalln(err)
	}
}
