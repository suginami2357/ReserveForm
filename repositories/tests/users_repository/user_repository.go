package users_repository

import (
	"ReserveForm/models/users"
	"errors"
)

type Repository struct {
}

var Data = []users.User{}

func (Repository) Create(user users.User) (*users.User, error) {
	Data = append(Data, user)
	return &user, nil
}

func (Repository) Take(src users.User) (*users.User, error) {
	if src.ID > 0 {
		for _, v := range Data {
			if src.ID == v.ID {
				return &v, nil
			}
		}
	} else if src.Email != "" {
		for _, v := range Data {
			if src.Email == v.Email {
				return &v, nil
			}
		}
	} else if src.Token != nil {
		for _, user := range Data {
			if string(src.Token) == string(user.Token) {
				return &user, nil
			}
		}
	}
	return nil, errors.New("search parameters is not applicable")
}

func (Repository) Update(user users.User) {
	for _, v := range Data {
		if user.ID == v.ID {
			v = user
		}
	}
}
