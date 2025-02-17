package models

import (
	"database/sql"
	"time"
)

type UserConfirmationToken struct {
	Id        int
	UserId    int
	Token     int
	Confirmed bool
	CreatedAt time.Time
	UpdatedAt sql.NullTime
}
