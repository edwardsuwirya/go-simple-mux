package main

import (
	"github.com/gorilla/sessions"
	"gosimplemux/appCookieStore"
	"gosimplemux/appMiddleware"
	"gosimplemux/deliveries"
)

//var users = []models.User{
//	{
//		Id:        "c01d7cf6-ec3f-47f0-9556-a5d6e9009a43",
//		FirstName: "Edi",
//		LastName:  "Uchida",
//	},
//}

//func userRoute(w http.ResponseWriter, r *http.Request) {
//	responder.Data(w, appStatus.Success, appStatus.StatusText(appStatus.Success), users)
//}
//
//func userPostRoute(w http.ResponseWriter, r *http.Request) {
//	var newUser models.User
//	if err := jsonParser.Parse(r, &newUser); err != nil {
//		http.Error(w, err.Error(), http.StatusBadRequest)
//		return
//	}
//	id := guuid.New()
//	newUser.Id = id.String()
//	users = append(users, newUser)
//	responder.Data(w, appStatus.Success, appStatus.StatusText(appStatus.Success), newUser)
//}
//
//func userPutRoute(w http.ResponseWriter, r *http.Request) {
//	userId, isExist := r.URL.Query()["id"]
//	var userUpdate models.User
//	var userIdx int
//	if isExist {
//		for idx, usr := range users {
//			if usr.Id == userId[0] {
//				userUpdate = usr
//				userIdx = idx
//				break
//			}
//		}
//		var usrReq models.User
//		if err := jsonParser.Parse(r, &usrReq); err != nil {
//			http.Error(w, err.Error(), http.StatusBadRequest)
//			return
//		}
//		userUpdate.FirstName = usrReq.FirstName
//		userUpdate.LastName = usrReq.LastName
//		users[userIdx] = userUpdate
//		responder.Data(w, appStatus.Success, appStatus.StatusText(appStatus.Success), userUpdate)
//	} else {
//		msg := appStatus.StatusText(appStatus.ErrorLackInfo)
//		responder.Error(w, appStatus.ErrorLackInfo, fmt.Sprintf(msg, "ID"))
//	}
//}
//
//func userDeleteRoute(w http.ResponseWriter, r *http.Request) {
//	userId, isExist := r.URL.Query()["id"]
//	var newUsers = make([]models.User, 0)
//	if isExist {
//		for _, usr := range users {
//			if usr.Id == userId[0] {
//				continue
//			}
//			newUsers = append(newUsers, usr)
//		}
//		users = newUsers
//		responder.Data(w, appStatus.Success, appStatus.StatusText(appStatus.Success), nil)
//	} else {
//		msg := appStatus.StatusText(appStatus.ErrorLackInfo)
//		responder.Error(w, appStatus.ErrorLackInfo, fmt.Sprintf(msg, "ID"))
//	}
//
//}
type appRouter struct {
	app                       *goSimpleMuxApp
	cookieStore               *sessions.CookieStore
	logRequestMiddleware      *appMiddleware.LogRequestMiddleware
	tokenValidationMiddleware *appMiddleware.TokenValidationMiddleware
}

func (ar *appRouter) InitMainRouter() {
	//var responder = appHttpResponse.NewJSONResponder()
	//var jsonParser = appHttpParser.NewJsonParser()

	ar.app.router.Use(ar.logRequestMiddleware.Log)
	deliveries.NewAuthDelivery(ar.app.router, ar.cookieStore).InitRoute()
	//userRouter := app.router.PathPrefix("/user").Subrouter()
	//userRouter.Use(tokenValidationMiddleware.Validate)
	//
	//userRouter.HandleFunc("", userRoute).Methods("GET")
	//userRouter.HandleFunc("", userPostRoute).Methods("POST")
	//userRouter.HandleFunc("", userPutRoute).Methods("PUT")
	//userRouter.HandleFunc("", userDeleteRoute).Methods("DELETE")
	//userRouter.HandleFunc("/{id}", userRoute).Methods("GET")
}

func NewAppRouter(app *goSimpleMuxApp) *appRouter {
	var cookieStore = appCookieStore.NewAppCookieStore().Store
	return &appRouter{
		app,
		cookieStore,
		appMiddleware.NewLogRequestMiddleware(),
		appMiddleware.NewTokenValidationMiddleware(cookieStore)}
}
