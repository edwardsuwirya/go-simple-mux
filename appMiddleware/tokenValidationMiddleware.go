package appMiddleware

import (
	"github.com/gorilla/sessions"
	"net/http"
	"time"
)

type TokenValidationMiddleware struct {
	Store *sessions.CookieStore
}

func NewTokenValidationMiddleware() *TokenValidationMiddleware {
	var store = sessions.NewCookieStore([]byte("rahasia..."))
	store.MaxAge(int(10 * time.Second))
	return &TokenValidationMiddleware{store}
}

func (v *TokenValidationMiddleware) Validate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		session, _ := v.Store.Get(r, "app-cookie")
		if isAuth, ok := session.Values["authenticated"].(bool); !isAuth || !ok {
			http.Error(w, "", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func (v *TokenValidationMiddleware) GetCookieStore() *sessions.CookieStore {
	return v.Store
}
