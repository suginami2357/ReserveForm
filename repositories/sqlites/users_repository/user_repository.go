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

	var user users.User
	var err error
	if src.ID > 0 {
		err = db.Where("id = ?", src.ID).First(&user).Error
	} else if src.Email != "" {
		err = db.Where("email = ?", src.Email).First(&user).Error
	} else if src.Token != nil {
		err = db.Where("token = ?", src.Token).First(&user).Error
	} else {
		err = errors.New("search parameters is not applicable")
	}
	return &user, err
}

func (Repository) Update(user users.User) {
	db := sqlites.Open()
	defer db.Close()
	db.Where("id = ?", user.ID).Save(&user)
}
