package dto

import (
	"database/sql"
	"github.com/Patrick-van-Halm/nuggets_books_blog-api/internal/models"
	"github.com/stroiman/go-automapper"
)

type SeriesBook struct {
	Id		string	`json:"id"`
	Name	string	`json:"name"`
	BookNum	uint8	`json:"book_num"`
}

func (sb *SeriesBook) Get(model *models.Series, bookId string, db *sql.DB) error {
	automapper.MapLoose(model, sb)

	m := models.SeriesBook{BookId: bookId}
	if err := m.Get(db); err != nil {
		return err
	}

	sb.BookNum = m.Number
	return nil
}