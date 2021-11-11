package contents_repository

import (
	"ReserveForm/models/contents"

	"github.com/jinzhu/gorm"
)

type Repository struct {
}

func (Repository) Index() []contents.Content {
	var places []contents.Content
	db, _ := gorm.Open("sqlite3", "data.sqlite3")
	db.Find(&places)
	defer db.Close()
	return places
}
