package entity

import "time"

type ContentEntity struct {
	ID          int64
	Title       string
	Excerpt     string
	Description string
	Image       string
	Tags        []string
	Status      string
	CategoryID  int64
	CreatedByID int64
	CreatedAt   time.Time
	Category    CategoryEntity
	User        UserEntity
}
