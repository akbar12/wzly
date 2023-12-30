package util

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
)

var (
	SuccessRetrieveData string = "successfully retrieve/save data"

	ErrInternalServerError error = errors.New("internal server error")
	ErrUnauthorized        error = errors.New("unauthorized")
	ErrInvalidRequestBody  error = errors.New("request body is invalid")
	ErrInvalidPathParam    error = errors.New("missing required path parameters")

	ErrInvalidParameter error = errors.New("invalid Paramater: %v")
)

type Pagination struct {
	Limit   int    `json:"LIMIT"`
	Page    int    `json:"PAGE"`
	OrderBy string `json:"ORDER_BY"`
}

func (p *Pagination) ParseFromBody() (param *PaginParam) {
	return &PaginParam{
		Limit:   uint(p.Limit),
		Page:    uint(p.Page),
		OrderBy: p.OrderBy,
	}
}

func DecodeJsonRequest(r *http.Request, dst interface{}) (err error) {
	err = json.NewDecoder(r.Body).Decode(dst)
	log.Println(err)
	return
}

func ResponseError401(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusUnauthorized)
	ErrorResponseBody(w, err)
}

func ResponseError400(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusBadRequest)
	ErrorResponseBody(w, err)
}

func ResponseError500(w http.ResponseWriter) {
	w.WriteHeader(http.StatusInternalServerError)
	ErrorResponseBody(w, ErrInternalServerError)
}

func ErrorResponseBody(w http.ResponseWriter, err error) {
	json.NewEncoder(w).Encode(ResponseBody{
		Success: false,
		Message: err.Error(),
	})
}

type ResponseBody struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

func ResponseSuccess200(w http.ResponseWriter, data interface{}, msg string) {
	json.NewEncoder(w).Encode(ResponseBody{
		Success: true,
		Data:    data,
		Message: msg,
	})
}
