package injections

import (
	"ReserveForm/models/contents"
	"ReserveForm/models/reserves"
	"ReserveForm/models/users"

	postgres_contents "ReserveForm/repositories/postgreses/contents_repository"
	postgres_reserves "ReserveForm/repositories/postgreses/reserves_repository"
	postgres_users "ReserveForm/repositories/postgreses/users_repository"

	test_contents "ReserveForm/repositories/tests/contents_repository"
	test_reserves "ReserveForm/repositories/tests/reserves_repository"
	test_users "ReserveForm/repositories/tests/users_repository"
)

type Type int

const (
	Postgres Type = iota + 1
	Test
)

func User(t Type) users.Repository {
	switch t {
	case Postgres:
		return new(postgres_users.Repository)
	case Test:
		return new(test_users.Repository)
	default:
		panic("argument is undefined")
	}
}

func Reserve(t Type) reserves.Repository {
	switch t {
	case Postgres:
		return new(postgres_reserves.Repository)
	case Test:
		return new(test_reserves.Repository)
	default:
		panic("argument is undefined")
	}
}

func Place(t Type) contents.Repository {
	switch t {
	case Postgres:
		return new(postgres_contents.Repository)
	case Test:
		return new(test_contents.Repository)
	default:
		panic("argument is undefined")
	}
}
