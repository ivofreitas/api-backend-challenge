package model

import "time"

type User struct {
	Username  string
	Role      string
	ManagedBy string
	CreateAt  time.Time
	UpdatedAt time.Time
}
