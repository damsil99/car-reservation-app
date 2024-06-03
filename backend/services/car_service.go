package services

import (
	"damsil99/car-reservation-app/backend/models"
	"database/sql"
)

func GetAllCars(db *sql.DB) ([]models.Car, error) {
	// Consulta SQL para seleccionar todos los autos de la base de datos
	query := "SELECT id, make, model, year, color FROM cars"

	// Ejecutar la consulta y obtener el resultado
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Crear una lista para almacenar los autos recuperados
	var cars []models.Car

	// Iterar sobre cada fila del resultado y crear un objeto Car
	for rows.Next() {
		var car models.Car
		if err := rows.Scan(&car.ID, &car.Make, &car.Model, &car.Year, &car.Color); err != nil {
			return nil, err
		}
		// Agregar el auto a la lista
		cars = append(cars, car)
	}

	// Manejar cualquier error que pueda haber ocurrido durante la iteración
	if err := rows.Err(); err != nil {
		return nil, err
	}

	// Devolver la lista de autos recuperados
	return cars, nil
}
