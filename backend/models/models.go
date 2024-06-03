package models

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func SetupDB() *sql.DB {
	// Configurar la conexión a la base de datos
	db, err := sql.Open("postgres", "user=youruser dbname=yourdb password=yourpassword host=yourhost sslmode=disable")
	if err != nil {
		panic(err)
	}

	// Verificar si la conexión a la base de datos es exitosa
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Conexión a la base de datos establecida")

	return db
}
