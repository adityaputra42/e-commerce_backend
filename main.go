package main

import (
	"context"
	"log"

	db "github.com/adityaputra42/e-commerce_backend/db/sqlc"
	"github.com/adityaputra42/e-commerce_backend/utils"
	"github.com/jackc/pgx/v5"
)

func main() {
	config, err := utils.LoadConfig(".")
	if err != nil {
		log.Fatal("Cannot load config", err)
	}
	conn, err := pgx.Connect(context.Background(), config.DbSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}
	db.NewStore(conn)
}
