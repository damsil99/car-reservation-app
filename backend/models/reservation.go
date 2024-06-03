package models

type Reservation struct {
	ID         int     `json:"id"`
	UserID     int     `json:"user_id"`
	CarID      int     `json:"car_id"`
	Extras     string  `json:"extras"`
	TotalPrice float64 `json:"total_price"`
}
