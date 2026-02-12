package query

import (
	"cmp"
	"log/slog"
	"slices"
	"time"

	"github.com/jmoiron/sqlx"
)

type (
	SchemaMigration struct {
		Version   int       `db:"version"`
		AppliedAt time.Time `db:"applied_at"`
	}
	Migration struct {
		Version int
		Fn      func(tx *sqlx.Tx) error
	}
)

var migrations = []Migration{
	{
		Version: 1,
		Fn: func(tx *sqlx.Tx) error {
			const q = `
				create table atena (
					id         UUID    NOT NULL PRIMARY KEY,
					year       INTEGER NOT NULL DEFAULT '2026',
					names      TEXT[]  NOT NULL,
					postcode   TEXT    NOT NULL DEFAULT '',
					address    TEXT    NOT NULL DEFAULT '',
					group_name TEXT    NOT NULL DEFAULT '',
					kind       INTEGER NOT NULL DEFAULT 0,
					is_sender  BOOLEAN NOT NULL DEFAULT false
				);
				create index atena_year_idx on atena (year);
			`
			_, err := tx.Exec(q)
			return err
		},
	},
}

func runMigrate(d *sqlx.DB) error {
	return Transaction(d, func(tx *sqlx.Tx) error {
		if err := createMigrationTable(tx); err != nil {
			return err
		}
		applied, err := getMigration(tx)
		if err != nil {
			return err
		}
		latestVersion := 0
		if len(applied) > 0 {
			latestApplied := slices.MaxFunc(applied, func(a, b *SchemaMigration) int {
				return cmp.Compare(a.Version, b.Version)
			})
			latestVersion = latestApplied.Version
		}
		slog.Info("current migrate version", slog.Int("version", latestVersion))
		for _, m := range migrations {
			if m.Version <= latestVersion {
				continue
			}
			slog.Info("migrating", slog.Int("version", m.Version))
			if err := m.Fn(tx); err != nil {
				return err
			}
			if err := saveMigration(tx, m.Version); err != nil {
				return err
			}
		}
		return nil
	})
}

func createMigrationTable(tx *sqlx.Tx) error {
	const q = `
CREATE TABLE IF NOT EXISTS schema_migrations (
	version INTEGER PRIMARY KEY,
	applied_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now()
)`
	_, err := tx.Exec(q)
	return err
}
func getMigration(tx *sqlx.Tx) ([]*SchemaMigration, error) {
	x := make([]*SchemaMigration, 0)
	const q = `select version, applied_at from schema_migrations order by version`
	err := tx.Select(&x, q)
	return x, err
}
func saveMigration(tx *sqlx.Tx, version int) error {
	const q = `insert into schema_migrations (version) values (:version)`
	_, err := tx.NamedExec(q, map[string]any{"version": version})
	return err
}
