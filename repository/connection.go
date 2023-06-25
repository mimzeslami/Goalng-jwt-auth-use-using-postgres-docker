package repository

import (
	"os"

	"github.com/go-pg/pg"
)

type PostgresImpl struct {
	db *pg.DB
}

func NewPostgresRepo() *PostgresImpl {
	db := pg.Connect(&pg.Options{
		Addr:     os.Getenv("DB_HOST") + ":" + os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Database: os.Getenv("DB_NAME"),
	})

	return &PostgresImpl{db}
}
