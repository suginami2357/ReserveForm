package reserves

import (
	"github.com/jinzhu/gorm"
)

type Reserve struct {
	gorm.Model
	UserID    uint
	PlaceID   uint
	PlaceName string
	Date      string
}

func (r Reserve) Format_MMDD() string {
	s := []rune(r.Date) //UTF-8のsliceにキャスト
	return string(s[5:7]) + "月" + string(s[8:10]) + "日"
}
