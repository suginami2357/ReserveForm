package main

import (
	"ReserveForm/commons/injections"
	"ReserveForm/commons/servers"
	// "github.com/jinzhu/gorm"
	// _ "github.com/mattn/go-sqlite3"
)

func main() {
	// db, _ := gorm.Open("sqlite3", "data.sqlite3")
	// defer db.Close()

	// user := users.User{}
	// db.Where("email = ?", "jinzhu").First(&user)

	// db.AutoMigrate(&users.User{}, &reserves.Reserve{}, &contents.Content{})
	// db.Create(&contents.Content{Name: "テスト1"})
	// db.Create(&contents.Content{Name: "テスト2"})
	// db.Create(&contents.Content{Name: "テスト3"})

	servers.Start(injections.Sqlite)
}
