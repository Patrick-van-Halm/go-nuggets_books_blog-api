package classes

import (
	"database/sql"
	"errors"
	"github.com/matoous/go-nanoid/v2"
)

type Book struct {
	Id			string         	`json:"id"`
	Author		*Author        	`json:"Author"`
	Title		string          `json:"title"`
	Genre 		*Genre          `json:"genre"`
	Age			int           	`json:"min-age"`
	CoverUrl	string          `json:"cover-url"`
	Series		*Series 		`json:"series"`
}

func GetAllBooks(db *sql.DB) ([]*Book, error) {
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
			author := Author{
				Id: authorId.String,
			}
			err := author.Get(db)
			if err != nil {
				return nil, err
			}
			book.Author = &author
		}

		// Get series if set
		sb := SeriesBook{BookId: book.Id}
		err := sb.Get(db)
		if err == nil {
			series := Series{Id: sb.SeriesId}
			err := series.Get(db)
			if err != nil {
				return nil, err
			}
			book.Series = &series
		}

		// Get genre if set
		if genreId.Valid {
			genre := Genre{Id: genreId.String}
			err := genre.Get(db)
			if err != nil {
				return nil, err
			}
			book.Genre = &genre
		}

		books = append(books, &book)
	}

	return books, nil
}

func (b *Book) Get(db *sql.DB) error {
	if b.Id == "" {
		return errors.New("id cannot be empty")
	}

	row := db.QueryRow("SELECT author_id, genre_id, title, age, cover_url FROM books WHERE id = $1", b.Id)

	var authorId sql.NullString
	var genreId sql.NullString

	// Get book data
	if err := row.Scan(&authorId, &genreId, &b.Title, &b.Age, &b.CoverUrl); err != nil {
		return err
	}

	// Get Author if set
	if authorId.Valid {
		author := Author{
			Id: authorId.String,
		}
		err := author.Get(db)
		if err != nil {
			return err
		}
		b.Author = &author
	}

	// Get series if set
	sb := SeriesBook{BookId: b.Id}
	err := sb.Get(db)
	if err == nil {
		series := Series{Id: sb.SeriesId}
		err := series.Get(db)
		if err != nil {
			return err
		}
		b.Series = &series
	}

	// Get genre if set
	if genreId.Valid {
		genre := Genre{Id: genreId.String}
		err := genre.Get(db)
		if err != nil {
			return err
		}
		b.Genre = &genre
	}

	return nil
}

func (b *Book) New(db *sql.DB) error {
	if b.Author == nil || b.CoverUrl == "" || b.Age <= -1 || b.Genre == nil || b.Title == "" {
		return errors.New("certain values are not set")
	}

	id, err := gonanoid.New(21)
	if err != nil {
		return err
	}

	_, err = db.Exec("INSERT INTO books (id, author_id, genre_id, title, age, cover_url) VALUES ($1, $2, $3, $4, $5, $6)", id, b.Author.Id, b.Genre.Id, b.Title, b.Age, b.CoverUrl)
	if err == nil {
		b.Id = id
	}
	return err
}