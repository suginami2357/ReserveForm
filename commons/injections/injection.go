package injections

import (
	"ReserveForm/models/places"
	"ReserveForm/models/reserves"
	"ReserveForm/models/users"
	sqlite_places "ReserveForm/repositories/sqlites/places_repository"
	sqlite_reserves "ReserveForm/repositories/sqlites/reserves_repository"
	sqlite_users "ReserveForm/repositories/sqlites/users_repository"
	test_places "ReserveForm/repositories/tests/places_repository"
	test_reserves "ReserveForm/repositories/tests/reserves_repository"
	test_users "ReserveForm/repositories/tests/users_repository"
)

type Type int

const (
	Sqlite Type = iota + 1
	Test
)

func User(t Type) users.Repository {
	switch t {
	case Sqlite:
		return new(sqlite_users.Repository)
	case Test:
		return new(test_users.Repository)
	default:
		panic("argument is undefined")
	}
}

func Reserve(t Type) reserves.Repository {
	switch t {
	case Sqlite:
		return new(sqlite_reserves.Repository)
	case Test:
		return new(test_reserves.Repository)
	default:
		panic("argument is undefined")
	}
}

func Place(t Type) places.Repository {
	switch t {
	case Sqlite:
		return new(sqlite_places.Repository)
	case Test:
		return new(test_places.Repository)
	default:
		panic("argument is undefined")
	}
}
