package appHttpSessions

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	guuid "github.com/google/uuid"
	"net/http"
)

type redisSession struct {
	pool *redis.Pool
}

func (rs *redisSession) GetCookie(r *http.Request, name string) string {
	c, err := r.Cookie(name)
	if err == http.ErrNoCookie {
		return ""
	}
	return c.Value
}
func (rs *redisSession) Get(r *http.Request, args ...string) interface{} {
	tokenId := rs.GetCookie(r, args[0])
	resp, _ := rs.pool.Get().Do("HGET", tokenId, args[1])
	return fmt.Sprintf("%s", resp)
}
func (rs *redisSession) Set(w http.ResponseWriter, r *http.Request, key string, val interface{}, forClear bool, args ...string) error {
	if forClear {
		cookieVal := rs.GetCookie(r, args[0])
		_, err := rs.pool.Get().Do("DEL", cookieVal)
		if err != nil {
			return err
		}
		return nil
	} else {
		tokenId := guuid.New().String()
		_, err := rs.pool.Get().Do("HSET", tokenId, key, val)
		http.SetCookie(w, &http.Cookie{
			Name:  args[0],
			Value: tokenId,
		})
		return err
	}

}

func NewRedisSession(pool *redis.Pool) AppSession {
	return &redisSession{pool}
}
