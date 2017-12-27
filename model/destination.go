package model

import "github.com/jmoiron/sqlx"

// Destination seems a NENGA ATENA
type Destination atenaAbstract

// All of destinations data
func (Destination) All(db *sqlx.DB) (*[]Destination, error) {
	x := new([]Destination)
	q := `
select *
from ` + AtenaTableName + ` a
where a.is_sender = false
and a.deleted = false
order by a.id desc`
	err := db.Select(x, q)
	return x, err
}
