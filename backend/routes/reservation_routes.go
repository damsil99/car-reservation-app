package routes

import (
	"encoding/json"
	"net/http"
	"strconv"

	"damsil99/car-reservation-app/backend/models"
	"damsil99/car-reservation-app/backend/repository"

	"github.com/gorilla/mux"
)

func RegisterReservationRoutes(r *mux.Router) {
	r.HandleFunc("/reservations", createReservation).Methods("POST")
	r.HandleFunc("/reservations/{user_id}", getUserReservations).Methods("GET")
}

func createReservation(w http.ResponseWriter, r *http.Request) {
	var reservation models.Reservation
	json.NewDecoder(r.Body).Decode(&reservation)

	db := models.SetupDB()
	defer db.Close()

	err := repository.CreateReservation(db, reservation)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func getUserReservations(w http.ResponseWriter, r *http.Request) {
	userID, _ := strconv.Atoi(mux.Vars(r)["user_id"])

	db := models.SetupDB()
	defer db.Close()

	reservations, err := repository.GetReservationsByUserID(db, userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(reservations)
}
