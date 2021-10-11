package users_repository

import (
	"ReserveForm/models/users"
	"ReserveForm/repositories/sqlites"
	"errors"
)

type Repository struct{}

func (Repository) Create(user users.User) (*users.User, error) {
	db := sqlites.Open()
	defer db.Close()
	err := db.Create(&user).Error
	return &user, err
}

func (Repository) Take(src users.User) (*users.User, error) {
	db := sqlites.Open()
	defer db.Close()

	if src.ID > 0 {
		db = db.Where("id = ?", src.ID)
	} else if src.Email != "" {
		db = db.Where("email = ?", src.Email)
	} else if src.Token != nil {
		db = db.Where("token = ?", src.Token)
	} else {
		return nil, errors.New("search parameters is not applicable")
	}

	var user users.User
	err := db.Take(&user).Error
	return &user, err
}

func (Repository) Update(user users.User) {
	db := sqlites.Open()
	defer db.Close()
	db.Where("id = ?", user.ID).Save(&user)
}
