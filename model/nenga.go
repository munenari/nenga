package model

import (
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
)

// Nenga seems many nenga post cards
type Nenga struct {
	Destinations []Destination
	Sender       Sender
}

type atenaAbstract struct {
	ID         uuid.UUID      `db:"id"`
	Names      pq.StringArray `db:"names"`
	Postcode   string         `db:"postcode"`
	Address    string         `db:"address"`
	Deleted    bool           `db:"deleted"`
	InsertedAt time.Time      `db:"inserted_at"`
	IsSender   bool           `db:"is_sender"`
}
