package main

import (
    "database/sql"
    "fmt"
    "github.com/lib/pq"
	"log"
	"github.com/marianna961/gainstat_backend/pkg/handlers"
    "github.com/marianna961/gainstat"
)

func main() {
	handlers := new(handler.Handler)
	srv := new(gain.Server)
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server %s", err.Error())
	}

}

func main_sql() {
    connStr := "user=gainstat_user password=qwerty dbname=gainstat sslmode=disable"
    db, err := sql.Open("postgres", connStr)
    if err != nil {
        panic(err)
    }

    err = db.Ping()
    if err != nil {
        panic(err)
    }

    fmt.Println("Successfully connected to PostgreSQL!")
}
