package request

import (
	"pos-api/src/model"
	"pos-api/src/repository"
	"pos-api/src/util"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestToQueryRepo(t *testing.T) {

	var totalItems int64 = 10
	var status int = 0
	issueDate10, _ := time.Parse(util.DateFormat, "10/10/2023")
	dueDate15, _ := time.Parse(util.DateFormat, "15/10/2023")

	for _, testCase := range []struct {
		ListRequest *InvoiceListRequest

		ExpectedListParam *repository.ListInvoiceParam
	}{
		{
			ListRequest: &InvoiceListRequest{
				InvoiceID:    "001",
				IssueDate:    "10/10/2023",
				Subject:      "Test Subject",
				TotalItems:   &totalItems,
				CustomerName: "John Doe",
				DueDate:      "15/10/2023",
				Status:       &status,
				Pagination: util.Pagination{
					Limit:   5,
					Page:    1,
					OrderBy: "subject",
				},
			},

			ExpectedListParam: &repository.ListInvoiceParam{
				InvoiceID:    "001",
				IssueDate:    issueDate10,
				Subject:      "Test Subject",
				TotalItems:   &totalItems,
				CustomerName: "John Doe",
				DueDate:      dueDate15,
				Status:       &status,
				Pagin: util.PaginParam{
					Limit:   5,
					Page:    1,
					OrderBy: "subject",
				},
			},
		},
	} {
		assert.Equal(t, *testCase.ExpectedListParam, *testCase.ListRequest.ToQueryRepo(), "TestToQueryRepo")
	}
}

func TestGetInvoice(t *testing.T) {
	now := time.Now()
	for _, testCase := range []struct {
		Invoice      *model.InvoiceModel
		InvoiceItems []model.InvoiceItemModel

		ExpectedInvoiceResponse InvoiceResponse
	}{
		{
			Invoice: &model.InvoiceModel{
				ID:            1,
				InvoiceNo:     001,
				IssuedDate:    now,
				DueDate:       now,
				Status:        int(model.UnpaidStatus),
				CustomerName:  "John Doe",
				Subject:       "John & Jhonson",
				DetailAddress: "Tangerang Selatan",
				TotalItems:    5,
				SubTotal:      4000,
				Tax:           400,
				GrandTotal:    4400,
			},
			InvoiceItems: []model.InvoiceItemModel{
				{
					InvoiceItemID: 1,
					InvoiceID:     1,
					ItemID:        3,
					ItemName:      "Soap",
					ItemType:      "self-care",
					Qty:           2,
					UnitPrice:     500,
					Amount:        1000,
				},
				{
					InvoiceItemID: 2,
					InvoiceID:     1,
					ItemID:        7,
					ItemName:      "toothpaste",
					ItemType:      "self-care",
					Qty:           3,
					UnitPrice:     1000,
					Amount:        3000,
				},
			},
			ExpectedInvoiceResponse: InvoiceResponse{
				ID:            1,
				InvoiceNo:     001,
				IssuedDate:    now.Format(util.DateFormat),
				DueDate:       now.Format(util.DateFormat),
				Status:        int(model.UnpaidStatus),
				StatusName:    string(model.UnpaidStatusName),
				CustomerName:  "John Doe",
				Subject:       "John & Jhonson",
				DetailAddress: "Tangerang Selatan",
				TotalItems:    5,
				SubTotal:      4000,
				Tax:           400,
				GrandTotal:    4400,
				InvoiceItem: []InvoiceItemReponse{
					{
						InvoiceItemID: 1,
						InvoiceID:     1,
						ItemID:        3,
						ItemName:      "Soap",
						ItemType:      "self-care",
						Qty:           2,
						UnitPrice:     500,
						Amount:        1000,
					},
					{
						InvoiceItemID: 2,
						InvoiceID:     1,
						ItemID:        7,
						ItemName:      "toothpaste",
						ItemType:      "self-care",
						Qty:           3,
						UnitPrice:     1000,
						Amount:        3000,
					},
				},
			},
		},
	} {
		actualInvoiceResponse := GetInvoice(testCase.Invoice, testCase.InvoiceItems)
		assert.Equal(t, testCase.ExpectedInvoiceResponse, actualInvoiceResponse, "TestGetInvoice")
	}
}
