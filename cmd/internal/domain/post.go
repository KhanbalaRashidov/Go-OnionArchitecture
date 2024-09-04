package domain

import "time"

type Post struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	Tags      []Tag     `json:"tags" gorm:"many2many:post_tags;"`
	Likes     int       `json:"likes"`
	UserID    uint      `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Duration  int       `json:"duration"` // dakika cinsinden
}
