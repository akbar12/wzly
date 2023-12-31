package model

import (
	"fmt"
	"pos-api/src/util"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCreateInvoice(t *testing.T) {
	now := time.Now()
	for _, testCase := range []struct {
		CreateInvoiceParam *CreateInvoiceParam

		ExpectedInvoice *InvoiceModel
		ExpectedErr     error
	}{
		{
			CreateInvoiceParam: &CreateInvoiceParam{
				IssuedDate: now,
				DueDate:    now,
				Status:     int(UnpaidStatus),
				CustomerID: 7,
				Customer: &CustomerModel{
					ID: 7,
				},
				Subject:  "John & Jhonson",
				Username: "Admin A",
				Now:      now,
			},

			ExpectedInvoice: &InvoiceModel{
				IssuedDate:   now,
				DueDate:      now,
				Status:       int(UnpaidStatus),
				CustomerID:   7,
				Subject:      "John &amp; Jhonson",
				CreatedBy:    "Admin A",
				CreatedDate:  now,
				ModifiedBy:   "Admin A",
				ModifiedDate: now,
			},
			ExpectedErr: nil,
		},
		{
			CreateInvoiceParam: &CreateInvoiceParam{
				IssuedDate: now,
				DueDate:    now,
				Status:     int(UnpaidStatus),
				CustomerID: 7,
				Subject:    "John & Jhonson",
				Username:   "Admin A",
				Now:        now,
			},

			ExpectedInvoice: nil,
			ExpectedErr:     fmt.Errorf(util.ErrInvalidParameter.Error(), "Customer"),
		},
	} {
		actualInvoice, err := CreateInvoice(testCase.CreateInvoiceParam)
		if testCase.ExpectedInvoice != nil {
			assert.NotNil(t, actualInvoice, "TestCreateInvoice")
			if actualInvoice != nil {
				assert.Equal(t, *testCase.ExpectedInvoice, *actualInvoice, "TestCreateInvoice")
			}
		}
		assert.Equal(t, testCase.ExpectedErr, err, "TestCreateInvoice")
	}
}

func TestSetTax(t *testing.T) {
	for _, testCase := range []struct {
		Invoice *InvoiceModel

		ExpectedTax float64
	}{
		{
			Invoice: &InvoiceModel{
				SubTotal: 1000,
			},

			ExpectedTax: 100,
		},
	} {
		testCase.Invoice.SetTax()
		assert.Equal(t, testCase.ExpectedTax, testCase.Invoice.Tax, "TestSetTax")
	}
}

func TestSetGrandTotal(t *testing.T) {
	for _, testCase := range []struct {
		Invoice *InvoiceModel

		ExpectedGrandTotal float64
	}{
		{
			Invoice: &InvoiceModel{
				SubTotal: 1000,
				Tax:      100,
			},

			ExpectedGrandTotal: 1100,
		},
	} {
		testCase.Invoice.SetGrandTotal()
		assert.Equal(t, testCase.ExpectedGrandTotal, testCase.Invoice.GrandTotal, "TestSetGrandTotal")
	}
}
