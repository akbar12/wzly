package model

import (
	"fmt"
	"html"
	"time"

	"pos-api/src/util"

	"github.com/shopspring/decimal"
)

type InvoiceModel struct {
	ID            int64     `db:"id"`
	InvoiceNo     int64     `db:"invoice_no"`
	IssuedDate    time.Time `db:"issued_date"`
	DueDate       time.Time `db:"due_date"`
	Status        int       `db:"status"`
	CustomerID    int64     `db:"customer_id"`
	CustomerName  string    `db:"customer_name" goqu:"skipinsert,skipupdate"`
	DetailAddress string    `db:"detail_address" goqu:"skipinsert,skipupdate"`
	Subject       string    `db:"subject"`
	TotalItems    int64     `db:"total_item"`
	SubTotal      float64   `db:"sub_total"`
	Tax           float64   `db:"tax"`
	GrandTotal    float64   `db:"grand_total"`
	CreatedDate   time.Time `db:"created_date"`
	CreatedBy     string    `db:"created_by"`
	ModifiedDate  time.Time `db:"modified_date"`
	ModifiedBy    string    `db:"modified_by"`
}

type Status int
type StatusName string

const (
	UnpaidStatus Status = 0
	PaidStatus   Status = 1

	UnpaidStatusName StatusName = "Unpaid"
	PaidStatusName   StatusName = "Paid"
)

func (i *InvoiceModel) GetStatusName() string {
	switch i.Status {
	case int(UnpaidStatus):
		return string(UnpaidStatusName)
	case int(PaidStatus):
		return string(PaidStatusName)
	default:
		return string(UnpaidStatusName)
	}
}

func (i *InvoiceModel) SetTax() {
	i.Tax, _ = (decimal.NewFromFloat(i.SubTotal).Mul(decimal.NewFromFloat(0.1))).Float64()
}

func (i *InvoiceModel) SetGrandTotal() {
	i.GrandTotal, _ = decimal.NewFromFloat(i.SubTotal).Add(decimal.NewFromFloat(i.Tax)).Float64()
}

type CreateInvoiceParam struct {
	IssuedDate time.Time
	DueDate    time.Time
	Status     int
	CustomerID int64
	Subject    string
	GrandTotal float64
	Username   string
	Now        time.Time

	Customer *CustomerModel
	Items    []ItemModel
	ItemsQty map[string]int64
}

func (p *CreateInvoiceParam) Validate() error {
	if p.IssuedDate.IsZero() {
		return fmt.Errorf(util.ErrInvalidParameter.Error(), "Issued Date")
	}

	if p.DueDate.IsZero() {
		return fmt.Errorf(util.ErrInvalidParameter.Error(), "Due Date")
	}

	if p.Status != int(PaidStatus) && p.Status != int(UnpaidStatus) {
		return fmt.Errorf(util.ErrInvalidParameter.Error(), "Status")
	}

	if p.Subject == "" {
		return fmt.Errorf(util.ErrInvalidParameter.Error(), "Subject")
	}

	if p.Customer == nil || p.Customer.ID == 0 {
		return fmt.Errorf(util.ErrInvalidParameter.Error(), "Customer")
	}

	return nil
}

func CreateInvoice(param *CreateInvoiceParam) (*InvoiceModel, error) {
	if err := param.Validate(); err != nil {
		return nil, err
	}
	invoice := &InvoiceModel{
		IssuedDate:   param.IssuedDate,
		DueDate:      param.DueDate,
		Status:       param.Status,
		CustomerID:   param.CustomerID,
		Subject:      html.EscapeString(param.Subject),
		CreatedDate:  param.Now,
		CreatedBy:    param.Username,
		ModifiedDate: param.Now,
		ModifiedBy:   param.Username,
	}
	return invoice, nil
}

func (i *InvoiceModel) Update(param *CreateInvoiceParam) error {
	if err := param.Validate(); err != nil {
		return err
	}
	i.IssuedDate = param.IssuedDate
	i.DueDate = param.DueDate
	i.Status = param.Status
	i.CustomerID = param.CustomerID
	i.Subject = html.EscapeString(param.Subject)
	i.ModifiedDate = param.Now
	i.ModifiedBy = param.Username

	return nil
}
