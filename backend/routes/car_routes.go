package routes

import (
	"encoding/json"
	"net/http"

	"damsil99/car-reservation-app/backend/models"
	"damsil99/car-reservation-app/backend/repository"

	"github.com/gorilla/mux"
)

func RegisterCarRoutes(r *mux.Router) {
	r.HandleFunc("/cars", getAllCars).Methods("GET")
}

func getAllCars(w http.ResponseWriter, r *http.Request) {
	db := models.SetupDB()
	defer db.Close()

	cars, err := repository.GetAllCars(db)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(cars)
}
