package config

import (
	"net/http"

	"github.com/GoGym/src/router"
	"github.com/go-chi/chi"
)

type RouteManager struct{}

func (rm *RouteManager) Init(r *chi.Mux) {
	handleStaticAssetsEndpoint(r)

	homeRouter := router.HomeRouter{}
	homeRouter.MapRoutes(r)

	loginRouter := router.LoginRouter{}
	loginRouter.MapRoutes(r)
}

const staticDir = "./src/static/"

func handleStaticAssetsEndpoint(r *chi.Mux) {
	fileServer := http.FileServer(http.Dir(staticDir))
	r.Handle("/static/*", http.StripPrefix("/static/", fileServer))
}
