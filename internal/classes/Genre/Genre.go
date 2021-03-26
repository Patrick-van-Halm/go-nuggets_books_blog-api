package genre

import (
	"database/sql"
)

type Genre struct {
	Id 		string `json:"id"`
	Name 	string `json:"name"`
}

func Get(db *sql.DB, id string) (*Genre, error) {
	row := db.QueryRow(`SELECT id, name FROM genres WHERE id = $1`, id)
	var genre Genre
	if err := row.Scan(&genre.Id, &genre.Name); err != nil {
		return nil, err
	}

	return &genre, nil
}