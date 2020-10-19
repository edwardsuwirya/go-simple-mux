package appMiddleware

import (
	"gosimplemux/appHttpSessions"
	"net/http"
)

type TokenRedisValidationMiddleware struct {
	appSession appHttpSessions.AppSession
}

func NewTokenRedisValidationMiddleware(appSession appHttpSessions.AppSession) *TokenRedisValidationMiddleware {
	return &TokenRedisValidationMiddleware{
		appSession,
	}
}

func (v *TokenRedisValidationMiddleware) Validate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		session := v.appSession.Get(r, appHttpSessions.CookieName, appHttpSessions.AuthCookieKey)
		if session == appHttpSessions.IsAuthenticated {
			next.ServeHTTP(w, r)
		} else {
			http.Error(w, "", http.StatusUnauthorized)
			return
		}

	})
}
