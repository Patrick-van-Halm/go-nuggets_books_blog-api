package book

import (
	"database/sql"
	"github.com/Patrick-van-Halm/nuggets_books_blog-api/internal/classes/Genre"
	"github.com/Patrick-van-Halm/nuggets_books_blog-api/internal/classes/Series"
	series_book "github.com/Patrick-van-Halm/nuggets_books_blog-api/internal/classes/SeriesBook"
	"github.com/Patrick-van-Halm/nuggets_books_blog-api/internal/classes/Author"
	"github.com/matoous/go-nanoid/v2"
)

type Book struct {
	Id			string			`json:"id"`
	Author		*author.Author	`json:"Author"`
	Title		string			`json:"title"`
	Genre 		*genre.Genre	`json:"genre"`
	Age			int				`json:"min-age"`
	CoverUrl	string			`json:"cover-url"`
	Series		*series.Series	`json:"series"`
}

func GetAll(db *sql.DB) ([]*Book, error) {
	books := make([]*Book, 0)

	rows, err := db.Query("SELECT id, author_id, genre_id, title, age, cover_url FROM books")
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var book Book
		var authorId sql.NullString
		var genreId sql.NullString

		// Get book data
		if err := rows.Scan(&book.Id, &authorId, &genreId, &book.Title, &book.Age, &book.CoverUrl); err != nil {
			return nil, err
		}

		// Get Author if set
		if authorId.Valid {
			author, err := author.Get(db, authorId.String)
			if err != nil {
				return nil, err
			}
			book.Author = author
		}

		// Get series if set
		seriesBook, err := series_book.GetWithBookId(db, book.Id)
		if err == nil {
			series, err := series.Get(db, seriesBook.SeriesId)
			if err != nil {
				return nil, err
			}
			book.Series = series
		}

		// Get genre if set
		if genreId.Valid {
			genre, err := genre.Get(db, genreId.String)
			if err != nil {
				return nil, err
			}
			book.Genre = genre
		}

		books = append(books, &book)
	}

	return books, nil
}

func Get(db *sql.DB, id string) (*Book, error) {
	row := db.QueryRow("SELECT id, author_id, genre_id, title, age, cover_url FROM books WHERE id = $1", id)

	var book Book
	var authorId sql.NullString
	var genreId sql.NullString

	// Get book data
	if err := row.Scan(&book.Id, &authorId, &genreId, &book.Title, &book.Age, &book.CoverUrl); err != nil {
		return nil, err
	}

	// Get Author if set
	if authorId.Valid {
		author, err := author.Get(db, authorId.String)
		if err != nil {
			return nil, err
		}
		book.Author = author
	}

	// Get series if set
	seriesBook, err := series_book.GetWithBookId(db, book.Id)
	if err == nil {
		series, err := series.Get(db, seriesBook.SeriesId)
		if err != nil {
			return nil, err
		}
		book.Series = series
	}

	// Get genre if set
	if genreId.Valid {
		genre, err := genre.Get(db, genreId.String)
		if err != nil {
			return nil, err
		}
		book.Genre = genre
	}

	return &book, nil
}

func New(db *sql.DB, book *Book) error {
	id, err := gonanoid.New(21)
	if err != nil {
		return err
	}

	_, err = db.Exec("INSERT INTO books (id, author_id, genre_id, title, age, cover_url) VALUES ($1, $2, $3, $4, $5, $6)", id, book.Author, book.Title, book.Genre, book.Age, book.CoverUrl)
	if err == nil {
		book.Id = id
	}
	return err
}