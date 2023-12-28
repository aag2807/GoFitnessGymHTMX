package utils

import (
	"fmt"
	"net/http"

	entities "github.com/GoGym/src/domain/entities/user"
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

func SetUserSession(w http.ResponseWriter, r *http.Request, u entities.User) {
	session := GetUserSession(r)
	session.Options = &sessions.Options{}
	session.Values["authenticated"] = true
	session.Values["userId"] = u.ID
	session.Save(r, w)
}

func ClearUserSession(w http.ResponseWriter, r *http.Request) {
	session := GetUserSession(r)
	session.Values["authenticated"] = false
	session.Values["userId"] = 0
	session.Save(r, w)
}

func GetUserSession(r *http.Request) *sessions.Session {
	session, _ := GetSessionHandler().Session.Get(r, "x-go-session")
	return session
}
