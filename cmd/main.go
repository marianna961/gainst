package main

import (
	"database/sql"
	"fmt"
	"log"
	//"net/http"
	_ "github.com/lib/pq"
    "github.com/marianna961/gainstat_backend/pkg/handler"
    "github.com/marianna961/gainstat_backend"
)

func main() {
	// Подключение к базе данных
	connStr := "user=gainstat_user password=qwerty dbname=gainstat sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("error occurred while connecting to the database: %s", err.Error())
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf("error occurred while pinging the database: %s", err.Error())
	}

	fmt.Println("Successfully connected to PostgreSQL!")

	// Инициализация маршрутов и запуск HTTP-сервера
	handlers := new(handler.Handler)
	srv := new(gainstat.Server)
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("error occurred while running the HTTP server: %s", err.Error())
	}
}
