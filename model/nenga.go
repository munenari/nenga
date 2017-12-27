package model

import (
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

// AtenaTableName switch by year suffix or so
var AtenaTableName = "ab_atena"

// Nenga seems many nenga post cards
type Nenga struct {
	Destinations []Destination
	Sender       Sender
}

// MustGet all of nenga data (include sender and destinations)
func (Nenga) MustGet(db *sqlx.DB) *Nenga {
	x, err := Nenga{}.Get(db)
	if err != nil {
		log.Fatalln(err)
	}
	return x
}

// Get all of nenga data (include sender and destinations)
func (Nenga) Get(db *sqlx.DB) (*Nenga, error) {
	sender, err := Sender{}.Get(db)
	if err != nil {
		log.Println("failed to load sender information")
		return nil, err
	}
	destinations, err := Destination{}.All(db)
	if err != nil {
		log.Println("failed to load destinations information")
		return nil, err
	}
	x := &Nenga{
		Sender:       *sender,
		Destinations: *destinations,
	}
	return x, nil
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
