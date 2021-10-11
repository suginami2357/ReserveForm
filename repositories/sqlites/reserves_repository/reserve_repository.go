package reserves_repository

import (
	"ReserveForm/models/reserves"
	"ReserveForm/models/users"
	"ReserveForm/repositories/sqlites"

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
	db := sqlites.Open()
	defer db.Close()

	var table = []field{}
	db.Table("reserves").
		Select("reserves.id, places.name, reserves.date").
		Joins("left join places on reserves.place_id = places.id").
		Where("user_id = ?", user.ID).
		Order("date, name").
		Scan(&table)

	var results []reserves.Reserve
	for _, v := range table {
		reserve := reserves.Reserve{Model: gorm.Model{ID: v.ID}, PlaceName: v.Name, Date: v.Date}
		results = append(results, reserve)
	}
	return results
}

func (Repository) Create(reserve reserves.Reserve) {
	db := sqlites.Open()
	defer db.Close()

	db.Create(&reserve)
}

func (Repository) Delete(reserve reserves.Reserve) {
	db := sqlites.Open()
	defer db.Close()

	db.Where("id = ?", reserve.ID).
		Where("user_id = ?", reserve.UserID).
		Unscoped(). //物理削除
		Delete(&reserves.Reserve{})
}
