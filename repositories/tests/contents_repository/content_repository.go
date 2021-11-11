package contents_repository

import (
	"ReserveForm/models/contents"
)

type Repository struct {
}

var Data = []contents.Content{}

func (Repository) Index() []contents.Content {
	return Data
}
