package contents_repository

import (
	"ReserveForm/models/contents"
	"ReserveForm/repositories/postgreses"
)

type Repository struct {
}

func (Repository) Index() []contents.Content {
	var contents []contents.Content
	db := postgreses.Open()
	db.Find(&contents)
	return contents
}
