package dto

import (
	"database/sql"
	"github.com/Patrick-van-Halm/nuggets_books_blog-api/internal/models"
)

type CreateBookDTO struct {
	AuthorId	string  		`json:"author_id"`
	Title		string          `json:"title"`
	GenreId		string 			`json:"genre_id"`
	Age			int           	`json:"min-age"`
	CoverUrl	string          `json:"cover-url"`
	SeriesId	string 			`json:"series_id"`
	BookNum		uint8			`json:"book_num"`
}

func (b *CreateBookDTO) Create(db *sql.DB) error {
	book := models.Book{
		AuthorId: b.AuthorId,
		Title:    b.Title,
		GenreId:  b.GenreId,
		Age:      b.Age,
		CoverUrl: b.CoverUrl,
	}

	err := book.New(db)
	if err != nil{
		return err
	}

	if b.SeriesId != "" {
		sb := models.SeriesBook{
			SeriesId: b.SeriesId,
			BookId:   book.Id,
			Number:   b.BookNum,
		}

		err := sb.New(db)
		if err != nil{
			return err
		}
	}

	return nil
}