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

	log.Println("Server running on port 8888")

	err := http.ListenAndServe(":8888", r)
	if http.ListenAndServe(":8888", r) != nil {
		log.Fatal("Error running server " + err.Error())
	}
}
