package config

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/GoGym/src/router"
	"github.com/GoGym/src/utils"
	"github.com/go-chi/chi"
)

type RouteManager struct{}

func (rm *RouteManager) Init(r *chi.Mux) {
	r.Use(ErrorCatcher)
	r.Use(SessionVerifier)
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
				renderer := utils.NewPartialRenderer()
				errorTemplate := renderer.GetTemplatePartialToRender("error-notification.html")
				w.Header().Add("HX-Retarget", "#errors")
				w.Header().Add("HX-Reswap", "innerHTML")
				errorTemplate.Execute(w, utils.ResponseMessage{Message: fmt.Sprintf("%v", err)})
			}
		}()
		next.ServeHTTP(w, r)
	})
}

func SessionVerifier(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		session, err := utils.GetSessionHandler().Session.Get(r, "x-go-session")
		log.Println(session.Values["authenticated"])

		if err != nil {
			panic(err)
		}

		if strings.Contains(r.URL.Path, "/session") {
			if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
				http.Error(w, "Forbidden", http.StatusForbidden)
				return
			}

			log.Println("Session is valid")
		}

		next.ServeHTTP(w, r)
	})
}
