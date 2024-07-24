package main

import (
	"context"
	"database/sql"
	_ "embed"
	"fmt"
	"log"
	"os"

	"github.com/google/uuid"
	"github.com/uoya/todo/tododb"

	_ "github.com/lib/pq"
)

type Todo struct {
	ID          uuid.UUID
	Name        string
	Description string
	Done        bool
}

func main() {
	err := run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

func run() error {
	ctx := context.Background()
	db, err := sql.Open("postgres", "host=localhost port=15432 user=postgres password=postgres dbname=postgres sslmode=disable")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
	}

	defer db.Close()

	queries := tododb.New(db)

	todos, err := queries.QueryTodos(ctx)
	log.Println(todos, err)
	return err
}
