package deliveries

import (
	"fmt"
	"github.com/gorilla/mux"
	"gosimplemux/appHttpParser"
	"gosimplemux/appHttpResponse"
	"gosimplemux/appStatus"
	"gosimplemux/models"
	"gosimplemux/useCases"
	"net/http"
)

const (
	userMainRoute = "/user"
)

type UserDelivery struct {
	router    *mux.Router
	parser    *appHttpParser.JsonParser
	responder appHttpResponse.IResponder
	service   useCases.IUserUseCase
}

func NewUserDelivery(router *mux.Router, parser *appHttpParser.JsonParser, responder appHttpResponse.IResponder, service useCases.IUserUseCase) *UserDelivery {
	return &UserDelivery{
		router, parser, responder, service,
	}
}

func (d *UserDelivery) InitRoute(mdw ...mux.MiddlewareFunc) {
	userRouter := d.router.PathPrefix(userMainRoute).Subrouter()
	userRouter.Use(mdw...)

	userRouter.HandleFunc("", d.userRoute).Methods("GET")
	userRouter.HandleFunc("", d.userPostRoute).Methods("POST")
	userRouter.HandleFunc("", d.userPutRoute).Methods("PUT")
	userRouter.HandleFunc("", d.userDeleteRoute).Methods("DELETE")
	userRouter.HandleFunc("/{id}", d.userRoute).Methods("GET")
}

func (d *UserDelivery) userRoute(w http.ResponseWriter, r *http.Request) {
	users := d.service.GetAll()
	d.responder.Data(w, appStatus.Success, appStatus.StatusText(appStatus.Success), users)
}

func (d *UserDelivery) userPostRoute(w http.ResponseWriter, r *http.Request) {
	var newUser models.User
	if err := d.parser.Parse(r, &newUser); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	d.service.Create(&newUser)
	d.responder.Data(w, appStatus.Success, appStatus.StatusText(appStatus.Success), newUser)
}

func (d *UserDelivery) userPutRoute(w http.ResponseWriter, r *http.Request) {
	userId, isExist := r.URL.Query()["id"]
	if isExist {
		var usrReq models.User
		if err := d.parser.Parse(r, &usrReq); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		userUpdate := d.service.Update(userId[0], &usrReq)
		d.responder.Data(w, appStatus.Success, appStatus.StatusText(appStatus.Success), userUpdate)
	} else {
		msg := appStatus.StatusText(appStatus.ErrorLackInfo)
		d.responder.Error(w, appStatus.ErrorLackInfo, fmt.Sprintf(msg, "ID"))
	}
}

func (d *UserDelivery) userDeleteRoute(w http.ResponseWriter, r *http.Request) {
	userId, isExist := r.URL.Query()["id"]
	if isExist {
		d.service.Delete(userId[0])
		d.responder.Data(w, appStatus.Success, appStatus.StatusText(appStatus.Success), nil)
	} else {
		msg := appStatus.StatusText(appStatus.ErrorLackInfo)
		d.responder.Error(w, appStatus.ErrorLackInfo, fmt.Sprintf(msg, "ID"))
	}

}
