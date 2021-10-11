package users

type Repository interface {
	Create(User) (*User, error)
	Take(User) (*User, error)
	Update(User)
}
