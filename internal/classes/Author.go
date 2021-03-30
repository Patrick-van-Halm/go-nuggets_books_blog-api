package classes

import (
	"database/sql"
	"errors"
	"github.com/matoous/go-nanoid/v2"
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

func (a *Author) New(db *sql.DB) error {
	if a.Name == "" {
		return errors.New("name cannot be empty")
	}

	id, err := gonanoid.New(21)
	if err != nil {
		return err
	}

	_, err = db.Exec("INSERT INTO authors (id, name) VALUES ($1, $2)", id, a.Name)
	if err == nil {
		a.Id = id
	}
	return err
}

func GetAllAuthors(db *sql.DB) ([]*Author, error) {
	authors := make([]*Author, 0)

	rows, err := db.Query(`SELECT id, name FROM authors`)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var a Author
		err := rows.Scan(&a.Id, &a.Name)
		if err != nil {
			return nil, err
		}

		authors = append(authors, &a)
	}

	return authors, nil
}