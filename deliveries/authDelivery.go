package deliveries

import (
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"log"
	"net/http"
)

const (
	loginRoute  = "/login"
	logoutRoute = "/logout"
)

type AuthDelivery struct {
	router      *mux.Router
	cookieStore *sessions.CookieStore
}

func NewAuthDelivery(r *mux.Router, c *sessions.CookieStore) *AuthDelivery {
	return &AuthDelivery{r, c}
}

func (d *AuthDelivery) InitRoute(mdw ...mux.MiddlewareFunc) {
	d.router.HandleFunc(loginRoute, d.authRoute).Methods("POST")
	d.router.HandleFunc(logoutRoute, d.authLogoutRoute).Methods("GET")
}

func (d *AuthDelivery) authRoute(w http.ResponseWriter, r *http.Request) {
	session, _ := d.cookieStore.Get(r, "app-cookie")
	//Ada mekanisme cek user name & password
	session.Values["authenticated"] = true
	err := session.Save(r, w)
	if err != nil {
		log.Fatalln(err)
	}
}
func (d *AuthDelivery) authLogoutRoute(w http.ResponseWriter, r *http.Request) {
	session, _ := d.cookieStore.Get(r, "app-cookie")
	session.Values["authenticated"] = false
	session.Options.MaxAge = -1
	err := session.Save(r, w)
	if err != nil {
		log.Fatalln(err)
	}
}
