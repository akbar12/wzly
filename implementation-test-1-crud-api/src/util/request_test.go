package util

import (
	"io"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestResponseError401(t *testing.T) {
	for _, testCase := range []struct {
		Err error

		ExpectedResponse string
		ExpectedHttpCode int
	}{
		{
			Err:              ErrUnauthorized,
			ExpectedResponse: `{"success":false,"data":null,"message":"unauthorized"}`,
			ExpectedHttpCode: 401,
		},
	} {
		w := httptest.NewRecorder()
		ResponseError401(w, testCase.Err)
		defer w.Result().Body.Close()
		data, err := io.ReadAll(w.Result().Body)
		respData := strings.ReplaceAll(string(data), "\n", "")
		assert.Nil(t, err)
		assert.Equal(t, testCase.ExpectedResponse, respData)
		assert.Equal(t, testCase.ExpectedHttpCode, w.Code)
	}
}

func TestResponseSuccess200(t *testing.T) {
	for _, testCase := range []struct {
		Data    string
		Message string

		ExpectedResponse string
		ExpectedHttpCode int
	}{
		{
			Data:    `invoice_id: 1`,
			Message: `successfully insert Invoice.`,

			ExpectedResponse: `{"success":false,"data":{"invoice_id":1},"message":"successfully insert Invoice."}`,
			ExpectedHttpCode: 200,
		},
	} {
		w := httptest.NewRecorder()
		ResponseSuccess200(w, testCase.Data, testCase.Message)
		defer w.Result().Body.Close()
		data, err := io.ReadAll(w.Result().Body)
		respData := strings.ReplaceAll(string(data), "\n", "")
		assert.Nil(t, err)
		assert.Equal(t, testCase.ExpectedResponse, respData)
		assert.Equal(t, testCase.ExpectedHttpCode, w.Code)
	}
}

func TestParseFromBody(t *testing.T) {
	for _, testCase := range []struct {
		Pagination *Pagination

		PaginParam *PaginParam
	}{
		{
			Pagination: &Pagination{
				Limit:   100,
				Page:    1,
				OrderBy: "subject",
			},
			PaginParam: &PaginParam{
				Limit:   uint(100),
				Page:    uint(1),
				OrderBy: "subject",
			},
		},
	} {
		assert.Equal(t, *testCase.Pagination.ParseFromBody(), *testCase.PaginParam)
	}
}
