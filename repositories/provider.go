package repositories

import (
	"ReserveForm/commons/injections"
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

func User(t injections.Type) users.Repository {
	switch t {
	case injections.Sqlite:
		return new(sqlite_users.Repository)
	case injections.Test:
		return new(test_users.Repository)
	default:
		panic("argument is undefined")
	}
}

func Reserve(t injections.Type) reserves.Repository {
	switch t {
	case injections.Sqlite:
		return new(sqlite_reserves.Repository)
	case injections.Test:
		return new(test_reserves.Repository)
	default:
		panic("argument is undefined")
	}
}

func Place(t injections.Type) places.Repository {
	switch t {
	case injections.Sqlite:
		return new(sqlite_places.Repository)
	case injections.Test:
		return new(test_places.Repository)
	default:
		panic("argument is undefined")
	}
}
