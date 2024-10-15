package models

type User struct {
    ID       uint   `gorm:"primaryKey" json:"id"`
    Name     string `json:"name"`
    Username string `gorm:"unique" json:"username"`
    Email    string `gorm:"unique" json:"email"`
    Password string `json:"-"`
    Items    []Item `gorm:"foreignKey:UserID"`
}
