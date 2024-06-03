package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"

	"damsil99/car-reservation-app/backend/models"
	"damsil99/car-reservation-app/backend/routes"
)

func main() {
	godotenv.Load()
	db := models.SetupDB()
	defer db.Close()

	r := mux.NewRouter()

	routes.RegisterAuthRoutes(r)
	routes.RegisterCarRoutes(r)
	routes.RegisterReservationRoutes(r)
	routes.RegisterUserRoutes(r)

	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8080", r))
}
