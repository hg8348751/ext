package ext

import (
	"net/http"

	"github.com/nevata/session"
)

//Handler 增加一个session参数
type Handler interface {
	ServeHTTP(*session.Session, http.ResponseWriter, *http.Request)
}

//HandlerFunc x
type HandlerFunc func(s *session.Session, w http.ResponseWriter, r *http.Request)

// ServeHTTP calls f(w, r).
func (f HandlerFunc) ServeHTTP(s *session.Session, w http.ResponseWriter, r *http.Request) {
	f(s, w, r)
}

//Auth 令牌检查
func Auth(inner Handler) Handler {
	return HandlerFunc(func(s *session.Session, w http.ResponseWriter, r *http.Request) {
		sess := session.SessionMgr().GetSession(w, r)
		if sess == nil {
			w.WriteHeader(http.StatusUnauthorized)
			//handlerError(w, fmt.Errorf("令牌无效！"))
			return
		}
		inner.ServeHTTP(sess, w, r)
	})
}
