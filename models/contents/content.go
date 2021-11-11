package contents

import "gorm.io/gorm"

type Content struct {
	gorm.Model
	Name string
}
