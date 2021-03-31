package models

import (
	"database/sql"
	"errors"
)

type SeriesBook struct {
	SeriesId	string	`json:"series_id"`
	BookId		string	`json:"book_id"`
	Number		uint8	`json:"number"`
}

func (sb *SeriesBook) Get(db *sql.DB) error {
	if (sb.SeriesId == "" || sb.Number == 0) && sb.BookId == "" {
		return errors.New("you need at least a book Id or a series Id and the book number")
	}

	if sb.BookId != "" {
		return sb.getWithBookId(db)
	}

	if sb.SeriesId != "" && sb.Number != 0 {
		return sb.getWithSeriesId(db)
	}

	return nil
}

func (sb *SeriesBook) getWithBookId(db *sql.DB) error {
	row := db.QueryRow(`SELECT series_id, number FROM series_books WHERE book_id = $1`, sb.BookId)
	if err := row.Scan(&sb.SeriesId, &sb.Number); err != nil {
		return err
	}

	return nil
}

func (sb *SeriesBook) getWithSeriesId(db *sql.DB) error {
	row := db.QueryRow(`SELECT book_id FROM series_books WHERE series_id = $1 AND book_id = $2`, sb.SeriesId, sb.Number)
	if err := row.Scan(&sb.BookId); err != nil {
		return err
	}

	return nil
}

func (sb *SeriesBook) New(db *sql.DB) error {
	if _, err := db.Exec(`INSERT INTO series_books(series_id, book_id, number) VALUES ($1, $2, $3)`, sb.SeriesId, sb.BookId, sb.Number); err != nil {
		return err
	}

	return nil
}

func (sb *SeriesBook) TypeName() string {
	return "series_book"
}