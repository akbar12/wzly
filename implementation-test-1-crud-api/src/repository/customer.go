package repository

import (
	"context"
	"database/sql"
	"log"

	"pos-api/src/model"
	"pos-api/src/util"

	"github.com/doug-martin/goqu/v9"
)

type CustomerRepoIface interface {
	ListCustomer(ctx context.Context, tx *sql.Tx, param ParamListCustomer) (list []model.CustomerModel, total int, err error)
}

type CustomerRepo struct {
	db *sql.DB
}

func InitCustomerRepo(db *sql.DB) CustomerRepoIface {
	return &CustomerRepo{
		db: db,
	}
}

type ParamListCustomer struct {
	CustomerID   []int64
	CustomerName string
	Pagin        util.PaginParam
}

func (i *CustomerRepo) ListCustomer(ctx context.Context, tx *sql.Tx, param ParamListCustomer) (list []model.CustomerModel, total int, err error) {
	dataset := util.SqlDialect.
		Select(
			goqu.I("c.id").As("id"),
			goqu.I("c.customer_no").As("customer_no"),
			goqu.I("c.customer_name").As("customer_name"),
			goqu.I("c.detail_address").As("detail_address"),
			goqu.I("c.created_by").As("created_by"),
			goqu.I("c.created_date").As("created_date"),
		).
		From(goqu.T("customer").As("c")).
		GroupBy(goqu.I("c.id")).
		Limit(param.Pagin.GetLimit()).
		Offset(param.Pagin.GetOffset()).
		Order(
			goqu.I(param.Pagin.GetOrderBy("c.id")).Desc(),
		)

	if len(param.CustomerID) > 0 {
		dataset = dataset.Where(
			goqu.I("c.id").In(param.CustomerID),
		)
	}

	if param.CustomerName != "" {
		dataset = dataset.Where(
			goqu.I("c.customer_name").ILike("%" + param.CustomerName + "%"),
		)
	}

	query, params, err := dataset.Prepared(true).ToSQL()
	if err != nil {
		log.Println("Prepared ", err)
		return
	}

	var rows *sql.Rows
	if tx != nil {
		rows, err = tx.QueryContext(ctx, query, params...)
	} else {
		rows, err = i.db.QueryContext(ctx, query, params...)
	}

	if err != nil && err != sql.ErrNoRows {
		log.Println("ErrNoRows", err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var c model.CustomerModel
		if err = rows.Scan(
			&c.ID,
			&c.CustomerNo,
			&c.CustomerName,
			&c.DetailAddress,
			&c.CreatedBy,
			&c.CreatedDate,
		); err != nil {
			log.Println(err)
			return
		}
		list = append(list, c)
	}

	if err = rows.Err(); err != nil && err != sql.ErrNoRows {
		log.Println("rows.Err ", err)
		return
	}

	return
}
