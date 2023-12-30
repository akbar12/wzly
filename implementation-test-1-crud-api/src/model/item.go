package model

import (
	"time"
)

type ItemModel struct {
	ID           int64     `db:"id"`
	ItemNo       string    `db:"item_no"`
	ItemName     string    `db:"item_name"`
	ItemType     string    `db:"item_type"`
	UnitPrice    float64   `db:"unit_price"`
	CreatedDate  time.Time `db:"created_date"`
	CreatedBy    string    `db:"created_by"`
	ModifiedDate time.Time `db:"modified_date"`
	ModifiedBy   string    `db:"modified_by"`
}
