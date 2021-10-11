package reserves

import "ReserveForm/models/users"

type Repository interface {
	Index(users.User) []Reserve
	Create(Reserve)
	Delete(Reserve)
}
