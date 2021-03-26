package review

import (
	"database/sql"
	"github.com/Patrick-van-Halm/nuggets_books_blog-api/internal/classes/Book"
)

type Review struct {
	Id		string		`json:"id"`
	Rating	uint8		`json:"rating"`
	Book	*book.Book	`json:"book"`
	Text	string		`json:"text"`
}

func GetAll(db *sql.DB) ([]*Review, error) {
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
		review.Book, err = book.Get(db, bookId)
		if err != nil {
			return nil, err
		}

		reviews = append(reviews, &review)
	}

	return reviews, nil
}

func Get(db *sql.DB, id string) (*Review, error) {
	row := db.QueryRow("SELECT id, book_id, rating, review FROM reviews WHERE id = $1", id)
	var review Review
	var bookId string
	if err := row.Scan(&review.Id, &bookId, &review.Rating, &review.Text); err != nil {
		return nil, err
	}

	book, err := book.Get(db, bookId)
	if err != nil {
		return nil, err
	}

	review.Book = book
	return &review, nil
}