package main

import (
	"context"
	"fmt"

	"backend/types"

	"github.com/uptrace/bun"
)

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
