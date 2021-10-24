package places_repository

import (
	"ReserveForm/models/places"

	"gorm.io/gorm"
)

type Repository struct {
}

func (Repository) Index() []places.Place {
	return []places.Place{
		places.Place{Model: gorm.Model{ID: 1}, Name: "テスト1"},
		places.Place{Model: gorm.Model{ID: 2}, Name: "テスト2"},
		places.Place{Model: gorm.Model{ID: 3}, Name: "テスト3"},
	}
}
