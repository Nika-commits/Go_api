package main

import (
	"net/http"

	"github.com/go-chi/chi"

	// "github.com/Nika-commits/`Go_api/internal/handlers"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetReportCaller(true)
	var r *chi.Mux = chi.NewRouter()
	handlers.Handler(r)

	fmt.Prinln("--------------- Starting Go API Service by Pranish ---------------")

	err := http.ListenAndServe("8080", r)

	if err != nil {
		log.Error(err)
	} 
}
