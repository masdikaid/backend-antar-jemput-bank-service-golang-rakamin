package entities

import (
	"database/sql"

	"gorm.io/gorm"
)

type Login struct {
	gorm.Model
	Username  string
	Password  string
	LoginAs   int
	LastLogin sql.NullTime
}
