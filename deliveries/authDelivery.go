package deliveries

import (
	"github.com/gorilla/mux"
	"gosimplemux/appHttpSessions"
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
	router     *mux.Router
	appSession appHttpSessions.AppSession
	parser     *appHttpParser.JsonParser
	responder  appHttpResponse.IResponder
	service    useCases.IUserAuthUseCase
}

func NewAuthDelivery(router *mux.Router, appSession appHttpSessions.AppSession, parser *appHttpParser.JsonParser, responder appHttpResponse.IResponder, service useCases.IUserAuthUseCase) IDelivery {
	return &AuthDelivery{router, appSession, parser, responder, service}
}

func (d *AuthDelivery) InitRoute(mdw ...mux.MiddlewareFunc) {
	d.router.HandleFunc(loginRoute, d.authRoute).Methods("POST")
	d.router.HandleFunc(logoutRoute, d.authLogoutRoute).Methods("GET")
}

func (d *AuthDelivery) authRoute(w http.ResponseWriter, r *http.Request) {

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
	err := d.appSession.Set(w, r, "authenticated", true, false, "app-cookie")
	if err != nil {
		http.Error(w, "", http.StatusInternalServerError)
		return
	}
	d.responder.Data(w, appStatus.Success, appStatus.StatusText(appStatus.Success), userInfo)
}
func (d *AuthDelivery) authLogoutRoute(w http.ResponseWriter, r *http.Request) {
	err := d.appSession.Set(w, r, "authenticated", false, true, "app-cookie")
	if err != nil {
		http.Error(w, "", http.StatusInternalServerError)
		return
	}
	d.responder.Data(w, appStatus.Success, appStatus.StatusText(appStatus.Success), "Logout")
}
