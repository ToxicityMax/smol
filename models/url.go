package models

type Url struct {
	ID       uint   `json:"id" gorm:"primary_key" `
	LongUrl  string `json:"long_url" binding:"required"`
	ShortUrl string `json:"short_url"`
}
