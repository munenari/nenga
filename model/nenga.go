package model

import (
	"log"
	"strings"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	"golang.org/x/text/unicode/norm"
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
	for i, d := range *destinations {
		d.Address = convertAddress2EM(d.Address)
		(*destinations)[i] = d
	}
	x := &Nenga{
		Sender:       *sender,
		Destinations: *destinations,
	}
	return x, nil
}

type atenaAbstract struct {
	ID        uuid.UUID      `db:"id"`
	Names     pq.StringArray `db:"names"`
	Postcode  string         `db:"postcode"`
	Address   string         `db:"address"`
	Deleted   bool           `db:"deleted"`
	IsSender  bool           `db:"is_sender"`
	GroupName string         `db:"group_name"`
	Kind      int            `db:"kind"`
}

func (ab *atenaAbstract) ConvertEM() {
	ab.Address = convertAddress2EM(ab.Address)
}

func convertAddress2EM(str string) string {
	hyphen := "―"
	replacers := []*strings.Replacer{
		// charsReplacer('a', 'z', 'ａ'),
		charsReplacer('A', 'Z', 'Ａ'),
		strings.NewReplacer("ー", hyphen, "-", hyphen, "―", hyphen, "-", hyphen, "－", hyphen, "−", hyphen),
		strings.NewReplacer("0", "〇", "1", "一", "2", "二", "3", "三", "4", "四", "5", "五", "6", "六", "7", "七", "8", "八", "9", "九"),
		strings.NewReplacer("\n", "<br>"),
		strings.NewReplacer("<br>", "<br>　"),
	}
	str = norm.NFKC.String(str)
	for _, r := range replacers {
		str = r.Replace(str)
	}
	return str
}

func convertName2EM(name string) string {
	return strings.TrimSpace(strings.ReplaceAll(name, " ", "　"))
}

func charsReplacer(oldStart, oldEnd, newStart int) *strings.Replacer {
	chars := make([]string, 0)
	for c := oldStart; c <= oldEnd; c++ {
		chars = append(chars, string(c), string(newStart+(c-oldStart)))
	}
	return strings.NewReplacer(chars...)
}
