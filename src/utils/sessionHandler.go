package utils

import (
	"fmt"
	"net/http"

	"github.com/gorilla/sessions"
)

var session *sessions.CookieStore

type SessionHandler struct {
	Session *sessions.CookieStore
}

func init() {
	fmt.Println("registering session handler")
	session = sessions.NewCookieStore([]byte("x-go-session"))
}

func GetSessionHandler() *SessionHandler {
	if session == nil {
		fmt.Println("registering session handler again")
		session = sessions.NewCookieStore([]byte("x-go-session"))
	}

	return &SessionHandler{Session: session}
}

func SetUserSession(w http.ResponseWriter, r *http.Request) {
	session := GetUserSession(r)
	session.Options = &sessions.Options{MaxAge: 216000}
	session.Values["authenticated"] = true
	session.Save(r, w)
}

func ClearUserSession(w http.ResponseWriter, r *http.Request) {
	session := GetUserSession(r)
	session.Values["authenticated"] = false
	session.Save(r, w)
}

func GetUserSession(r *http.Request) *sessions.Session {
	session, _ := GetSessionHandler().Session.Get(r, "x-go-session")
	return session
}
