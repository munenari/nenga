package domain

import "github.com/google/uuid"

type NengaAtena struct {
	ID        uuid.UUID `db:"id"         json:"id"`
	Year      int       `db:"year"       json:"year"`
	Names     []string  `db:"names"      json:"names"`
	Postcode  string    `db:"postcode"   json:"postcode"`
	Address   string    `db:"address"    json:"address"`
	GroupName string    `db:"group_name" json:"group_name"`
	Kind      int       `db:"kind"       json:"kind"`
	IsSender  bool      `db:"is_sender"  json:"is_sender"`
}
