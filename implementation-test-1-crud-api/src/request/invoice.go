package request

import (
	"time"

	"pos-api/src/model"
	"pos-api/src/repository"
	"pos-api/src/util"
)

type InvoiceListRequest struct {
	InvoiceID    string `json:"INVOICE_NO"`
	IssueDate    string `json:"ISSUE_DATE"`
	Subject      string `json:"SUBJECT"`
	TotalItems   *int64 `json:"TOTAL_ITEMS"`
	CustomerName string `json:"CUSTOMER_NAME"`
	DueDate      string `json:"DUE_DATE"`
	Status       *int   `json:"STATUS"`
	util.Pagination
}

func (i *InvoiceListRequest) ToQueryRepo() *repository.ListInvoiceParam {
	issueDate, _ := time.Parse(util.DateFormat, i.IssueDate)
	dueDate, _ := time.Parse(util.DateFormat, i.DueDate)
	return &repository.ListInvoiceParam{
		InvoiceID:    i.InvoiceID,
		IssueDate:    issueDate,
		Subject:      i.Subject,
		TotalItems:   i.TotalItems,
		CustomerName: i.CustomerName,
		DueDate:      dueDate,
		Status:       i.Status,
		Pagin:        *i.Pagination.ParseFromBody(),
	}
}

type InvoiceResponse struct {
	ID            int64                `json:"ID"`
	InvoiceNo     int64                `json:"INVOICE_NO"`
	IssuedDate    string               `json:"ISSUED_DATE"`
	DueDate       string               `json:"DUE_DATE"`
	Status        int                  `json:"STATUS"`
	StatusName    string               `json:"STATUS_NAME"`
	CustomerName  string               `json:"CUSTOMER_NAME"`
	Subject       string               `json:"SUBJECT"`
	DetailAddress string               `json:"DETAIL_ADDRESS"`
	TotalItems    int64                `json:"TOTAL_ITEMS"`
	SubTotal      float64              `json:"SUB_TOTAL"`
	Tax           float64              `json:"TAX"`
	GrandTotal    float64              `json:"GRAND_TOTAL"`
	InvoiceItem   []InvoiceItemReponse `json:"ITEMS"`
}

func GetInvoice(invoiceModel *model.InvoiceModel, invoiceItems []model.InvoiceItemModel) InvoiceResponse {
	res := InvoiceResponse{
		ID:            invoiceModel.ID,
		InvoiceNo:     invoiceModel.InvoiceNo,
		IssuedDate:    invoiceModel.IssuedDate.Format(util.DateFormat),
		DueDate:       invoiceModel.DueDate.Format(util.DateFormat),
		Status:        invoiceModel.Status,
		StatusName:    invoiceModel.GetStatusName(),
		CustomerName:  invoiceModel.CustomerName,
		Subject:       invoiceModel.Subject,
		DetailAddress: invoiceModel.DetailAddress,
		TotalItems:    invoiceModel.TotalItems,
		SubTotal:      invoiceModel.SubTotal,
		Tax:           invoiceModel.Tax,
		GrandTotal:    invoiceModel.GrandTotal,
	}
	if len(invoiceItems) > 0 {
		res.InvoiceItem = ListInvoiceItem(invoiceItems)
	}

	return res
}

type InvoiceItemReponse struct {
	InvoiceItemID int64   `json:"invoice_item_id"`
	InvoiceID     int64   `json:"invoice_id"`
	ItemID        int64   `json:"item_id"`
	ItemName      string  `json:"item_name" goqu:"skipinsert,skipupdate"`
	ItemType      string  `json:"item_type" goqu:"skipinsert,skipupdate"`
	Qty           int64   `json:"qty"`
	UnitPrice     float64 `json:"unit_price"`
	Amount        float64 `json:"amount"`
}

func ListInvoiceItem(invoiceItems []model.InvoiceItemModel) (list []InvoiceItemReponse) {
	for _, iItem := range invoiceItems {
		list = append(list, InvoiceItemReponse{
			InvoiceItemID: iItem.InvoiceItemID,
			InvoiceID:     iItem.InvoiceID,
			ItemID:        iItem.ItemID,
			ItemName:      iItem.ItemName,
			ItemType:      iItem.ItemType,
			Qty:           iItem.Qty,
			UnitPrice:     iItem.UnitPrice,
			Amount:        iItem.Amount,
		})
	}
	return
}

type InvoiceListResponse struct {
	Invoices []InvoiceResponse `json:"INVOICES"`
	Total    int               `json:"TOTAL"`
}

func ListInvoice(invoices []model.InvoiceModel, total int) InvoiceListResponse {
	list := []InvoiceResponse{}
	for _, i := range invoices {
		list = append(list, GetInvoice(&i, []model.InvoiceItemModel{}))
	}
	return InvoiceListResponse{
		Invoices: list,
		Total:    total,
	}
}

type InvoiceCreateRequest struct {
	IssuedDate string               `json:"ISSUED_DATE"`
	DueDate    string               `json:"DUE_DATE"`
	Status     int                  `json:"STATUS"`
	CustomerID int64                `json:"CUSTOMER_ID"`
	Subject    string               `json:"SUBJECT"`
	Items      []InvoiceItemRequest `json:"ITEMS"`
}

type InvoiceItemRequest struct {
	ItemID   int64 `json:"ITEM_ID"`
	Quantity int64 `json:"QTY"`
}

func (i *InvoiceCreateRequest) ItemData() (itemIDs []int64, qtyMap map[int64]int64) {
	qtyMap = make(map[int64]int64, 0)
	for _, i := range i.Items {
		itemIDs = append(itemIDs, i.ItemID)
		qtyMap[i.ItemID] = qtyMap[i.ItemID] + i.Quantity
	}
	return
}

func (i *InvoiceCreateRequest) ToCreateModel() *model.CreateInvoiceParam {
	issuedDate, _ := time.Parse(util.DateFormat, i.IssuedDate)
	dueDate, _ := time.Parse(util.DateFormat, i.DueDate)
	return &model.CreateInvoiceParam{
		IssuedDate: issuedDate,
		DueDate:    dueDate,
		Status:     i.Status,
		CustomerID: i.CustomerID,
		Subject:    i.Subject,
		Now:        time.Now(),
	}
}
