package postgreses

import (

	// "github.com/jinzhu/gorm"

	"os"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/stdlib"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

//private
func Open() *gorm.DB {

	config, _ := pgx.ParseConfig(os.Getenv("DATABASE_URL"))
	dialector := postgres.New(postgres.Config{Conn: stdlib.OpenDB(*config)})
	opts := &gorm.Config{Logger: logger.Discard}
	gormDB, _ := gorm.Open(dialector, opts)

	// dsn, _ := pq.ParseURL(os.Getenv("DATABASE_URL"))
	// dsn += "user=postgres dbname=password password=password"
	// sqlDB, _ := sql.Open("postgres", dsn)
	// // config := postgres.Config{
	// // 	DriverName: "postgres",
	// // 	DSN: dsn,
	// // 	Conn: ,
	// // }
	// dialector := postgres.New(postgres.Config{DriverName: "postgres", Conn: sqlDB})
	// gormDB, _ := gorm.Open(dialector, &gorm.Config{})

	//HerokuでDBのセットアップをする場合 DATABASE_URL という環境変数が設定される
	// url := os.Getenv("DATABASE_URL")
	// connection, _ := pq.ParseURL(url)
	// connection += " host=postgres user=postgres password=password dbname=postgres sslmode=disable"
	// connection += "user=postgres dbname=password password=password sslmode=disable"
	// sqlDB, _ := sql.Open("postgres", connection)
	// gormDB, _ := gorm.Open(postgres.New(postgres.Config{
	// 	Conn: sqlDB,
	// }), &gorm.Config{})

	// dsn := "host=localhost user=postgres password=password dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	// gormDB, _ := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	return gormDB
}
