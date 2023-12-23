package test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/GoGym/src/utils"
	"github.com/gorilla/sessions"
	"github.com/stretchr/testify/assert"
)

var SessionHandler *utils.SessionHandler

func setup() {
	session := sessions.NewCookieStore([]byte("x-go-session"))
	SessionHandler = &utils.SessionHandler{Session: session}
}

func teardown() {
	SessionHandler = nil
}

func TestSessionHandlerCanBeProperlyInstantiated(t *testing.T) {
	setup()
	assert := assert.New(t)

	assert.NotNil(SessionHandler)

	teardown()
}

func TestSessionHandlerCanSetUserSession(t *testing.T) {
	setup()
	defer teardown()
	assert := assert.New(t)

	_, r := CreateMockHttpResponseRequest("/login/user")

	assert.NotNil(SessionHandler.Session.Get(r, "x-go-session"))
}

func TestSessionHandlerWillMaintainValuesWhenTwoRequestsAreMade(t *testing.T) {
	setup()
	defer teardown()
	assert := assert.New(t)
	w, r := CreateMockHttpResponseRequest("/login/user")
	utils.SetUserSession(w, r)
	cookie := w.Result().Cookies()[0]

	_, r = CreateMockHttpResponseRequest("/session/home")
	r.AddCookie(cookie)

	session, _ := utils.GetSessionHandler().Session.Get(r, "x-go-session")
	assert.Equal(true, session.Values["authenticated"])
}

func CreateMockHttpResponseRequest(url string) (*httptest.ResponseRecorder, *http.Request) {
	w := httptest.NewRecorder()

	r, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}
	return w, r
}
