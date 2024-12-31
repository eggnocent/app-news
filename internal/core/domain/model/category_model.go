package model

import "time"

type Category struct {
	ID          int64     `gorm:"id"`
	Title       string    `gorm:"title"`
	Slug        string    `gorm:"slug"`
	CreatedByID int64     `gorm:"created_by_id"`
	User        User      `gorm:"foreignKey:CreatedByID"`
	CreatedAt   time.Time `gorm:"created_at"`
	UpdatedAt   time.Time `gorm:"updated_at"`
}
