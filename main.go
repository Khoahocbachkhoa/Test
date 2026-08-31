package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"project/internal/model"
	"strconv"
	"time"

	"github.com/jackc/pgx/v5"
)

var tasks = []model.Task{
	{
		ID:          10,
		Title:       "Learning Docker",
		Description: "Image, container, dockerfile, volume, ...",
		Status:      "Ongoing",
	},
	{
		ID:          20,
		Title:       "Learning Linux",
		Description: "Bash scripting, user management, file permisson, ...",
		Status:      "Finished",
	},
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "OK")
}

func exampleHandler(w http.ResponseWriter, r *http.Request) {
	task1 := model.Task{
		ID:          10,
		Title:       "Learning Docker",
		Description: "Image, container, dockerfile, volume, ...",
		Status:      "Ongoing",
	}

	task2 := model.Task{
		ID:          20,
		Title:       "Learning Linux",
		Description: "Bash scripting, user management, file permisson, ...",
		Status:      "Finished",
	}

	tasks := []model.Task{task1, task2}

	jsonTask, _ := json.Marshal(tasks) // Ignore error

	// Return json array
	fmt.Fprintf(w, "%s", jsonTask)
}

func getTasks(w http.ResponseWriter, r *http.Request) {
	jsonData, _ := json.Marshal(tasks)

	fmt.Fprintf(w, "%s", jsonData)
}

func getTaskByID(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	// find task by Id
	var res model.Task
	var found bool

	for _, task := range tasks {
		if strconv.Itoa(task.ID) == id {
			res = task
			found = true
			break
		}
	}

	if !found {
		fmt.Fprintf(w, "Not found!")
	} else {
		json.NewEncoder(w).Encode(res)
	}
}

func main_() {
	http.HandleFunc("/health", healthHandler)
	http.HandleFunc("/example", exampleHandler)
	http.HandleFunc("GET /tasks", getTasks)
	http.HandleFunc("GET /task/{id}", getTaskByID)

	fmt.Println("Server is listening on :8080")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

func main1() {
	connStr := os.Getenv("DATABASE_URL")

	fmt.Println(connStr)

	ctx := context.Background()
	conn, err := pgx.Connect(ctx, connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close(ctx)

	// Verify connect
	err = conn.Ping(ctx)
	if err != nil {
		log.Fatal(err)
	}

	// Test query
	rows, err := conn.Query(ctx, "select * from tasks")

	if err != nil {
		fmt.Fprintf(os.Stderr, "Query failed: %v", err)
	}

	defer rows.Close()

	fmt.Println(rows) //?

	for rows.Next() {
		var (
			id          int
			title       string
			description string
			status      string
			created     time.Time
			updated     time.Time
		)

		err := rows.Scan(
			&id,
			&title,
			&description,
			&status,
			&created,
			&updated,
		)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(id, title, description, status, created, updated)
	}

	fmt.Println("OKe")
}

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

}
