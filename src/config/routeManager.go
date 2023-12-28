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
	r.Use(VerifyIsHTMXCall)

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
		log.Println(r.URL)
		// session, err := utils.GetSessionHandler().Session.Get(r, "x-go-session")
		// if err != nil {
		// 	panic(err)
		// }

		// // if strings.Contains(r.URL.Path, "/session") {
		// // 	if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
		// // 		http.Error(w, "Forbidden", http.StatusForbidden)
		// // 		return
		// // 	}
		// // }

		next.ServeHTTP(w, r)
	})
}

func VerifyIsHTMXCall(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		route := r.URL.Path
		if r.Header.Get("HX-Request") != "" {
			next.ServeHTTP(w, r)
		} else if strings.Contains(route, "/static") {
			next.ServeHTTP(w, r)
		} else if strings.Contains(route, "/session") && strings.Contains(route, "/home") && !strings.Contains(route, "/page") {
			next.ServeHTTP(w, r)
		} else if strings.Contains(route, "/session") && r.Header.Get("HX-Request") == "" {
			http.Redirect(w, r, "/session/home", http.StatusSeeOther)
		} else {
			next.ServeHTTP(w, r)
		}
	})
}
