package appHttpSessions

import (
	"github.com/gorilla/sessions"
	"net/http"
)

type cookieSession struct {
	cookieStore *sessions.CookieStore
}

func (cs *cookieSession) Get(r *http.Request, args ...string) interface{} {
	session, _ := cs.cookieStore.Get(r, args[0])
	return session
}
func (cs *cookieSession) Set(w http.ResponseWriter, r *http.Request, key string, val interface{}, forClear bool, args ...string) error {
	session := cs.Get(r, args[0]).(*sessions.Session)
	session.Values[key] = val
	if forClear {
		session.Options.MaxAge = -1
	}
	err := session.Save(r, w)
	return err
}

func NewCookieSession(cookieStore *sessions.CookieStore) AppSession {
	return &cookieSession{cookieStore: cookieStore}
}
