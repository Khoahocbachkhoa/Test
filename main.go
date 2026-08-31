package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"project/internal/handler"
	"project/internal/repository"
	"project/internal/service"

	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
)

func main() {
	// Load .env
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Cannot read .env:", err)
	}

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
	taskRepo := repository.NewTaskRepository(conn)

	// create service
	taskService := service.NewTaskService(taskRepo)

	// create handler
	taskHandler := handler.NewTaskHandler(taskService)

	// register route
	mux := http.NewServeMux()

	mux.HandleFunc("GET /health", taskHandler.Health)
	mux.HandleFunc("GET /tasks", taskHandler.GetTasks)
	mux.HandleFunc("POST /tasks", taskHandler.CreateTask)
	mux.HandleFunc("GET /tasks/{id}", taskHandler.GetTaskByID)
	mux.HandleFunc("PUT /tasks/{id}", taskHandler.UpdateTask)
	mux.HandleFunc("DELETE /tasks/{id}", taskHandler.DeleteTask)

	log.Println("Server listening on :8080")

	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}
