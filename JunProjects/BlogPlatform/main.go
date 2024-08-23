package main

import (
	"log"
	"net/http"

	"example.com/moviereview/database"
	"example.com/moviereview/router"
)

func main() {
	database.InitDatabase()
	r := router.Router()

	log.Println("Server is running on port: 8000")
	log.Fatal(http.ListenAndServe(":8000", r))
}