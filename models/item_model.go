package models

type Item struct {
    ID      uint   `gorm:"primaryKey" json:"id"`
    Title   string `json:"title"`
    Content string `json:"content"`
    UserID  uint   `json:"user_id"`
}
