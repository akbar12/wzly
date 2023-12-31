package controller

import (
	"net/http"
	"strconv"
	"time"

	"pos-api/src/model"
	"pos-api/src/repository"
	"pos-api/src/request"
	"pos-api/src/util"

	"github.com/gorilla/mux"
)

type InvoiceCtrlIface interface {
	InvoiceList(w http.ResponseWriter, r *http.Request)
	DetailInvoice(w http.ResponseWriter, r *http.Request)
	InvoiceCreate(w http.ResponseWriter, r *http.Request)
	InvoiceUpdate(w http.ResponseWriter, r *http.Request)
}

type InvoiceController struct {
	InvoiceRepo     repository.InvoiceRepoIface
	InvoiceItemRepo repository.InvoiceItemRepoIface
	ItemRepo        repository.ItemRepoIface
	CustomerRepo    repository.CustomerRepoIface
}

func InitInvoiceCtrl(
	invoiceRepo repository.InvoiceRepoIface,
	invoiceItemRepo repository.InvoiceItemRepoIface,
	itemRepo repository.ItemRepoIface,
	customerRepo repository.CustomerRepoIface) InvoiceCtrlIface {
	return &InvoiceController{
		InvoiceRepo:     invoiceRepo,
		InvoiceItemRepo: invoiceItemRepo,
		ItemRepo:        itemRepo,
		CustomerRepo:    customerRepo,
	}
}

func (i *InvoiceController) InvoiceList(w http.ResponseWriter, r *http.Request) {
	var err error
	body := request.InvoiceListRequest{}
	err = util.DecodeJsonRequest(r, &body)
	if err != nil {
		util.ResponseError400(w, util.ErrInvalidRequestBody)
		return
	}

	repoParam := body.ToQueryRepo()
	invoices, total, err := i.InvoiceRepo.ListInvoice(r.Context(), repoParam)
	if err != nil {
		util.ResponseError500(w)
		return
	}

	resp := request.ListInvoice(invoices, total)
	util.ResponseSuccess200(w, resp, util.SuccessRetrieveData)
}

func (i *InvoiceController) DetailInvoice(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	_id, ok := vars["id"]
	invoiceID, err := strconv.Atoi(_id)
	if !ok || err != nil {
		util.ResponseError400(w, util.ErrInvalidPathParam)
		return
	}

	invoice, err := i.InvoiceRepo.DetailInvoice(r.Context(), invoiceID)
	if err != nil {
		util.ResponseError500(w)
		return
	}

	if invoice.ID == 0 {
		util.ResponseError400(w, util.ErrInvalidPathParam)
		return
	}

	invoiceItem, err := i.InvoiceItemRepo.List(r.Context(), invoice.ID)
	if err != nil {
		util.ResponseError500(w)
		return
	}

	resp := request.GetInvoice(&invoice, invoiceItem)

	util.ResponseSuccess200(w, resp, util.SuccessRetrieveData)
}

func (i *InvoiceController) InvoiceCreate(w http.ResponseWriter, r *http.Request) {
	var err error
	body := request.InvoiceCreateRequest{}
	err = util.DecodeJsonRequest(r, &body)
	if err != nil {
		util.ResponseError400(w, util.ErrInvalidRequestBody)
		return
	}

	paramListCustomer := repository.ParamListCustomer{
		CustomerID: []int64{body.CustomerID},
	}
	customers, _, err := i.CustomerRepo.ListCustomer(r.Context(), nil, paramListCustomer)
	if err != nil {
		util.ResponseError500(w)
		return
	}
	if len(customers) == 0 {
		customers = append(customers, model.CustomerModel{})
	}

	createParam := body.ToCreateModel()
	createParam.Customer = &customers[0]
	createParam.Username = r.Context().Value("username").(string)
	invoice, err := model.CreateInvoice(createParam)
	if err != nil {
		util.ResponseError400(w, err)
		return
	}

	tx, err := i.InvoiceRepo.BeginTx(r.Context())
	if err != nil {
		util.ResponseError500(w)
		return
	}

	defer func() {
		if err != nil {
			tx.Rollback()
		}
		return
	}()

	err = i.InvoiceRepo.Insert(r.Context(), tx, invoice)
	if err != nil {
		util.ResponseError500(w)
		return
	}

	itemIDs, qtyMap := body.ItemData()
	paramListItem := repository.ParamListItem{
		ItemID: itemIDs,
	}
	items, _, err := i.ItemRepo.ListItem(r.Context(), paramListItem)
	if err != nil {
		util.ResponseError500(w)
		return
	}

	invoiceItems, err := model.CreateBulkInvoiceItem(invoice, items, qtyMap, time.Now())
	if err != nil {
		util.ResponseError400(w, err)
		return
	}

	err = i.InvoiceItemRepo.Insert(r.Context(), tx, invoiceItems)
	if err != nil {
		util.ResponseError500(w)
		return
	}

	err = i.InvoiceRepo.Update(r.Context(), tx, invoice)
	if err != nil {
		util.ResponseError500(w)
		return
	}

	err = tx.Commit()
	if err != nil {
		util.Error(err, nil)
		return
	}

	util.ResponseSuccess200(w, nil, util.SuccessRetrieveData)
}

func (i *InvoiceController) InvoiceUpdate(w http.ResponseWriter, r *http.Request) {
	var err error
	body := request.InvoiceCreateRequest{}
	err = util.DecodeJsonRequest(r, &body)
	if err != nil {
		util.ResponseError400(w, util.ErrInvalidRequestBody)
		return
	}

	vars := mux.Vars(r)
	_id, ok := vars["id"]
	invoiceID, err := strconv.Atoi(_id)
	if !ok || err != nil {
		util.ResponseError400(w, util.ErrInvalidPathParam)
		return
	}

	invoice, err := i.InvoiceRepo.DetailInvoice(r.Context(), invoiceID)
	if err != nil {
		util.ResponseError500(w)
		return
	}

	paramListCustomer := repository.ParamListCustomer{
		CustomerID: []int64{body.CustomerID},
	}
	customers, _, err := i.CustomerRepo.ListCustomer(r.Context(), nil, paramListCustomer)
	if err != nil {
		util.ResponseError500(w)
		return
	}
	if len(customers) == 0 {
		customers = append(customers, model.CustomerModel{})
	}

	updateParam := body.ToCreateModel()
	updateParam.Customer = &customers[0]
	updateParam.Username = r.Context().Value("username").(string)
	err = invoice.Update(updateParam)
	if err != nil {
		util.ResponseError400(w, err)
		return
	}

	tx, err := i.InvoiceRepo.BeginTx(r.Context())
	if err != nil {
		util.ResponseError500(w)
		return
	}

	defer func() {
		if err != nil {
			tx.Rollback()
		}
		return
	}()

	err = i.InvoiceRepo.Update(r.Context(), tx, &invoice)
	if err != nil {
		util.ResponseError500(w)
		return
	}

	err = i.InvoiceItemRepo.Delete(r.Context(), tx, invoice.ID)
	if err != nil {
		util.ResponseError500(w)
		return
	}

	itemIDs, qtyMap := body.ItemData()
	paramListItem := repository.ParamListItem{
		ItemID: itemIDs,
	}
	items, _, err := i.ItemRepo.ListItem(r.Context(), paramListItem)
	if err != nil {
		util.ResponseError500(w)
		return
	}

	invoiceItems, err := model.CreateBulkInvoiceItem(&invoice, items, qtyMap, time.Now())
	if err != nil {
		util.ResponseError400(w, err)
		return
	}

	err = i.InvoiceItemRepo.Insert(r.Context(), tx, invoiceItems)
	if err != nil {
		util.ResponseError500(w)
		return
	}

	err = i.InvoiceRepo.Update(r.Context(), tx, &invoice)
	if err != nil {
		util.ResponseError500(w)
		return
	}

	err = tx.Commit()
	if err != nil {
		util.Error(err, nil)
		return
	}

	util.ResponseSuccess200(w, nil, util.SuccessRetrieveData)
}
