package classes

import (
	"database/sql"
	"errors"
)

type Author struct {
	Id		string `json:"id"`
	Name	string `json:"name"`
}

func (a *Author) Get(db *sql.DB) error {
	if a.Id == "" {
		return errors.New("id cannot be empty")
	}

	row := db.QueryRow(`SELECT name FROM authors WHERE id = $1`, a.Id)
	if err := row.Scan(&a.Name); err != nil {
		return err
	}
	return nil
}