package classes

import (
	"database/sql"
	"errors"
)

type Genre struct {
	Id 		string `json:"id"`
	Name 	string `json:"name"`
}

func (g *Genre) Get(db *sql.DB) error {
	if g.Id == "" {
		return errors.New("id cannot be empty")
	}

	row := db.QueryRow(`SELECT name FROM genres WHERE id = $1`, g.Id)
	if err := row.Scan(&g.Name); err != nil {
		return err
	}

	return nil
}