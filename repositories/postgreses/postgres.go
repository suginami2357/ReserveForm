package postgreses

import (

	// "github.com/jinzhu/gorm"

	"database/sql"
	"os"

	"github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

//private
func Open() *gorm.DB {

	//HerokuでDBのセットアップをする場合、 DATABASE_URL という環境変数が設定される
	url := os.Getenv("DATABASE_URL")
	connection, _ := pq.ParseURL(url)
	connection += " sslmode=require"
	sqlDB, _ := sql.Open("postgres", connection)
	gormDB, _ := gorm.Open(postgres.New(postgres.Config{
		Conn: sqlDB,
	}), &gorm.Config{})

	// connection, _ := pq.ParseURL(os.Getenv("DATABASE_URL"))
	// // connection += " host=postgres user=postgres password=password dbname=postgres sslmode=disable"
	// connection += "host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai"
	// sqlDB, _ := sql.Open("postgres", connection)
	// gormDB, _ := gorm.Open(postgres.New(postgres.Config{
	// 	Conn: sqlDB,
	// }), &gorm.Config{})

	// dsn := "host=localhost user=postgres password=password dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	// gormDB, _ := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	return gormDB
}
