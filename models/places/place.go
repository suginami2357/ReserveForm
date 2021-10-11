package places

import "gorm.io/gorm"

type Place struct {
	gorm.Model
	Name string
}
