package main

import (
	"github.com/gorilla/mux"
	"gosimplemux/appMiddleware"
	"gosimplemux/appUtils/appHttpParser"
	"gosimplemux/appUtils/appHttpResponse"
	"gosimplemux/deliveries"
	"gosimplemux/infra"
	"gosimplemux/manager"
)

type appRouter struct {
	app                            *goSimpleMuxApp
	sessionManger                  manager.SessionManager
	logRequestMiddleware           *appMiddleware.LogRequestMiddleware
	tokenValidationMiddleware      *appMiddleware.TokenValidationMiddleware
	tokenRedisValidationMiddleware *appMiddleware.TokenRedisValidationMiddleware
	parser                         *appHttpParser.JsonParser
	responder                      appHttpResponse.IResponder
	infra                          infra.Infra
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
			del: deliveries.NewAuthDelivery(ar.app.router, ar.sessionManger.CookieRedis(), ar.parser, ar.responder, serviceManager.UserAuthUseCase()),
			mdw: nil,
		},
		{
			del: deliveries.NewUserDelivery(ar.app.router, ar.parser, ar.responder, serviceManager.UserUseCase()),
			mdw: []mux.MiddlewareFunc{
				ar.tokenRedisValidationMiddleware.Validate,
			},
		},
	}
	for _, r := range appRoutes {
		r.del.InitRoute(r.mdw...)
	}
}

func NewAppRouter(app *goSimpleMuxApp) *appRouter {
	sessionManager := manager.NewSessionManger(app.infra)
	return &appRouter{
		app,
		sessionManager,
		appMiddleware.NewLogRequestMiddleware(),
		appMiddleware.NewTokenValidationMiddleware(sessionManager.CookieInMemory()),
		appMiddleware.NewTokenRedisValidationMiddleware(sessionManager.CookieRedis()),
		appHttpParser.NewJsonParser(),
		appHttpResponse.NewJSONResponder(),
		app.infra,
	}
}
