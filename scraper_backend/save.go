package main

import (
    "context"
		"time"
		"backend/rod_utils"
		"fmt"

		"github.com/uptrace/bun"
)

type PageDataWrapper struct {
    bun.BaseModel `bun:"table:page_data_wrappers,alias:pdw"`

    ID              int64       `bun:"id,pk,autoincrement"`
    StartStation    string      `bun:"start_station"`
    EndStation      string      `bun:"end_station"`
    RetrievalTime   time.Time   `bun:"retrieval_time"`
    
    PageDataResults []rod_utils.PageData  `bun:"rel:has-many,join:id=wrapper_id"` 
}


func SaveToDB(db *bun.DB, wrappers []PageDataWrapper) error {
    ctx := context.Background()

    // 1. Loop through every route we scraped
    for _, wrapper := range wrappers {
        
        // Start a transaction for this route
        err := db.RunInTx(ctx, nil, func(ctx context.Context, tx bun.Tx) error {
            
            // A. Insert the Parent (Wrapper)
            // returning the ID is handled automatically by Bun
            _, err := tx.NewInsert().Model(&wrapper).Exec(ctx)
            if err != nil {
                return err
            }

            // B. Prepare the children
            // We must assign the new Wrapper ID to all the children
            for i := range wrapper.PageDataResults {
                wrapper.PageDataResults[i].WrapperID = wrapper.ID
            }

            // C. Insert all children (Tickets) in one bulk query
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
            // We continue to the next route even if one fails
        }
    }
    return nil
}
