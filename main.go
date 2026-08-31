package main

import (
	"context"
	"log"
	"os"
	"project/internal/repository"

	"github.com/jackc/pgx/v5"
)

func main() {
	// Connect database
	connStr := os.Getenv("DATABASE_URL")
	conn, err := pgx.Connect(context.Background(), connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close(context.Background())

	if err := conn.Ping(context.Background()); err != nil {
		log.Fatal(err)
	}

	// create repository
	repo := repository.NewTaskRepository(conn)

	// create service

}
