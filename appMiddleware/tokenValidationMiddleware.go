package appMiddleware

import (
	"github.com/gorilla/sessions"
	"gosimplemux/appHttpSessions"
	"net/http"
)

type TokenValidationMiddleware struct {
	appSession appHttpSessions.AppSession
}

func NewTokenValidationMiddleware(appSession appHttpSessions.AppSession) *TokenValidationMiddleware {
	return &TokenValidationMiddleware{appSession}
}

func (v *TokenValidationMiddleware) Validate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		session := v.appSession.Get(r, "app-cookie").(*sessions.Session)
		if isAuth, ok := session.Values["authenticated"].(bool); !isAuth || !ok {
			http.Error(w, "", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}
