package query

import (
	_ "github.com/duckdb/duckdb-go/v2"
	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

func Init(ddbFilepath string) error {
	d, err := sqlx.Open("duckdb", ddbFilepath)
	if err != nil {
		return err
	}
	if err := runMigrate(d); err != nil {
		return err
	}
	db = d
	return nil
}

func DB() *sqlx.DB {
	return db
}
func Close() error {
	return db.Close()
}

func Transaction(db *sqlx.DB, fn func(*sqlx.Tx) error) error {
	tx, err := db.Beginx()
	if err != nil {
		return err
	}
	defer tx.Rollback()
	if err := fn(tx); err != nil {
		return err
	}
	return tx.Commit()
}
