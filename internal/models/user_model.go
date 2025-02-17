package models

import (
	"database/sql"
	"time"
)

type User struct {
	Id        int
	Email     string
	Password  string
	Active    bool
	CreatedAt time.Time
	UpdatedAt sql.NullTime
}
