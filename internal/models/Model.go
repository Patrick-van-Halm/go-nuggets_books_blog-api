package models

import "database/sql"

type Model interface {
	Get(db *sql.DB) error
	New(db *sql.DB) error
	TypeName() string
}
