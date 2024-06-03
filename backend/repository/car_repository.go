package repository

import (
	"database/sql"

	"damsil99/car-reservation-app/backend/models"
)

func GetAllCars(db *sql.DB) ([]models.Car, error) {
	rows, err := db.Query("SELECT id, make, model, year, color, type, transmission, fuel, price FROM cars")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	cars := []models.Car{}
	for rows.Next() {
		var car models.Car
		if err := rows.Scan(&car.ID, &car.Make, &car.Model, &car.Year, &car.Color, &car.Type, &car.Transmission, &car.Fuel, &car.Price); err != nil {
			return nil, err
		}
		cars = append(cars, car)
	}
	return cars, nil
}
