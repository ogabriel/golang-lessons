package main

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

// However, for many uses, such as a web application server, concurrency is required. To use a connection pool replace
// the import github.com/jackc/pgx/v5 with github.com/jackc/pgx/v5/pgxpool and connect with pgxpool.New() instead of
// pgx.Connect().
func main() {
	dbpool, err := pgxpool.New(context.Background(), os.Getenv("DATABASE_URL"))

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	defer dbpool.Close()

	var greeting string

	err = dbpool.QueryRow(context.Background(), "select 'Hello, world!'").Scan(&greeting)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(greeting)
}

// The *pgx.Conn returned by pgx.Connect() represents a single connection and is not concurrency safe. This is entirely
// appropriate for a simple command line example such as above. However, for many uses, such as a web application
// server, concurrency is required.
func mainOld() {
	// export DATABASE_URL=postgres://postgres:postgres@127.0.0.1:5432/pgx_getting_started
	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	defer conn.Close(context.Background())

	var greeting string

	err = conn.QueryRow(context.Background(), "select 'Hello, world!'").Scan(&greeting)

	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(greeting)
}
