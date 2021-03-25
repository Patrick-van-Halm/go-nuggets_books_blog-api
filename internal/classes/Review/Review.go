package review

import (
	"database/sql"
	"github.com/Patrick-van-Halm/nuggets_books_blog-api/internal/hasher"
	"github.com/Patrick-van-Halm/nuggets_books_blog-api/internal/classes/Book"
	"log"
	"strconv"
)

type Review struct {
	Id		string		`json:"id"`
	Rating	uint8		`json:"rating"`
	Book	*book.Book	`json:"book"`
	Text	string		`json:"text"`
}

func GetAll(db *sql.DB) []*Review {
	reviews := make([]*Review, 0)

	rows, err := db.Query("SELECT id, book_id, rating, review FROM reviews")
	if err != nil {
		log.Println(err)
		return reviews
	}

	for rows.Next() {
		var review Review
		var bookId uint
		if err := rows.Scan(&review.Id, &bookId, &review.Rating, &review.Text); err != nil {
			log.Println(err)
			return reviews
		}

		id, _ := strconv.Atoi(review.Id)
		review.Id = hasher.HashID(id)
		review.Book = book.Get(db, bookId)

		reviews = append(reviews, &review)
	}

	return reviews
}

func GetWithHash(db *sql.DB, hashedId string) *Review {
	id := hasher.GetFromHashID(hashedId)
	row := db.QueryRow("SELECT id, book_id, rating, review FROM reviews WHERE id = $1", id)
	var review Review
	var bookId uint
	if err := row.Scan(&review.Id, &bookId, &review.Rating, &review.Text); err != nil {
		log.Println(err)
		return nil
	}

	review.Id = hashedId
	review.Book = book.Get(db, bookId)

	return &review
}