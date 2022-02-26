package postgreses

import (

	// "github.com/jinzhu/gorm"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

//private
func Open() *gorm.DB {
	// dsn := "host=localhost user=all password=password dbname=all port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	dsn := "postgres://chyfofnwozqyct:21e7e4cc5818d2a03ea783d235aa3131ef730e0d697196ede51be37fc411d05e@ec2-18-215-8-186.compute-1.amazonaws.com:5432/d7vogf43k1o4kp"
	gormDB, _ := gorm.Open(postgres.Open(dsn))

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
