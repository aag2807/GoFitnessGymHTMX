package controller

import (
	"fmt"
	"log"
	"net/http"

	"github.com/GoGym/src/domain/service"
	"github.com/GoGym/src/utils"
	lib "github.com/aag2807/triplex-to-go"
)

type LoginController struct {
	Arguments    lib.Arguments
	State        lib.State
	Renderer     *utils.PartialRenderer
	userService  *service.LoginService
	pageRenderer *utils.TemplateRenderer
}

func NewLoginController() *LoginController {
	return &LoginController{
		Arguments:    lib.Arguments{},
		State:        lib.State{},
		Renderer:     utils.NewPartialRenderer(),
		userService:  service.NewLoginService(),
		pageRenderer: utils.NewTemplateRenderer("src/templates"),
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
		user, _ := c.userService.GetUserByEmail(email)
		utils.SetUserSession(w, req, user)
		http.Redirect(w, req, "/session/home", http.StatusSeeOther)
	}
}

// Logout logouts hte user
func (c *LoginController) Logout(w http.ResponseWriter, req *http.Request) {
	utils.ClearUserSession(w, req)
	http.Redirect(w, req, "/", http.StatusSeeOther)
}

func (c *LoginController) RenderLoginPage(w http.ResponseWriter, req *http.Request) {
	w.Header().Add("Content-Type", "text/html")
	w.Header().Add("cache-control", "max-age=120")

	err := c.pageRenderer.RenderHTMLTemplate(w, "login.html", nil)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
		return
	}
}

func (c *LoginController) RenderForgotPasswordPage(w http.ResponseWriter, req *http.Request) {
	w.Header().Add("Content-Type", "text/html")
	w.Header().Add("cache-control", "max-age=120")
	err := c.pageRenderer.RenderHTMLTemplate(w, "forgot-password.html", nil)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
		return
	}
}

func (c *LoginController) HandleIndexPageRedirection(w http.ResponseWriter, req *http.Request) {
	session := utils.GetUserSession(req)
	log.Println(session.Values, " in index redirection")
	if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
		http.Redirect(w, req, "/login", http.StatusSeeOther)
		return
	}

	http.Redirect(w, req, "/session/home", http.StatusSeeOther)
}

func (c *LoginController) RenderSignUpPage(w http.ResponseWriter, req *http.Request) {
	w.Header().Add("Content-Type", "text/html")
	w.Header().Add("cache-control", "max-age=120")
	err := c.pageRenderer.RenderHTMLTemplate(w, "sign-up.html", nil)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
		return
	}
}

func (c *LoginController) SignUp(w http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	username := req.FormValue("username")
	email := req.FormValue("email")
	password := req.FormValue("password")
	retypePassword := req.FormValue("retypePassword")

	c.Arguments.NotWhiteSpace(username, "username cannot be empty")
	c.Arguments.NotWhiteSpace(email, "email cannot be empty")
	c.Arguments.NotWhiteSpace(password, "password cannot be empty")
	c.Arguments.NotWhiteSpace(retypePassword, "retypePassword cannot be empty")

	log.Println(username, email, password, retypePassword)
}
