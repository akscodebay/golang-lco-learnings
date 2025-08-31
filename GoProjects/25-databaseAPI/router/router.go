package router

import (
	"log"
	"net/http"

	"github.com/akscodebay/databaseapi/controller"
	"github.com/gorilla/mux"
)

func GenerateRoutes() {
	router := mux.NewRouter()

	router.HandleFunc("/movies", controller.CreateMovie).Methods("POST")
	router.HandleFunc("/movies", controller.GetMovies).Methods("GET")
	router.HandleFunc("/movies/{id}", controller.GetMovie).Methods("GET")
	router.HandleFunc("/movies", controller.UpdateMovie).Methods("PUT")
	router.HandleFunc("/movies/{id}", controller.DeleteMovie).Methods("DELETE")
	router.HandleFunc("/movies", controller.DeleteAll).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", router))
}
