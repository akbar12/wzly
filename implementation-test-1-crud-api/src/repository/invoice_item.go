package repository

import (
	"context"
	"database/sql"
	"log"

	"pos-api/src/model"
	"pos-api/src/util"

	"github.com/doug-martin/goqu/v9"
)

type InvoiceItemRepoIface interface {
	Insert(ctx context.Context, tx *sql.Tx, invoiceItems []model.InvoiceItemModel) (err error)
	Delete(ctx context.Context, tx *sql.Tx, byInvoiceID int64) (err error)
	List(ctx context.Context, byInvoiceID int64) (list []model.InvoiceItemModel, err error)
}

type InvoiceItemRepo struct {
	db *sql.DB
}

func InitInvoiceItemRepo(db *sql.DB) InvoiceItemRepoIface {
	return &InvoiceItemRepo{
		db: db,
	}
}

func (i *InvoiceItemRepo) Insert(ctx context.Context, tx *sql.Tx, invoiceItems []model.InvoiceItemModel) (err error) {
	query, values, err := util.SqlDialect.Insert(goqu.T("invoice_item")).Rows(invoiceItems).Prepared(true).ToSQL()
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

	_, err = tx.ExecContext(ctx, query, values...)
	if err != nil {
		util.Error(err, nil)
	}
	return
}

func (i *InvoiceItemRepo) Delete(ctx context.Context, tx *sql.Tx, byInvoiceID int64) (err error) {
	query, value, err := util.SqlDialect.Delete(goqu.T("invoice_item")).
		Where(goqu.Ex{
			"invoice_id": byInvoiceID,
		}).Prepared(true).ToSQL()
	if err != nil {
		log.Println("1 ", err)
		return err
	}

	_, err = tx.ExecContext(ctx, query, value...)
	if err != nil {
		log.Println("2 ", err)
		return err
	}
	return nil
}

func (i *InvoiceItemRepo) List(ctx context.Context, byInvoiceID int64) (list []model.InvoiceItemModel, err error) {
	dataset := util.SqlDialect.
		Select(
			goqu.I("ii.invoice_item_id").As("invoice_item_id"),
			goqu.I("ii.invoice_id").As("invoice_id"),
			goqu.I("ii.item_id").As("item_id"),
			goqu.I("i.item_name").As("item_name"),
			goqu.I("i.item_type").As("item_type"),
			goqu.I("ii.qty").As("qty"),
			goqu.I("ii.unit_price").As("unit_price"),
			goqu.I("ii.amount").As("amount"),
		).
		From(goqu.T("invoice_item").As("ii")).
		LeftJoin(goqu.T("item").As("i"),
			goqu.On(
				goqu.I("i.id").Eq(goqu.I("ii.item_id")),
			),
		).
		GroupBy("ii.invoice_item_id").
		Order(
			goqu.I("ii.invoice_item_id").Asc(),
		)

	if byInvoiceID != 0 {
		dataset = dataset.Where(
			goqu.I("ii.invoice_id").Eq(byInvoiceID),
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
		var invc model.InvoiceItemModel
		if err = rows.Scan(
			&invc.InvoiceItemID,
			&invc.InvoiceID,
			&invc.ItemID,
			&invc.ItemName,
			&invc.ItemType,
			&invc.Qty,
			&invc.UnitPrice,
			&invc.Amount,
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
