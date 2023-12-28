package router

import (
	"net/http"

	"github.com/GoGym/src/controller"
	"github.com/GoGym/src/utils"
	"github.com/go-chi/chi"
)

type LoginRouter struct {
	Controller *controller.Controller
	renderer   *utils.TemplateRenderer
}

func (r *LoginRouter) MapRoutes(chiRouter *chi.Mux) {
	loginController := controller.NewLoginController()

	chiRouter.Get("/", func(w http.ResponseWriter, req *http.Request) {
		loginController.HandleIndexPageRedirection(w, req)
	})

	chiRouter.Get("/login", func(w http.ResponseWriter, req *http.Request) {
		loginController.RenderLoginPage(w, req)
	})

	chiRouter.Get("/forgot-password", func(w http.ResponseWriter, req *http.Request) {
		loginController.RenderForgotPasswordPage(w, req)
	})

	chiRouter.Get("/login/user", func(w http.ResponseWriter, req *http.Request) {
		loginController.Login(w, req)
	})

	chiRouter.Get("/logout/user", func(w http.ResponseWriter, req *http.Request) {
		loginController.Logout(w, req)
	})

	chiRouter.Get("/sign-up", func(w http.ResponseWriter, req *http.Request) {
		loginController.RenderSignUpPage(w, req)
	})

	chiRouter.Post("/sign-up/user", func(w http.ResponseWriter, req *http.Request) {
		loginController.SignUp(w, req)
	})
}
