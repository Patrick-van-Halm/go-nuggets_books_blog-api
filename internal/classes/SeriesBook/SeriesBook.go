package series_book

import (
	"database/sql"
)

type SeriesBook struct {
	SeriesId	string	`json:"series_id"`
	BookId		string	`json:"book_id"`
	Number		int	`json:"number"`
}

func GetWithBookId(db *sql.DB, bookId string) (*SeriesBook, error) {
	row := db.QueryRow(`SELECT series_id, book_id, number FROM series_books WHERE book_id = $1`, bookId)
	var seriesBook SeriesBook
	if err := row.Scan(&seriesBook.SeriesId, &seriesBook.BookId, &seriesBook.Number); err != nil {
		return nil, err
	}

	return &seriesBook, nil
}

func New(db *sql.DB, seriesBook *SeriesBook) error {
	if _, err := db.Exec(`INSERT INTO series_books(series_id, book_id, number) VALUES ($1, $2, $3)`, seriesBook.SeriesId, seriesBook.BookId, seriesBook.Number); err != nil {
		return err
	}

	return nil
}