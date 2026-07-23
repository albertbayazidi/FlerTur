package types

import (
	"time"

	"github.com/uptrace/bun"
)

type PageData struct {
    bun.BaseModel `bun:"table:page_data_results,alias:pdr"`

    ID             int64    `bun:"id,pk,autoincrement"`
    WrapperID      int64    `bun:"wrapper_id"` // The Foreign Key
    Duration       string   `bun:"duration"`
    StartTime      string   `bun:"start_time"`
    Price          int      `bun:"price"`
    NumberOfTrains int      `bun:"number_of_trains"`
    TrainIds       []string `bun:"train_ids,array"` // 'array' handles the Postgres text[] type
    URL            string   `bun:"url"`
}

type PageDataWrapper struct {
    bun.BaseModel `bun:"table:page_data_wrappers,alias:pdw"`

    ID              int64       `bun:"id,pk,autoincrement"`
    StartStation    string      `bun:"start_station"`
    EndStation      string      `bun:"end_station"`
    RetrievalTime   time.Time   `bun:"retrieval_time"`
    
    PageDataResults []PageData  `bun:"rel:has-many,join:id=wrapper_id"` 
}

type Route struct {
	Start string
	End   string
}

type UrlAndMetaData struct {
	Date time.Time
	StationLatLonPair string
	Url string
}
