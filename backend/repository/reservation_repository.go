package repository

import (
	"database/sql"

	"damsil99/car-reservation-app/backend/models"
)

func CreateReservation(db *sql.DB, reservation models.Reservation) error {
	query := `INSERT INTO reservations (user_id, car_id, extras, total_price) VALUES ($1, $2, $3, $4)`
	_, err := db.Exec(query, reservation.UserID, reservation.CarID, reservation.Extras, reservation.TotalPrice)
	return err
}

func GetReservationsByUserID(db *sql.DB, userID int) ([]models.Reservation, error) {
	rows, err := db.Query("SELECT id, user_id, car_id, extras, total_price FROM reservations WHERE user_id = $1", userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	reservations := []models.Reservation{}
	for rows.Next() {
		var reservation models.Reservation
		if err := rows.Scan(&reservation.ID, &reservation.UserID, &reservation.CarID, &reservation.Extras, &reservation.TotalPrice); err != nil {
			return nil, err
		}
		reservations = append(reservations, reservation)
	}
	return reservations, nil
}
