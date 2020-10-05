package main

import (
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"gosimplemux/appMiddleware"
	"gosimplemux/appUtils/appCookieStore"
	"gosimplemux/appUtils/appHttpParser"
	"gosimplemux/appUtils/appHttpResponse"
	"gosimplemux/deliveries"
	"gosimplemux/infra"
	"gosimplemux/manager"
)

type appRouter struct {
	app                       *goSimpleMuxApp
	cookieStore               *sessions.CookieStore
	logRequestMiddleware      *appMiddleware.LogRequestMiddleware
	tokenValidationMiddleware *appMiddleware.TokenValidationMiddleware
	parser                    *appHttpParser.JsonParser
	responder                 appHttpResponse.IResponder
	infra                     infra.Infra
}

type appRoutes struct {
	del deliveries.IDelivery
	mdw []mux.MiddlewareFunc
}

func (ar *appRouter) InitMainRouter() {
	ar.app.router.Use(ar.logRequestMiddleware.Log)
	var serviceManager = manager.NewServiceManger(ar.infra)
	appRoutes := []appRoutes{
		{
			del: deliveries.NewAuthDelivery(ar.app.router, ar.cookieStore, ar.parser, ar.responder, serviceManager.UserAuthUseCase()),
			mdw: nil,
		},
		{
			del: deliveries.NewUserDelivery(ar.app.router, ar.parser, ar.responder, serviceManager.UserUseCase()),
			mdw: []mux.MiddlewareFunc{
				ar.tokenValidationMiddleware.Validate,
			},
		},
	}
	for _, r := range appRoutes {
		r.del.InitRoute(r.mdw...)
	}
}

func NewAppRouter(app *goSimpleMuxApp) *appRouter {
	var cookieStore = appCookieStore.NewAppCookieStore().Store
	return &appRouter{
		app,
		cookieStore,
		appMiddleware.NewLogRequestMiddleware(),
		appMiddleware.NewTokenValidationMiddleware(cookieStore),
		appHttpParser.NewJsonParser(),
		appHttpResponse.NewJSONResponder(),
		app.infra,
	}
}
