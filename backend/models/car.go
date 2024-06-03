package models

type Car struct {
	ID           int     `json:"id"`
	Make         string  `json:"make"`
	Model        string  `json:"model"`
	Year         int     `json:"year"`
	Color        string  `json:"color"`
	Type         string  `json:"type"`
	Transmission string  `json:"transmission"`
	Fuel         string  `json:"fuel"`
	Price        float64 `json:"price"`
}
