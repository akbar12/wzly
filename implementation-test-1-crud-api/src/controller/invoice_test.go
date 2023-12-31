package controller

import (
	"bytes"
	"context"
	"database/sql"
	"io"
	"net/http/httptest"
	"pos-api/src/model"
	"pos-api/src/repository"
	"pos-api/src/util"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockInvoiceRepo struct {
	mock.Mock
}

func (m *MockInvoiceRepo) BeginTx(ctx context.Context) (*sql.Tx, error) {
	args := m.Called(ctx)
	return args.Get(0).(*sql.Tx), args.Error(1)
}

func (m *MockInvoiceRepo) ListInvoice(ctx context.Context, param *repository.ListInvoiceParam) (list []model.InvoiceModel, total int, err error) {
	args := m.Called(ctx, param)
	return args.Get(0).([]model.InvoiceModel), args.Int(1), args.Error(2)
}

func (m *MockInvoiceRepo) DetailInvoice(ctx context.Context, id int) (model.InvoiceModel, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(model.InvoiceModel), args.Error(1)
}

func (m *MockInvoiceRepo) Insert(ctx context.Context, tx *sql.Tx, invoice *model.InvoiceModel) (err error) {
	args := m.Called(ctx, tx, invoice)
	return args.Error(0)
}

func (m *MockInvoiceRepo) Update(ctx context.Context, tx *sql.Tx, invoice *model.InvoiceModel) (err error) {
	args := m.Called(ctx, tx, invoice)
	return args.Error(0)
}

func TestResponseError401(t *testing.T) {

	for _, testCase := range []struct {
		RequestBody string

		ExpectedResponse string
		ExpectedHttpCode int
	}{
		{
			RequestBody: `{
				"INVOICE_NO":"",
				"ISSUE_DATE":"",
				"SUBJECT":"",
				"CUSTOMER_NAME":"",
				"DUE_DATE":""
			}`,

			ExpectedHttpCode: 200,
			ExpectedResponse: `{
				"success":true,
				"data":{
					"INVOICES":[
						{
							"ID":27,
							"INVOICE_NO":27,
							"ISSUED_DATE":"01/12/2023",
							"DUE_DATE":"01/12/2023",
							"STATUS":1,
							"STATUS_NAME":"Paid",
							"CUSTOMER_NAME":"Budi Susanto",
							"SUBJECT":"Food for Budi 123",
							"DETAIL_ADDRESS":"",
							"TOTAL_ITEMS":2,
							"SUB_TOTAL":105.81,
							"TAX":10.58,
							"GRAND_TOTAL":116.39,
							"ITEMS":null
						}
					],
					"TOTAL":0
				},
				"message":"successfully retrieve/save data"
			}`,
		},
	} {
		w := httptest.NewRecorder()
		r := httptest.NewRequest("GET", "/", bytes.NewReader([]byte(testCase.RequestBody)))

		mockInvoiceRepo := new(MockInvoiceRepo)
		invoiceController := InvoiceController{
			InvoiceRepo: mockInvoiceRepo,
		}

		issuedDate, _ := time.Parse(util.DateFormat, "01/12/2023")
		dueDate, _ := time.Parse(util.DateFormat, "01/12/2023")

		mockInvoiceRepo.On(
			"ListInvoice",
			context.Background(),
			&repository.ListInvoiceParam{},
		).Return(
			[]model.InvoiceModel{
				{
					ID:           27,
					InvoiceNo:    27,
					IssuedDate:   issuedDate,
					DueDate:      dueDate,
					Status:       1,
					CustomerName: "Budi Susanto",
					Subject:      "Food for Budi 123",
					TotalItems:   2,
					SubTotal:     105.81,
					Tax:          10.58,
					GrandTotal:   116.39,
				},
			},
			0,
			nil,
		)
		invoiceController.InvoiceList(w, r)
		mockInvoiceRepo.AssertExpectations(t)
		defer w.Result().Body.Close()
		data, err := io.ReadAll(w.Result().Body)
		respData := strings.ReplaceAll(string(data), "\n", "")
		assert.Nil(t, err)
		cleanExpResp := strings.Replace(strings.Replace(testCase.ExpectedResponse, "\t", "", -1), "\n", "", -1)
		assert.Equal(t, cleanExpResp, respData)
		assert.Equal(t, testCase.ExpectedHttpCode, w.Code)
	}
}
