package main

import (
	"log"
	"net/http"

	"github.com/GoGym/src/config"
	"github.com/go-chi/chi"
)

func main() {
	rm := config.RouteManager{}
	r := chi.NewRouter()
	rm.Init(r)

	log.Println("Server running on port 8080")
	http.ListenAndServe(":8080", r)
}
