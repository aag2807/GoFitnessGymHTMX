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
			err := r.renderer.LoadAndExecuteTemplateWithDashboardLayout(w, "home.html", nil)
			if err != nil {
				fmt.Println(err)
				http.Error(w, "Error rendering template", http.StatusInternalServerError)
				return
			}
		})
	})
}
