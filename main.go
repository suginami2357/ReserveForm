package main

import (
	// "github.com/jinzhu/ gorm"

	"ReserveForm/commons/injections"
	"ReserveForm/commons/servers"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// db, _ := gorm.Open("sqlite3", "data.sqlite3")
	// defer db.Close()
	// db.AutoMigrate(&users.User{}, &reserves.Reserve{}, &places.Place{})
	// db.Create(&places.Place{Name: "テスト1"})
	// db.Create(&places.Place{Name: "テスト2"})
	// db.Create(&places.Place{Name: "テスト3"})

	servers.Start(injections.Sqlite)
}
