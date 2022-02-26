package main

import (
	"ReserveForm/commons/injections"
	"ReserveForm/commons/servers"
	"ReserveForm/models/contents"
	"ReserveForm/models/reserves"
	"ReserveForm/models/users"
	"ReserveForm/repositories/postgreses"
)

func main() {
	db := postgreses.Open()
	db.AutoMigrate(&users.User{}, &reserves.Reserve{}, &contents.Content{})
	db.Create(&contents.Content{Name: "テスト1"})
	db.Create(&contents.Content{Name: "テスト2"})
	db.Create(&contents.Content{Name: "テスト3"})

	servers.Start(injections.Postgres)
}
