package series

import (
	"database/sql"
	gonanoid "github.com/matoous/go-nanoid/v2"
)

type Series struct {
	Id			string	`json:"id"`
	Name		string	`json:"name"`
}

func Get(db *sql.DB, id string) (*Series, error) {
	row := db.QueryRow(`SELECT id, name FROM series WHERE id = $1`, id)
	var series Series
	if err := row.Scan(&series.Id, &series.Name); err != nil {
		return nil, err
	}

	return &series, nil
}

func New(db *sql.DB, series *Series) error {
	id, err := gonanoid.New(21)
	if err != nil {
		return err
	}

	if _, err := db.Exec(`INSERT INTO series(id, name) VALUES ($1, $2)`, id, series.Name); err != nil {
		return err
	}

	series.Id = id
	return nil
}