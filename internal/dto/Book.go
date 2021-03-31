package dto

import (
	"database/sql"
	"github.com/Patrick-van-Halm/nuggets_books_blog-api/internal/models"
	"github.com/stroiman/go-automapper"
)

type BookDTO struct {
	Id			string         	`json:"id"`
	Author		*models.Author  `json:"author"`
	Title		string          `json:"title"`
	Genre 		*models.Genre   `json:"genre"`
	Age			int           	`json:"min-age"`
	CoverUrl	string          `json:"cover-url"`
	Series		*SeriesBook 	`json:"series"`
}

func (b *BookDTO) Get(book *models.Book, db *sql.DB) error {
	automapper.MapLoose(book, b)

	if book.GenreId != "" {
		genre := models.Genre{Id: book.GenreId}
		err := genre.Get(db)
		if err != nil {
			return err
		}
		b.Genre = &genre
	}

	if book.AuthorId != "" {
		author := models.Author{Id: book.AuthorId}
		err := author.Get(db)
		if err != nil {
			return err
		}
		b.Author = &author
	}

	if book.SeriesId != "" {
		series := models.Series{Id: book.SeriesId}
		if err := series.Get(db); err != nil {
			return err
		}

		sb := SeriesBook{}
		if err := sb.Get(&series, book.Id, db); err != nil {
			return err
		}

		b.Series = &sb
	}

	return nil
}