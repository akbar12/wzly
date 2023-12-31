package repository

import (
	"context"
	"database/sql"
	"time"

	"pos-api/src/model"
	"pos-api/src/util"

	"github.com/doug-martin/goqu/v9"
)

type InvoiceRepoIface interface {
	BeginTx(context.Context) (*sql.Tx, error)
	ListInvoice(ctx context.Context, param *ListInvoiceParam) (list []model.InvoiceModel, total int, err error)
	DetailInvoice(ctx context.Context, id int) (model.InvoiceModel, error)
	Insert(ctx context.Context, tx *sql.Tx, invoice *model.InvoiceModel) (err error)
	Update(ctx context.Context, tx *sql.Tx, invoice *model.InvoiceModel) (err error)
}

type InvoiceRepo struct {
	db *sql.DB
}

func InitInvoiceRepo(db *sql.DB) InvoiceRepoIface { //2
	return &InvoiceRepo{
		db: db,
	}
}

type ListInvoiceParam struct {
	InvoiceID    string
	IssueDate    time.Time
	Subject      string
	TotalItems   *int64
	CustomerName string
	DueDate      time.Time
	Status       *int
	Pagin        util.PaginParam
}

func (i *InvoiceRepo) BeginTx(ctx context.Context) (*sql.Tx, error) {
	tx, err := i.db.BeginTx(ctx, nil)
	if err != nil {
		util.Error(err, nil)
	}
	return tx, err
}

func (i *InvoiceRepo) ListInvoice(ctx context.Context, param *ListInvoiceParam) (list []model.InvoiceModel, total int, err error) {
	dataset := util.SqlDialect.
		Select(
			goqu.I("i.id").As("id"),
			goqu.I("i.invoice_no").As("invoice_no"),
			goqu.I("i.subject").As("subject"),
			goqu.I("i.issued_date").As("issue_date"),
			goqu.I("i.due_date").As("due_date"),
			goqu.I("c.customer_name").As("customer_name"),
			goqu.I("i.total_item").As("total_item"),
			goqu.I("i.status").As("status"),
			goqu.I("i.sub_total").As("sub_total"),
			goqu.I("i.tax").As("tax"),
			goqu.I("i.grand_total").As("grand_total"),
		).
		From(goqu.T("invoice").As("i")).
		LeftJoin(goqu.T("customer").As("c"),
			goqu.On(
				goqu.I("c.id").Eq(goqu.I("i.customer_id")),
			),
		).
		GroupBy("i.id").
		Limit(param.Pagin.GetLimit()).
		Offset(param.Pagin.GetOffset()).
		Order(
			goqu.I(param.Pagin.GetOrderBy("id")).Desc(),
		)

	if param.InvoiceID != "" {
		dataset = dataset.Where(
			goqu.I("i.invoice_no").ILike("%" + param.InvoiceID + "%"),
		)
	}

	if !param.IssueDate.IsZero() {
		dataset = dataset.Where(
			goqu.I("i.issue_date").Eq(param.IssueDate.Format(util.DateFormat)),
		)
	}

	if param.Subject != "" {
		dataset = dataset.Where(
			goqu.I("i.subject").ILike("%" + param.Subject + "%"),
		)
	}

	if param.TotalItems != nil {
		dataset = dataset.Where(
			goqu.I("i.total_items").Eq(*param.TotalItems),
		)
	}

	if param.CustomerName != "" {
		dataset = dataset.Where(
			goqu.I("c.customer_name").ILike("%" + param.CustomerName + "%"),
		)
	}

	if !param.DueDate.IsZero() {
		dataset = dataset.Where(
			goqu.I("i.due_date").Eq(param.DueDate.Format(util.DateFormat)),
		)
	}

	if param.TotalItems != nil {
		dataset = dataset.Where(
			goqu.I("i.total_items").Eq(*param.TotalItems),
		)
	}

	query, params, err := dataset.Prepared(true).ToSQL()
	if err != nil {
		util.Error(err, nil)
		return
	}

	rows, err := i.db.QueryContext(ctx, query, params...)
	if err != nil && err != sql.ErrNoRows {
		util.Error(err, nil)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var invc model.InvoiceModel
		if err = rows.Scan(
			&invc.ID,
			&invc.InvoiceNo,
			&invc.Subject,
			&invc.IssuedDate,
			&invc.DueDate,
			&invc.CustomerName,
			&invc.TotalItems,
			&invc.Status,
			&invc.SubTotal,
			&invc.Tax,
			&invc.GrandTotal,
		); err != nil {
			util.Error(err, nil)
			return
		}
		list = append(list, invc)
	}

	if err = rows.Err(); err != nil && err != sql.ErrNoRows {
		util.Error(err, nil)
		return
	}

	return
}

func (i *InvoiceRepo) DetailInvoice(ctx context.Context, invoiceID int) (invoice model.InvoiceModel, err error) {
	dataset := util.SqlDialect.
		Select(
			goqu.I("i.id").As("id"),
			goqu.I("i.invoice_no").As("invoice_no"),
			goqu.I("i.subject").As("subject"),
			goqu.I("i.issued_date").As("issue_date"),
			goqu.I("i.due_date").As("due_date"),
			goqu.I("c.customer_name").As("customer_name"),
			goqu.I("c.detail_address").As("detail_address"),
			goqu.I("i.total_item").As("total_item"),
			goqu.I("i.sub_total").As("sub_total"),
			goqu.I("i.tax").As("tax"),
			goqu.I("i.grand_total").As("grand_total"),
			goqu.I("i.status").As("status"),
		).
		From(goqu.T("invoice").As("i")).
		LeftJoin(goqu.T("customer").As("c"),
			goqu.On(
				goqu.I("c.id").Eq(goqu.I("i.customer_id")),
			),
		)

	if invoiceID != 0 {
		dataset = dataset.Where(
			goqu.I("i.id").Eq(invoiceID),
		)
	}

	query, params, err := dataset.Prepared(true).ToSQL()
	if err != nil {
		util.Error(err, nil)
		return
	}

	err = i.db.QueryRow(query, params...).
		Scan(
			&invoice.ID,
			&invoice.InvoiceNo,
			&invoice.Subject,
			&invoice.IssuedDate,
			&invoice.DueDate,
			&invoice.CustomerName,
			&invoice.DetailAddress,
			&invoice.TotalItems,
			&invoice.SubTotal,
			&invoice.Tax,
			&invoice.GrandTotal,
			&invoice.Status,
		)
	if err != nil && err != sql.ErrNoRows {
		util.Error(err, nil)
		return
	}
	return
}

func (i *InvoiceRepo) Insert(ctx context.Context, tx *sql.Tx, invoice *model.InvoiceModel) (err error) {
	query, values, err := util.SqlDialect.Insert(goqu.T("invoice")).Rows(invoice).Prepared(true).ToSQL()
	if err != nil {
		util.Error(err, nil)
		return
	}

	if tx == nil {
		tx, err = i.db.BeginTx(ctx, nil)
		if err != nil {
			util.Error(err, nil)
			return
		}

		defer func() {
			if err != nil {
				tx.Rollback()
				return
			}
			tx.Commit()
		}()
	}

	result, err := tx.ExecContext(ctx, query, values...)
	if err != nil {
		util.Error(err, nil)
		return
	}

	id, err := result.LastInsertId()
	if err != nil {
		util.Error(err, nil)
		return
	}

	invoice.ID = id
	invoice.InvoiceNo = id

	err = i.Update(ctx, tx, invoice)
	if err != nil {
		util.Error(err, nil)
		return
	}

	return
}

func (i *InvoiceRepo) Update(ctx context.Context, tx *sql.Tx, invoice *model.InvoiceModel) (err error) {
	if tx == nil {
		tx, err = i.db.BeginTx(ctx, nil)
		if err != nil {
			util.Error(err, nil)
			return
		}

		defer func() {
			if err != nil {
				tx.Rollback()
				return
			}
			tx.Commit()
		}()
	}

	query, values, err := util.SqlDialect.
		Update(goqu.T("invoice")).
		Set(goqu.Ex{
			"invoice_no":    invoice.InvoiceNo,
			"issued_date":   invoice.IssuedDate,
			"due_date":      invoice.DueDate,
			"status":        invoice.Status,
			"customer_id":   invoice.CustomerID,
			"subject":       invoice.Subject,
			"total_item":    invoice.TotalItems,
			"sub_total":     invoice.SubTotal,
			"tax":           invoice.Tax,
			"grand_total":   invoice.GrandTotal,
			"modified_date": invoice.ModifiedDate,
			"modified_by":   invoice.ModifiedBy,
		}).
		Where(
			goqu.I("id").Eq(invoice.ID),
		).
		Prepared(true).ToSQL()

	if err != nil {
		util.Error(err, nil)
		return
	}

	_, err = tx.ExecContext(ctx, query, values...)
	if err != nil {
		util.Error(err, nil)
		return
	}

	return
}
