package routes

import (
	"damsil99/car-reservation-app/backend/models"
	"damsil99/car-reservation-app/backend/repository"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func RegisterUserRoutes(r *mux.Router) {
	r.HandleFunc("/users/{username}", getUser).Methods("GET")
}

func getUser(w http.ResponseWriter, r *http.Request) {
	username := mux.Vars(r)["username"]

	db := models.SetupDB()
	defer db.Close()

	user, err := repository.GetUserByUsername(db, username)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(user)
}
