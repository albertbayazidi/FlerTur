package main

import (
	"context"
	"database/sql"
	"fmt"
	"os"

	"backend/types"

	"github.com/joho/godotenv"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

func ConnectDB() *bun.DB {
	_ = godotenv.Load(".env", "../../.env")

	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		dsn = "postgres://USER:PASSWORD@localhost:5432/DB?sslmode=disable"
	}

	sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))
	db := bun.NewDB(sqldb, pgdialect.New())

	return db
}

func SaveToDB(db *bun.DB, wrappers []types.PageDataWrapper) error {
	ctx := context.Background()

	for _, wrapper := range wrappers {

		err := db.RunInTx(ctx, nil, func(ctx context.Context, tx bun.Tx) error {

			_, err := tx.NewInsert().Model(&wrapper).Exec(ctx)
			if err != nil {
				return err
			}

			for i := range wrapper.PageDataResults {
				wrapper.PageDataResults[i].WrapperID = wrapper.ID
			}

			if len(wrapper.PageDataResults) > 0 {
				_, err = tx.NewInsert().Model(&wrapper.PageDataResults).Exec(ctx)
				if err != nil {
					return err
				}
			}
			return nil
		})
		if err != nil {
			fmt.Printf("Error saving route %s->%s: %v\n", wrapper.StartStation, wrapper.EndStation, err)
		}
	}
	return nil
}
