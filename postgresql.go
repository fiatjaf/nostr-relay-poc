// +build !lite

package main

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func initDB() (*sqlx.DB, error) {
	db, err := sqlx.Connect("postgres", s.PostgresDatabase)
	if err != nil {
		return nil, err
	}

	_, err = db.Exec(`
CREATE TABLE IF NOT EXISTS event (
  id text NOT NULL,
  pubkey text NOT NULL,
  created_at integer NOT NULL,
  kind integer NOT NULL,
  tags jsonb NOT NULL,
  content text NOT NULL,
  sig text NOT NULL
);

CREATE UNIQUE INDEX IF NOT EXISTS ididx ON event (id);
CREATE UNIQUE INDEX IF NOT EXISTS pubkeytimeidx ON event (pubkey, created_at);
    `)
	log.Print(err)
	return db, nil
}

const relatedEventsCondition = `tags @@ '$[*][1] == "' || ? || '"'`
