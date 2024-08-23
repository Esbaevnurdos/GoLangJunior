package router

import (
	"example.com/moviereview/handlers"
	"github.com/gorilla/mux"
)


func Router() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/register", handlers.Register).Methods("POST")
	r.HandleFunc("/login", handlers.Login).Methods("POST")

	r.HandleFunc("/register", handlers.CreateMovie).Methods("POST")
	r.HandleFunc("/login", handlers.GetMovie).Methods("POST")

	r.HandleFunc("/register", handlers.CreateReview).Methods("POST")
	r.HandleFunc("/login", handlers.GetReviews).Methods("POST")
	
    return r
}