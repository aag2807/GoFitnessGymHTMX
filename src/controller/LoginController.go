package controller

import (
	"net/http"

	"github.com/GoGym/src/domain/service"
	"github.com/GoGym/src/utils"
	lib "github.com/aag2807/triplex-to-go"
)

type LoginController struct {
	Arguments   lib.Arguments
	State       lib.State
	Renderer    *utils.PartialRenderer
	userService *service.LoginService
}

func NewLoginController() *LoginController {
	return &LoginController{
		Arguments:   lib.Arguments{},
		State:       lib.State{},
		Renderer:    utils.NewPartialRenderer(),
		userService: service.NewLoginService(),
	}
}

// Login is the handler for the /login/user that manages if the user can log in
func (c *LoginController) Login(w http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	email := req.FormValue("email")
	password := req.FormValue("password")
	c.Arguments.NotWhiteSpace(email, "email cannot be empty")
	c.Arguments.NotWhiteSpace(password, "password cannot be empty")

	if valid, err := c.userService.Login(email, password); err != nil {
		panic(err)
	} else if valid {
		w.Header().Add("HX-Reswap", "none")
		w.Header().Add("HX-Redirect", "/home")
		w.WriteHeader(http.StatusSeeOther)
	}
}
