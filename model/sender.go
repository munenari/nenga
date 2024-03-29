package model

import "github.com/jmoiron/sqlx"

// Sender seems a NENGA senders
type Sender atenaAbstract

// Get a Sender data
func (Sender) Get(db *sqlx.DB) (*Sender, error) {
	x := new(Sender)
	q := `
select *
from ` + AtenaTableName + ` a
where a.is_sender = true
and a.deleted = false
limit 1`
	err := db.Get(x, q)
	x.Address = convertAddress2EM(x.Address)
	return x, err
}
