package models

import "time"

type BasicModel struct {
	ID        uint
	CreatedAt time.Time
	UpdatedAt time.Time
}
