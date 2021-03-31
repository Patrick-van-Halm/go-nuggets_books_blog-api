package models

import (
	"database/sql"
	"errors"
	"github.com/matoous/go-nanoid/v2"
)

type Book struct {
	Id			string         	`json:"id"`
	AuthorId	string		    `json:"author_id"`
	SeriesId	string	 		`json:"series_id"`
	Title		string          `json:"title"`
	GenreId 	string          `json:"genre_id"`
	Age			int           	`json:"min-age"`
	CoverUrl	string          `json:"cover-url"`
}

func GetAllBooks(db *sql.DB) ([]*Book, error) {
	books := make([]*Book, 0)

	rows, err := db.Query("SELECT id, author_id, genre_id, title, age, cover_url, series_books.series_id FROM books LEFT JOIN series_books ON books.id = series_books.book_id")
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var book Book
		var authorId sql.NullString
		var genreId sql.NullString
		var seriesId sql.NullString

		// Get book data
		if err := rows.Scan(&book.Id, &authorId, &genreId, &book.Title, &book.Age, &book.CoverUrl, &seriesId); err != nil {
			return nil, err
		}
		if authorId.Valid {
			book.AuthorId = authorId.String
		}

		if genreId.Valid {
			book.GenreId = genreId.String
		}

		if seriesId.Valid {
			book.SeriesId = seriesId.String
		}
		books = append(books, &book)
	}

	return books, nil
}

func (b *Book) Get(db *sql.DB) error {
	if b.Id == "" {
		return errors.New("id cannot be empty")
	}

	row := db.QueryRow("SELECT author_id, genre_id, title, age, cover_url, series_books.series_id FROM books LEFT JOIN series_books ON books.id = series_books.book_id WHERE id = $1", b.Id)

	var authorId sql.NullString
	var genreId sql.NullString
	var seriesId sql.NullString

	// Get book data
	if err := row.Scan(&authorId, &genreId, &b.Title, &b.Age, &b.CoverUrl, &seriesId); err != nil {
		return err
	}

	// Get Author if set
	if authorId.Valid {
		b.AuthorId = authorId.String
	}

	// Get series if set
	if seriesId.Valid {
		b.SeriesId = seriesId.String
	}

	// Get genre if set
	if genreId.Valid {
		b.GenreId = genreId.String
	}

	return nil
}

func (b *Book) New(db *sql.DB) error {
	if b.AuthorId == "" || b.CoverUrl == "" || b.Age <= -1 || b.GenreId == "" || b.Title == "" {
		return errors.New("certain values are not set")
	}

	id, err := gonanoid.New(21)
	if err != nil {
		return err
	}

	_, err = db.Exec("INSERT INTO books (id, author_id, genre_id, title, age, cover_url) VALUES ($1, $2, $3, $4, $5, $6)", id, b.AuthorId, b.GenreId, b.Title, b.Age, b.CoverUrl)
	if err == nil {
		b.Id = id
	}
	return err
}

func (b *Book) TypeName() string {
	return "book"
}