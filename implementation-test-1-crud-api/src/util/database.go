package util

import (
	"database/sql"
	"embed"
	"time"

	"github.com/doug-martin/goqu/v9"
	_ "github.com/doug-martin/goqu/v9/dialect/mysql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	mysqlMigrate "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/golang-migrate/migrate/v4/source/iofs"
)

var SqlDialect = goqu.Dialect("mysql")

//go:embed migrations/*.sql
var fs embed.FS

func InitDB() (db *sql.DB) {
	db, err := sql.Open("mysql", GetConfigString("db.conn_str"))
	if err != nil {
		panic(err)
	}
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(100)
	db.SetMaxIdleConns(100)
	dbMigrate(db)
	return
}

func dbMigrate(db *sql.DB) {
	driver, err := mysqlMigrate.WithInstance(db, &mysqlMigrate.Config{})
	if err != nil {
		panic(err)
	}

	d, err := iofs.New(fs, "migrations")
	if err != nil {
		panic(err)
	}

	m, err := migrate.NewWithInstance("iofs", d, "mysql", driver)
	if err != nil {
		panic(err)
	}

	err = m.Up()
	if err != nil {
		panic(err)
	}
}
