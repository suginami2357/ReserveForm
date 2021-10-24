package reserves_repository

import (
	"ReserveForm/models/reserves"
	"ReserveForm/models/users"
)

type Repository struct {
}

var Data = []reserves.Reserve{}

func (Repository) Index(user users.User) []reserves.Reserve {
	return Data
}

func (Repository) Create(reserve reserves.Reserve) {
	Data = append(Data, reserve)
}

func (Repository) Delete(reserve reserves.Reserve) {
	result := []reserves.Reserve{}
	for _, v := range Data {
		if v.ID != reserve.ID {
			result = append(result, v)
		}
	}
	Data = result
}
