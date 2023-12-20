package router

import (
	"github.com/GoGym/src/controller"
	"github.com/go-chi/chi"
)

type Router interface {
	MapRoutes(chiRouter *chi.Mux)
	SetController(controller controller.Controller)
}
