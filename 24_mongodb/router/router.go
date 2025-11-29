package router

import (
	"mongo_db/controller"

	"github.com/gorilla/mux"
)

func Router() *mux.Router{
	r := mux.NewRouter()
	r.HandleFunc("/api/movies", controller.GetMyAllMovies).Methods("GET")
	r.HandleFunc("/api/movies", controller.CreateMovie).Methods("POST")
	r.HandleFunc("/api/movies/{id}", controller.MarkedAsWatched).Methods("PUT")
	r.HandleFunc("/api/movies/{id}", controller.DeleteMovie).Methods("DELETE")
	r.HandleFunc("/api/deleteallmovie", controller.DeleteAllMovie).Methods("DELETE")

	return r

}
