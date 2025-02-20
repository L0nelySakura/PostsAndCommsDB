package graph

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

import (
	"context"
	"fmt"

	"github.com/google/uuid"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Resolver struct {
	DB *pgxpool.Pool
}

func NewDatabase(connectionString string) *pgxpool.Pool {

	pool, err := pgxpool.New(context.Background(), connectionString)
	if err != nil {
		return nil
	}
	createTables(pool)
	return pool
}

func NewResolver(connectionString string) *Resolver {
	return &Resolver{
		DB: NewDatabase(connectionString),
	}
}

func createTables(db *pgxpool.Pool) {
	queries := []string{
		`
		CREATE TABLE IF NOT EXISTS posts (
			id UUID PRIMARY KEY,
			title TEXT NOT NULL,
			content TEXT NOT NULL,
			author TEXT NOT NULL,
			comments_enabled BOOLEAN NOT NULL
		);
		`,

		`
		CREATE TABLE IF NOT EXISTS comments (
			id UUID PRIMARY KEY,
			content TEXT NOT NULL,
			author TEXT NOT NULL,
			post_id UUID NOT NULL REFERENCES posts(id),
			parent_id UUID
		);
		`,

		`CREATE INDEX IF NOT EXISTS idx_comments_post ON comments(post_id)`,
		`CREATE INDEX IF NOT EXISTS idx_comments_parent ON comments(parent_id)`,
	}

	for _, q := range queries {
		_, err := db.Exec(context.Background(), q)
		if err != nil {
			fmt.Printf("Failed to create tables: %v", err)
		}
	}
}

func GenerateID() string {
	return uuid.New().String()
}
