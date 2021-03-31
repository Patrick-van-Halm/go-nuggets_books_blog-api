package models

import (
	"database/sql"
	"github.com/matoous/go-nanoid/v2"
)

type Series struct {
	Id			string	`json:"id"`
	Name		string	`json:"name"`
}

func (s *Series) Get(db *sql.DB) error {
	row := db.QueryRow(`SELECT name FROM series WHERE id = $1`, s.Id)
	if err := row.Scan(&s.Name); err != nil {
		return err
	}

	return nil
}

func (s *Series) New(db *sql.DB) error {
	id, err := gonanoid.New(21)
	if err != nil {
		return err
	}

	if _, err := db.Exec(`INSERT INTO series(id, name) VALUES ($1, $2)`, id, s.Name); err != nil {
		return err
	}

	s.Id = id
	return nil
}

func (s *Series) TypeName() string {
	return "series"
}

func GetAllSeries(db *sql.DB) ([]*Series, error) {
	series := make([]*Series, 0)

	rows, err := db.Query(`SELECT id, name FROM series`)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var s Series
		err := rows.Scan(&s.Id, &s.Name)
		if err != nil {
			return nil, err
		}

		series = append(series, &s)
	}

	return series, nil
}