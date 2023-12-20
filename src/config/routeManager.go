package config

import (
	"fmt"
	"log"
	"net/http"

	"github.com/GoGym/src/router"
	"github.com/go-chi/chi"
)

type RouteManager struct{}

func (rm *RouteManager) Init(r *chi.Mux) {
	r.Use(ErrorCatcher)
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

func ErrorCatcher(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Println("Recovered from error: " + fmt.Sprintf("%v", err))
				w.WriteHeader(http.StatusInternalServerError)
				http.Error(w, fmt.Sprintf("%v", err), http.StatusInternalServerError)
			}
		}()
		next.ServeHTTP(w, r)
	})
}
