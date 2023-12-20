package controller

import (
	"fmt"
	"net/http"
)

type LoginController struct{}

// Login is the handler for the /login/user that manages if the user can log in
func (c *LoginController) Login(w http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	email := req.FormValue("email")
	password := req.FormValue("password")

	fmt.Println(email, password)

	w.WriteHeader(http.StatusNoContent)
}
