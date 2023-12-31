package repository

import (
	"context"
	"database/sql"

	"pos-api/src/model"
	"pos-api/src/util"

	"github.com/doug-martin/goqu/v9"
)

type ItemRepoIface interface {
	ListItem(ctx context.Context, param ParamListItem) (list []model.ItemModel, total int, err error)
}

type ItemRepo struct {
	db *sql.DB
}

func InitItemRepo(db *sql.DB) ItemRepoIface {
	return &ItemRepo{
		db: db,
	}
}

type ParamListItem struct {
	ItemID   []int64
	ItemName string
	ItemType string
	Pagin    util.PaginParam
}

func (i *ItemRepo) ListItem(ctx context.Context, param ParamListItem) (list []model.ItemModel, total int, err error) {
	dataset := util.SqlDialect.
		Select(
			goqu.I("i.id").As("id"),
			goqu.I("i.item_no").As("item_no"),
			goqu.I("i.item_name").As("item_name"),
			goqu.I("i.item_type").As("item_type"),
			goqu.I("i.unit_price").As("unit_price"),
			goqu.I("i.created_by").As("created_by"),
			goqu.I("i.created_date").As("created_date"),
		).
		From(goqu.T("item").As("i")).
		GroupBy(goqu.I("i.id")).
		Limit(param.Pagin.GetLimit()).
		Offset(param.Pagin.GetOffset()).
		Order(
			goqu.I(param.Pagin.GetOrderBy("i.id")).Desc(),
		)

	if len(param.ItemID) > 0 {
		dataset = dataset.Where(
			goqu.I("i.id").In(param.ItemID),
		)
	}

	if param.ItemName != "" {
		dataset = dataset.Where(
			goqu.I("i.item_name").ILike("%" + param.ItemName + "%"),
		)
	}

	if param.ItemType != "" {
		dataset = dataset.Where(
			goqu.I("i.item_type").ILike("%" + param.ItemType + "%"),
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
		var item model.ItemModel
		if err = rows.Scan(
			&item.ID,
			&item.ItemNo,
			&item.ItemName,
			&item.ItemType,
			&item.UnitPrice,
			&item.CreatedBy,
			&item.CreatedDate,
		); err != nil {
			util.Error(err, nil)
			return
		}
		list = append(list, item)
	}

	if err = rows.Err(); err != nil && err != sql.ErrNoRows {
		util.Error(err, nil)
		return
	}

	return
}
