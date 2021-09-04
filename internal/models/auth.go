package models

import (
	"time"

	"gorm.io/gorm"
)

type Login struct {
	gorm.Model
	Username  string
	Password  string
	LoginAs   int
	LastLogin time.Time
}
