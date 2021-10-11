package places_repository

import (
	"ReserveForm/models/places"

	"github.com/jinzhu/gorm"
)

type Repository struct {
}

func (Repository) Index() []places.Place {
	var places []places.Place
	db, _ := gorm.Open("sqlite3", "data.sqlite3")
	db.Find(&places)
	defer db.Close()
	return places
}
