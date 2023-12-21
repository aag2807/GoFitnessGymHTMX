package router

import (
	"fmt"
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
	userController := controller.NewLoginController()
	r.renderer = utils.NewTemplateRenderer("src/templates")
	r.renderer.LoadTemplates()

	chiRouter.Get("/", func(w http.ResponseWriter, req *http.Request) {
		http.Redirect(w, req, "/login", http.StatusSeeOther)
	})

	chiRouter.Get("/login", func(w http.ResponseWriter, req *http.Request) {
		w.Header().Add("Content-Type", "text/html")
		w.Header().Add("cache-control", "max-age=120")

		err := r.renderer.RenderHTMLTemplate(w, "login.html", nil)
		if err != nil {
			fmt.Println(err)
			http.Error(w, "Error rendering template", http.StatusInternalServerError)
			return
		}
	})

	chiRouter.Get("/forgot-password", func(w http.ResponseWriter, req *http.Request) {
		w.Header().Add("Content-Type", "text/html")
		w.Header().Add("cache-control", "max-age=120")
		err := r.renderer.RenderHTMLTemplate(w, "forgot-password.html", nil)
		if err != nil {
			fmt.Println(err)
			http.Error(w, "Error rendering template", http.StatusInternalServerError)
			return
		}
	})

	chiRouter.Post("/login/user", func(w http.ResponseWriter, req *http.Request) {
		userController.Login(w, req)
	})

	chiRouter.Get("/logout/user", func(w http.ResponseWriter, req *http.Request) {
		userController.Logout(w, req)
	})
}
