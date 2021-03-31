package models

import (
	"database/sql"
	"errors"
	"github.com/matoous/go-nanoid/v2"
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

func GetAllGenres(db *sql.DB) ([]*Genre, error) {
	genres := make([]*Genre, 0)

	rows, err := db.Query(`SELECT id, name FROM genres`)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var g Genre
		err := rows.Scan(&g.Id, &g.Name)
		if err != nil {
			return nil, err
		}

		genres = append(genres, &g)
	}

	return genres, nil
}

func (g *Genre) New(db *sql.DB) error {
	if g.Name == "" {
		return errors.New("name cannot be empty")
	}

	id, err := gonanoid.New(21)

	_, err = db.Exec(`INSERT INTO genres (id, name) VALUES ($1, $2)`, id, g.Name)
	if err != nil {
		return err
	}
	g.Id = id
	return nil
}

func (g *Genre) TypeName() string {
	return "genre"
}