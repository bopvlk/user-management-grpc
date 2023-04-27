package models

import (
	"database/sql"
	"time"
)

type User struct {
	ID        uint
	FirstName string
	LastName  string
	Email     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeleteAt  sql.NullTime
}
