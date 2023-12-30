package model

import (
	"fmt"
	"time"

	"pos-api/src/util"

	"github.com/shopspring/decimal"
)

type InvoiceItemModel struct {
	InvoiceItemID int64     `db:"invoice_item_id"`
	InvoiceID     int64     `db:"invoice_id"`
	ItemID        int64     `db:"item_id"`
	ItemName      string    `db:"item_name" goqu:"skipinsert,skipupdate"`
	ItemType      string    `db:"item_type" goqu:"skipinsert,skipupdate"`
	Qty           int64     `db:"qty"`
	UnitPrice     float64   `db:"unit_price"`
	Amount        float64   `db:"amount"`
	CreatedBy     string    `db:"created_by"`
	CreatedDate   time.Time `db:"created_date"`
}

func CreateInvoiceItem(
	invoice *InvoiceModel, item *ItemModel, qty int64,
	username string, now time.Time,

) (*InvoiceItemModel, error) {
	if item.ID == 0 {
		return nil, fmt.Errorf(util.ErrInvalidParameter.Error(), "Item(s)")
	}
	if qty < 1 {
		return nil, fmt.Errorf(util.ErrInvalidParameter.Error(), "quantity")
	}

	amount, _ := decimal.NewFromInt(qty).Mul(decimal.NewFromFloat(item.UnitPrice)).Float64()

	return &InvoiceItemModel{
		InvoiceID:   invoice.ID,
		ItemID:      int64(item.ID),
		Qty:         qty,
		UnitPrice:   item.UnitPrice,
		Amount:      amount,
		CreatedBy:   username,
		CreatedDate: now,
	}, nil
}

func CreateBulkInvoiceItem(
	invoice *InvoiceModel,
	items []ItemModel,
	qtyMap map[int64]int64,
	now time.Time,
) (listInvoiceItem []InvoiceItemModel, err error) {
	var subtotal float64
	var totalQty int64
	for _, i := range items {
		qty, ok := qtyMap[int64(i.ID)]
		if !ok {
			continue
		}
		totalQty += qty
		subtotal, _ = decimal.NewFromFloat(subtotal).Add(decimal.NewFromInt(qty).Mul(decimal.NewFromFloat(i.UnitPrice))).Float64()
		var invoiceItem *InvoiceItemModel
		invoiceItem, err = CreateInvoiceItem(invoice, &i, qty, invoice.ModifiedBy, now)
		if err != nil {
			return
		}
		listInvoiceItem = append(listInvoiceItem, *invoiceItem)
	}

	invoice.TotalItems = int64(len(items))
	invoice.SubTotal = subtotal
	invoice.SetTax()
	invoice.SetGrandTotal()
	return
}
