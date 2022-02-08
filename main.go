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
	gormDB := postgreses.Open()
	gormDB.AutoMigrate(&users.User{}, &reserves.Reserve{}, &contents.Content{})
	// gormDB.Create(&contents.Content{Name: "テスト1"})
	// gormDB.Create(&contents.Content{Name: "テスト2"})
	// gormDB.Create(&contents.Content{Name: "テスト3"})
	servers.Start(injections.Postgres)
}
