package postgreses

import (

	// "github.com/jinzhu/gorm"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	_ "github.com/lib/pq"
)

//private
func Open() *gorm.DB {
	// mydb_dsn := os.Getenv("DATABASE_URL")
	// connection, _ := pq.ParseURL(mydb_dsn)
	// connection += " sslmode=require"
	// sqlDB, _ := sql.Open("postgres", connection)
	// gormDB, _ := gorm.Open(postgres.New(postgres.Config{
	// 	Conn: sqlDB,
	// }), &gorm.Config{})

	dsn := "host=localhost user=postgres password=password dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	gormDB, _ := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	return gormDB
}
