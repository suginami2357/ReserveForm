package repositories

import (
	"ReserveForm/commons/injections"
	"ReserveForm/models/contents"
	"ReserveForm/models/reserves"
	"ReserveForm/models/users"
	sqlite_contents "ReserveForm/repositories/postgreses/contents_repository"
	sqlite_reserves "ReserveForm/repositories/postgreses/reserves_repository"
	sqlite_users "ReserveForm/repositories/postgreses/users_repository"
	test_contents "ReserveForm/repositories/tests/contents_repository"
	test_reserves "ReserveForm/repositories/tests/reserves_repository"
	test_users "ReserveForm/repositories/tests/users_repository"
)

func User(t injections.Type) users.Repository {
	switch t {
	case injections.Postgres:
		return new(sqlite_users.Repository)
	case injections.Test:
		return new(test_users.Repository)
	default:
		panic("argument is undefined")
	}
}

func Reserve(t injections.Type) reserves.Repository {
	switch t {
	case injections.Postgres:
		return new(sqlite_reserves.Repository)
	case injections.Test:
		return new(test_reserves.Repository)
	default:
		panic("argument is undefined")
	}
}

func Content(t injections.Type) contents.Repository {
	switch t {
	case injections.Postgres:
		return new(sqlite_contents.Repository)
	case injections.Test:
		return new(test_contents.Repository)
	default:
		panic("argument is undefined")
	}
}
