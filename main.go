package main

import (
	"ReserveForm/commons/injections"
	"ReserveForm/commons/servers"
)

func main() {
	// dsn := "host=localhost user=postgres password=password dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	// gormDB, _ := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	// gormDB.AutoMigrate(&users.User{}, &reserves.Reserve{}, &contents.Content{})
	// gormDB.Create(&contents.Content{Name: "テスト1"})
	// gormDB.Create(&contents.Content{Name: "テスト2"})
	// gormDB.Create(&contents.Content{Name: "テスト3"})
	servers.Start(injections.Postgres)
}
