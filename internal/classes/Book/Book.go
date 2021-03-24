package book

import (
	"database/sql"
	"github.com/Patrick-van-Halm/nuggets_books_blog-api/internal/env"
	"github.com/Patrick-van-Halm/nuggets_books_blog-api/internal/hasher"
	_ "github.com/lib/pq"
	"log"
	"strconv"
)

type Book struct {
	Id			string	`json:"id"`
	Author		string	`json:"author"`
	Title		string	`json:"title"`
	Genre 		string	`json:"genre"`
	Age			int		`json:"min-age"`
	CoverUrl	string	`json:"cover-url"`
	Series		*Series	`json:"series"`
}

type Series struct {
	Name	string	`json:"name"`
	Number	int32	`json:"number"`
}

func GetAll() []*Book {
	books := make([]*Book, 0)
	db, err := sql.Open("postgres", env.GetEnvVar("DB_CONNECTION_STRING"))
	if err != nil {
		log.Println(err)
		return books
	}

	rows, err := db.Query("SELECT id, author, title, genre, age, cover_url, series, series_book_num FROM books")
	if err != nil {
		log.Println(err)
		return books
	}

	for rows.Next() {
		var book Book
		var series sql.NullString
		var seriesNum sql.NullInt32
		if err := rows.Scan(&book.Id, &book.Author, &book.Title, &book.Genre, &book.Age, &book.CoverUrl, &series, &seriesNum); err != nil {
			log.Println(err)
			return books
		}
		id, _ := strconv.Atoi(book.Id)
		book.Id = hasher.HashID(id)

		if series.Valid {
			book.Series = &Series{
				Name:   series.String,
				Number: seriesNum.Int32,
			}
		}

		books = append(books, &book)
	}

	return books
}

func GetWithHash(hashedId string) *Book {
	db, err := sql.Open("postgres", env.GetEnvVar("DB_CONNECTION_STRING"))
	if err != nil {
		log.Println(err)
		return nil
	}

	id := hasher.GetFromHashID(hashedId)
	row := db.QueryRow("SELECT id, author, title, genre, age, cover_url, series, series_book_num FROM books WHERE id = $1", id)
	var book Book
	var series sql.NullString
	var seriesNum sql.NullInt32
	if err := row.Scan(&book.Id, &book.Author, &book.Title, &book.Genre, &book.Age, &book.CoverUrl, &series, &seriesNum); err != nil {
		log.Println(err)
		return nil
	}

	if series.Valid {
		book.Series = &Series{
			Name:   series.String,
			Number: seriesNum.Int32,
		}
	}

	book.Id = hashedId

	return &book
}

func Get(id uint) *Book {
	db, err := sql.Open("postgres", env.GetEnvVar("DB_CONNECTION_STRING"))
	if err != nil {
		log.Println(err)
		return nil
	}

	row := db.QueryRow("SELECT id, author, title, genre, age, cover_url, series, series_book_num FROM books WHERE id = $1", id)
	var book Book
	var series sql.NullString
	var seriesNum sql.NullInt32
	if err := row.Scan(&book.Id, &book.Author, &book.Title, &book.Genre, &book.Age, &book.CoverUrl, &series, &seriesNum); err != nil {
		log.Println(err)
		return nil
	}

	hashedId, _ := strconv.Atoi(book.Id)
	book.Id = hasher.HashID(hashedId)

	if series.Valid {
		book.Series = &Series{
			Name:   series.String,
			Number: seriesNum.Int32,
		}
	}

	return &book
}