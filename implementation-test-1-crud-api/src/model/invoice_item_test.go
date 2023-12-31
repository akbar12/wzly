package model

import (
	"fmt"
	"pos-api/src/util"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCreateInvoiceItem(t *testing.T) {
	now := time.Now()
	for _, testCase := range []struct {
		Invoice  *InvoiceModel
		Item     *ItemModel
		Qty      int64
		Username string
		Now      time.Time

		ExpectedInvoiceItem *InvoiceItemModel
		ExpectedErr         error
	}{
		{
			Invoice: &InvoiceModel{ID: 1},
			Item: &ItemModel{
				ID:        5,
				UnitPrice: 1000,
			},
			Qty:      5,
			Username: "Admin A",
			Now:      now,

			ExpectedInvoiceItem: &InvoiceItemModel{
				InvoiceID:   1,
				ItemID:      5,
				Qty:         5,
				UnitPrice:   1000,
				Amount:      5000,
				CreatedBy:   "Admin A",
				CreatedDate: now,
			},
			ExpectedErr: nil,
		},
		{
			Item: &ItemModel{},

			ExpectedInvoiceItem: nil,
			ExpectedErr:         fmt.Errorf(util.ErrInvalidParameter.Error(), "Item(s)"),
		},
		{
			Item: &ItemModel{ID: 1},

			ExpectedInvoiceItem: nil,
			ExpectedErr:         fmt.Errorf(util.ErrInvalidParameter.Error(), "quantity"),
		},
	} {
		actualInvoiceItem, err := CreateInvoiceItem(
			testCase.Invoice,
			testCase.Item,
			testCase.Qty,
			testCase.Username,
			testCase.Now,
		)
		assert.Equal(t, testCase.ExpectedInvoiceItem, actualInvoiceItem, "TestCreateInvoiceItem")
		assert.Equal(t, testCase.ExpectedErr, err, "TestCreateInvoiceItem")
	}
}
