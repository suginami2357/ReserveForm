package reserves_repository

import (
	"ReserveForm/models/reserves"
	"ReserveForm/models/users"
	"ReserveForm/repositories/postgreses"

	"github.com/jinzhu/gorm"
)

type Repository struct {
}

func (Repository) Index(user users.User) []reserves.Reserve {
	type field struct {
		ID   uint
		Name string
		Date string
	}
	db := postgreses.Open()

	var table = []field{}
	db.Table("reserves").
		Select("reserves.id, contents.name, reserves.date").
		Joins("left join contents on reserves.content_id = contents.id").
		Where("user_id = ?", user.ID).
		Order("date, name").
		Scan(&table)

	var results []reserves.Reserve
	for _, v := range table {
		reserve := reserves.Reserve{Model: gorm.Model{ID: v.ID}, ContentName: v.Name, Date: v.Date}
		results = append(results, reserve)
	}
	return results
}

func (Repository) Create(reserve reserves.Reserve) {
	db := postgreses.Open()
	db.Create(&reserve)
}

func (Repository) Delete(reserve reserves.Reserve) {
	db := postgreses.Open()
	db.Where("id = ?", reserve.ID).
		Where("user_id = ?", reserve.UserID).
		Unscoped(). //çİçċé¤
		Delete(&reserves.Reserve{})
}
