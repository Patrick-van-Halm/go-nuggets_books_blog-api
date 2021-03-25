package book

import (
	"database/sql"
	"github.com/Patrick-van-Halm/nuggets_books_blog-api/internal/hasher"
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

func GetAll(db *sql.DB) []*Book {
	books := make([]*Book, 0)

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

func GetWithHash(db *sql.DB, hashedId string) *Book {
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

func Get(db *sql.DB, id uint) *Book {
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

func New(db *sql.DB,book Book) error {
	if book.Series != nil {
		_, err := db.Exec("INSERT INTO books (author, title, genre, age, cover_url, series, series_book_num) VALUES ($1, $2, $3, $4, $5, $6, $7)", book.Author, book.Title, book.Genre, book.Age, book.CoverUrl, book.Series.Name, book.Series.Number)
		return err
	}

	_, err := db.Exec("INSERT INTO books (author, title, genre, age, cover_url) VALUES ($1, $2, $3, $4, $5)", book.Author, book.Title, book.Genre, book.Age, book.CoverUrl)
	return err
}