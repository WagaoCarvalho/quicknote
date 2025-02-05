package models

import (
	"database/sql"
	"time"
)

type Note struct {
	Id        int
	Title     string
	Content   string
	Color     string
	CreatedAt time.Time
	UpdatedAt sql.NullTime
}
