package models

import "time"

type URL struct {
	ID             uint   `json:"id" gorm:"primary_key" `
	LongUrl        string `json:"url" binding:"required"`
	ShortUrl       string `json:"short_url"`
	Password       string
	CreatedAt      string    `json:"created_at" gorm:"autoCreateTime;type:time"`
	ExpirationDate time.Time `json:"expiration_date"`
	Visited        int       `json:"visited"`
}
