package controller

import (
	"net/http"

	"github.com/GoGym/src/utils"
	lib "github.com/aag2807/triplex-to-go"
)

type LoginController struct {
	Arguments lib.Arguments
	State     lib.State
	Renderer  *utils.PartialRenderer
}

func NewLoginController() *LoginController {

	return &LoginController{
		Arguments: lib.Arguments{},
		State:     lib.State{},
		Renderer:  utils.NewPartialRenderer(),
	}
}

// Login is the handler for the /login/user that manages if the user can log in
func (c *LoginController) Login(w http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	email := req.FormValue("email")
	password := req.FormValue("password")
	c.Arguments.NotWhiteSpace(email, "email cannot be empty")
	c.Arguments.NotWhiteSpace(password, "password cannot be empty")

	// validate using service
	// if valid, redirect to home
	// if not valid, redirect to login with error message

	if false {
		http.Redirect(w, req, "/home", http.StatusSeeOther)
	} else {
		errorTemplate := c.Renderer.GetTemplatePartialToRender("error-notification.html")
		errorTemplate.Execute(w, utils.ResponseMessage{Message: "Invalid credentials"})
	}
}
