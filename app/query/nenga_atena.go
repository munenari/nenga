package query

import (
	"github.com/munenari/nenga/app/domain"

	"github.com/jmoiron/sqlx"
)

type Atena struct{}

func (Atena) List(tx sqlx.Ext, year int) ([]*domain.NengaAtena, error) {
	const q = `select * from atena where year = ? order by is_sender desc, group_name, id`
	x := make([]*domain.NengaAtena, 0)
	err := sqlx.Select(tx, &x, q, year)
	return x, err
}

func (Atena) Save(tx sqlx.Ext, x *domain.NengaAtena) error {
	const q = `
insert into atena (
	id, year, names, postcode, address, group_name, kind, is_sender
) values (
	:id, :year, :names, :postcode, :address, :group_name, :kind, :is_sender
)`
	_, err := sqlx.NamedExec(tx, q, x)
	return err
}

func (Atena) Delete(tx sqlx.Ext, id string) error {
	const q = `delete from atena where id = ?`
	_, err := tx.Exec(q, id)
	return err
}
