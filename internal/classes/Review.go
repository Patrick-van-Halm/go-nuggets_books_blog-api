package classes

import (
	"database/sql"
)

type Review struct {
	Id		string `json:"id"`
	Rating	uint8  `json:"rating"`
	Book	*Book    `json:"book"`
	Text	string   `json:"text"`
}

func GetAllReviews(db *sql.DB) ([]*Review, error) {
	reviews := make([]*Review, 0)

	rows, err := db.Query("SELECT id, book_id, rating, review FROM reviews")
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var review Review
		var bookId string
		if err := rows.Scan(&review.Id, &bookId, &review.Rating, &review.Text); err != nil {
			return nil, err
		}
		book := Book{Id: bookId}
		err := book.Get(db)
		if err != nil {
			return nil, err
		}

		reviews = append(reviews, &review)
	}

	return reviews, nil
}

func (r *Review) Get(db *sql.DB) error {
	row := db.QueryRow("SELECT book_id, rating, review FROM reviews WHERE id = $1", r.Id)
	var review Review
	var bookId string
	if err := row.Scan(&bookId, &review.Rating, &review.Text); err != nil {
		return err
	}

	book := Book {Id: bookId}
	err := book.Get(db)
	if err != nil {
		return err
	}

	review.Book = &book
	return nil
}