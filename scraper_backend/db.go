package main

import (
    "database/sql"
    "log"
    "os"
		"strings"

    "github.com/joho/godotenv"
    "github.com/uptrace/bun"
    "github.com/uptrace/bun/dialect/pgdialect"
    "github.com/uptrace/bun/driver/pgdriver"
)

func ConnectDB() *bun.DB {
    err := godotenv.Load("../.env")
    if err != nil {
        log.Println("Note: .env file not found, using system environment variables.")
    }

    dsn := os.Getenv("DATABASE_URL")

		// THIS WAS ADDED TO GET THIGNS TO WORK, BUT SHOULD PROB NOT BE LIKE THIS WHEN IN PROD
    if !strings.Contains(dsn, "sslmode") {
        if strings.Contains(dsn, "?") {
            dsn += "&sslmode=disable"
        } else {
            dsn += "?sslmode=disable"
        }
    } 
    sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))
    db := bun.NewDB(sqldb, pgdialect.New())

    return db
}
