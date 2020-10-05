package main

import (
	"github.com/gorilla/mux"
	"gosimplemux/infra"
	"log"
	"net/http"
)

type goSimpleMuxApp struct {
	infra  infra.Infra
	router *mux.Router
}

func (app *goSimpleMuxApp) run() {
	h := app.infra.ApiServer()
	log.Println("Listening on", h)
	NewAppRouter(app).InitMainRouter()
	err := http.ListenAndServe(h, app.router)
	if err != nil {
		log.Fatalln(err)
	}
}

func NewGoSimpleMuxApp() *goSimpleMuxApp {
	r := mux.NewRouter()
	appInfra := infra.NewInfra()
	return &goSimpleMuxApp{
		infra:  appInfra,
		router: r,
	}
}

func main() {
	NewGoSimpleMuxApp().run()
}
