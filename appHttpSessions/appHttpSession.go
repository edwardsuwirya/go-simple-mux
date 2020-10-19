package appHttpSessions

import (
	"net/http"
)

type AppSession interface {
	Get(r *http.Request, args ...string) interface{}
	Set(w http.ResponseWriter, r *http.Request, key string, val interface{}, forClear bool, args ...string) error
}
