package utils

import (
	"net/http"

	"github.com/gorilla/sessions"
)

type SessionHandler struct {
	Session *sessions.CookieStore
}

var session *sessions.CookieStore

func GetSessionHandler() *SessionHandler {
	if session == nil {
		session = sessions.NewCookieStore([]byte("x-go-session"))
	}

	return &SessionHandler{Session: session}
}

func SetUserSession(w http.ResponseWriter, r *http.Request) {
	session, _ := GetSessionHandler().Session.Get(r, "x-go-session")
	session.Options = &sessions.Options{MaxAge: 216000}
	session.Values["authenticated"] = true
	session.Save(r, w)
}

func ClearUserSession(w http.ResponseWriter, r *http.Request) {
	session, _ := GetSessionHandler().Session.Get(r, "x-go-session")
	session.Values["authenticated"] = false
	session.Save(r, w)
}
