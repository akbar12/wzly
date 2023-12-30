package model

import "time"

type CustomerModel struct {
	ID            int64     `db:"id"`
	CustomerNo    string    `db:"customer_no"`
	CustomerName  string    `db:"customer_name"`
	DetailAddress string    `db:"detail_address"`
	CreatedDate   time.Time `db:"created_date"`
	CreatedBy     string    `db:"created_by"`
	ModifiedDate  time.Time `db:"modified_date"`
	ModifiedBy    string    `db:"modified_by"`
}
