package main

import (
	"flag"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type goSimpleMuxApp struct {
	host   string
	port   string
	router *mux.Router
}

func (app *goSimpleMuxApp) run() {
	h := fmt.Sprintf("%s:%s", app.host, app.port)
	log.Println("Listening on", h)
	NewAppRouter(app).InitMainRouter()
	err := http.ListenAndServe(h, app.router)
	if err != nil {
		log.Fatalln(err)
	}
}

func NewGoSimpleMuxApp() *goSimpleMuxApp {
	hostPtr := flag.String("host", "localhost", "Listening on host")
	portPtr := flag.String("port", "6969", "Listening on port")
	flag.Parse()
	r := mux.NewRouter()
	return &goSimpleMuxApp{
		host:   *hostPtr,
		port:   *portPtr,
		router: r,
	}
}

func main() {
	NewGoSimpleMuxApp().run()
}
