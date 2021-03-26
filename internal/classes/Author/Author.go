package author

import (
	"database/sql"
)

type Author struct {
	Id		string `json:"id"`
	Name	string `json:"name"`
}

func Get(db *sql.DB, id string) (*Author, error) {
	row := db.QueryRow(`SELECT id, name FROM authors WHERE id = $1`, id)
	var author Author
	if err := row.Scan(&author.Id, &author.Name); err != nil {
		return nil, err
	}

	return &author, nil
}