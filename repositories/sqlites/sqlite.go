package sqlites

import (
	"github.com/jinzhu/gorm"
)

//private
func Open() *gorm.DB {
	db, _ := gorm.Open("sqlite3", "data.sqlite3")
	return db
}
