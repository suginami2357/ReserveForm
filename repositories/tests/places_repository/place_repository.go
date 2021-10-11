package places_repository

import (
	"ReserveForm/models/places"
)

type Repository struct {
}

func (Repository) Index() []places.Place {
	var places []places.Place
	// db, _ := gorm.Open("sqlite3", "data.sqlite3")
	// db.Find(&places)
	// defer db.Close()
	return places
}
