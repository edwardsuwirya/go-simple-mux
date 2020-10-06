package deliveries

import (
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"gosimplemux/appUtils/appHttpParser"
	"gosimplemux/appUtils/appHttpResponse"
	"gosimplemux/appUtils/appStatus"
	"gosimplemux/models"
	"gosimplemux/useCases"
	"net/http"
)

const (
	loginRoute  = "/login"
	logoutRoute = "/logout"
)

type AuthDelivery struct {
	router      *mux.Router
	cookieStore *sessions.CookieStore
	parser      *appHttpParser.JsonParser
	responder   appHttpResponse.IResponder
	service     useCases.IUserAuthUseCase
}

func NewAuthDelivery(router *mux.Router, cookie *sessions.CookieStore, parser *appHttpParser.JsonParser, responder appHttpResponse.IResponder, service useCases.IUserAuthUseCase) IDelivery {
	return &AuthDelivery{router, cookie, parser, responder, service}
}

func (d *AuthDelivery) InitRoute(mdw ...mux.MiddlewareFunc) {
	d.router.HandleFunc(loginRoute, d.authRoute).Methods("POST")
	d.router.HandleFunc(logoutRoute, d.authLogoutRoute).Methods("GET")
}

func (d *AuthDelivery) authRoute(w http.ResponseWriter, r *http.Request) {
	session, _ := d.cookieStore.Get(r, "app-cookie")
	var userAuth models.UserAuth
	if err := d.parser.Parse(r, &userAuth); err != nil {
		http.Error(w, "", http.StatusBadRequest)
		return
	}
	userInfo := d.service.UserNamePasswordValidation(userAuth.UserName, userAuth.UserPassword)
	if userInfo == nil {
		http.Error(w, "", http.StatusUnauthorized)
		return
	}
	session.Values["authenticated"] = true
	err := session.Save(r, w)
	d.responder.Data(w, appStatus.Success, appStatus.StatusText(appStatus.Success), userInfo)
	if err != nil {
		http.Error(w, "", http.StatusInternalServerError)
		return
	}
}
func (d *AuthDelivery) authLogoutRoute(w http.ResponseWriter, r *http.Request) {
	session, _ := d.cookieStore.Get(r, "app-cookie")
	session.Values["authenticated"] = false
	session.Options.MaxAge = -1
	err := session.Save(r, w)
	if err != nil {
		http.Error(w, "", http.StatusInternalServerError)
		return
	}
	d.responder.Data(w, appStatus.Success, appStatus.StatusText(appStatus.Success), "Logout")
}
