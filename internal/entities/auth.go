package entities

import (
	"database/sql"
	"time"

	"gorm.io/gorm"
)

type Login struct {
	gorm.Model
	Username  string
	Password  string
	LoginAs   int
	LastLogin sql.NullTime
}

type Session struct {
	gorm.Model
	LoginID   uint
	Login     Login
	SID       string
	ExpiredAt time.Time
}
