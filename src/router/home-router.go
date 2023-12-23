package router

import (
	"fmt"
	"net/http"

	"github.com/GoGym/src/controller"
	"github.com/GoGym/src/utils"
	"github.com/go-chi/chi"
)

type HomeRouter struct {
	Controller *controller.Controller
	renderer   *utils.TemplateRenderer
}

func (r *HomeRouter) MapRoutes(chiRouter *chi.Mux) {
	r.renderer = utils.NewTemplateRenderer("src/templates")
	r.renderer.LoadTemplates()

	chiRouter.Route("/session", func(chiRouter chi.Router) {
		chiRouter.Get("/home", func(w http.ResponseWriter, req *http.Request) {
			w.Header().Add("Content-Type", "text/html")
			err := r.renderer.LoadAndExecuteTemplateWithDashboardLayout(w, "pages/home.html", nil)
			if err != nil {
				fmt.Println(err)
				http.Error(w, "Error rendering template", http.StatusInternalServerError)
				return
			}
		})

		chiRouter.Get("/admin-profile", func(w http.ResponseWriter, req *http.Request) {
			err := r.renderer.LoadAndExecuteTemplateWithDashboardLayout(w, "pages/admin-profile.html", nil)
			if err != nil {
				fmt.Println(err)
				http.Error(w, "Error rendering template", http.StatusInternalServerError)
				return
			}
		})

		chiRouter.Get("/coaches", func(w http.ResponseWriter, req *http.Request) {
			w.Header().Add("Content-Type", "text/html")
			err := r.renderer.LoadAndExecuteTemplateWithDashboardLayout(w, "pages/coaches.html", nil)
			if err != nil {
				fmt.Println(err)
				http.Error(w, "Error rendering template", http.StatusInternalServerError)
				return
			}
		})

		chiRouter.Get("/inventory", func(w http.ResponseWriter, req *http.Request) {
			w.Header().Add("Content-Type", "text/html")
			err := r.renderer.LoadAndExecuteTemplateWithDashboardLayout(w, "pages/inventory.html", nil)
			if err != nil {
				fmt.Println(err)
				http.Error(w, "Error rendering template", http.StatusInternalServerError)
				return
			}
		})

		chiRouter.Get("/members", func(w http.ResponseWriter, req *http.Request) {
			w.Header().Add("Content-Type", "text/html")
			err := r.renderer.LoadAndExecuteTemplateWithDashboardLayout(w, "pages/members.html", nil)
			if err != nil {
				fmt.Println(err)
				http.Error(w, "Error rendering template", http.StatusInternalServerError)
				return
			}
		})

		chiRouter.Get("/payment", func(w http.ResponseWriter, req *http.Request) {
			w.Header().Add("Content-Type", "text/html")
			err := r.renderer.LoadAndExecuteTemplateWithDashboardLayout(w, "pages/payment.html", nil)
			if err != nil {
				fmt.Println(err)
				http.Error(w, "Error rendering template", http.StatusInternalServerError)
				return
			}
		})

		chiRouter.Get("/plans", func(w http.ResponseWriter, req *http.Request) {
			w.Header().Add("Content-Type", "text/html")
			err := r.renderer.LoadAndExecuteTemplateWithDashboardLayout(w, "pages/plans.html", nil)
			if err != nil {
				fmt.Println(err)
				http.Error(w, "Error rendering template", http.StatusInternalServerError)
				return
			}
		})

		chiRouter.Get("/registration", func(w http.ResponseWriter, req *http.Request) {
			w.Header().Add("Content-Type", "text/html")
			err := r.renderer.LoadAndExecuteTemplateWithDashboardLayout(w, "pages/registration.html", nil)
			if err != nil {
				fmt.Println(err)
				http.Error(w, "Error rendering template", http.StatusInternalServerError)
				return
			}
		})

		chiRouter.Get("/reports", func(w http.ResponseWriter, req *http.Request) {
			w.Header().Add("Content-Type", "text/html")
			err := r.renderer.LoadAndExecuteTemplateWithDashboardLayout(w, "pages/reports.html", nil)
			if err != nil {
				fmt.Println(err)
				http.Error(w, "Error rendering template", http.StatusInternalServerError)
				return
			}
		})
	})
}
