package manager

import (
	"gosimplemux/appHttpSessions"
	"gosimplemux/appUtils/appCookieStore"
	"gosimplemux/infra"
)

type SessionManager interface {
	CookieInMemory() appHttpSessions.AppSession
	CookieRedis() appHttpSessions.AppSession
}
type sessionManager struct {
	infra infra.Infra
}

func (sm *sessionManager) CookieInMemory() appHttpSessions.AppSession {
	var cookieStore = appCookieStore.NewAppCookieStore().Store
	sess := appHttpSessions.NewCookieSession(cookieStore)
	return sess
}
func (sm *sessionManager) CookieRedis() appHttpSessions.AppSession {
	pool := sm.infra.RedisServer()
	sess := appHttpSessions.NewRedisSession(pool)
	return sess
}
func NewSessionManger(infra infra.Infra) SessionManager {
	return &sessionManager{infra: infra}
}
