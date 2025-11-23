package models

import "time"

type Car struct {
	ID        uint   "gorm:\"primaryKey\""
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
