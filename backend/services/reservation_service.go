package services

import (
	"damsil99/car-reservation-app/backend/models"
	"database/sql"
)

func CreateReservation(db *sql.DB, reservation models.Reservation) error {
	// Preparar la consulta SQL para insertar una nueva reserva en la base de datos
	query := "INSERT INTO reservations (user_id, car_id, extras, total_price) VALUES ($1, $2, $3, $4)"

	// Ejecutar la consulta SQL con los valores de la reserva proporcionados
	_, err := db.Exec(query, reservation.UserID, reservation.CarID, reservation.Extras, reservation.TotalPrice)
	if err != nil {
		// Manejar cualquier error que pueda ocurrir al ejecutar la consulta
		return err
	}

	// La reserva se ha creado correctamente
	return nil
}
