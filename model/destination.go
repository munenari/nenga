package model

import "github.com/jmoiron/sqlx"

// Destination seems a NENGA ATENA
type Destination atenaAbstract

// All of destinations data
func (Destination) All(db *sqlx.DB) (*[]Destination, error) {
	x := new([]Destination)
	q := `
select *
from ab_atena a
where a.is_sender = false
and a.deleted = false
order by a.inserted_at desc`
	err := db.Select(x, q)
	return x, err
}
