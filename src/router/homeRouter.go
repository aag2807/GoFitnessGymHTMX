package router

import (
	"net/http"

	"github.com/GoGym/src/controller"
	"github.com/go-chi/chi"
)

type HomeRouter struct{}

func (r *HomeRouter) MapRoutes(chiRouter *chi.Mux) {
	dashboardController := controller.NewDashboardController()

	chiRouter.Route("/session", func(chiRouter chi.Router) {
		chiRouter.Get("/home", func(w http.ResponseWriter, req *http.Request) {
			dashboardController.RenderInitialDashboardHome(w, req)
		})

		chiRouter.Get("/admin-profile", func(w http.ResponseWriter, req *http.Request) {
			dashboardController.RenderAdminProfile(w, req)
		})

		chiRouter.Get("/coaches", func(w http.ResponseWriter, req *http.Request) {
			dashboardController.RenderCoachesPage(w, req)
		})

		chiRouter.Get("/inventory", func(w http.ResponseWriter, req *http.Request) {
			dashboardController.RenderInventoryPage(w, req)
		})

		chiRouter.Get("/members", func(w http.ResponseWriter, req *http.Request) {
			dashboardController.RenderMembersPage(w, req)
		})

		chiRouter.Get("/payment", func(w http.ResponseWriter, req *http.Request) {
			dashboardController.RenderPaymentPage(w, req)
		})

		chiRouter.Get("/plans", func(w http.ResponseWriter, req *http.Request) {
			dashboardController.RenderPlansPage(w, req)
		})

		chiRouter.Get("/registration", func(w http.ResponseWriter, req *http.Request) {
			dashboardController.RenderRegistrationPage(w, req)
		})

		chiRouter.Get("/reports", func(w http.ResponseWriter, req *http.Request) {
			dashboardController.RenderReportsPage(w, req)
		})
	})
}
