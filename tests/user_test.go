package tests

import (
    "restapi/models"
    "restapi/services"
    "gorm.io/gorm"
    "testing"
)

func TestCreateUser(t *testing.T) {
    db, _ := gorm.Open("mysql", "user:password@tcp(localhost:3306)/dbname")
    user := models.User{Name: "John Doe", Username: "john", Email: "john@example.com", Password: "password"}
    err := services.CreateUser(db, &user)
    if err != nil {
        t.Errorf("Failed to create user: %v", err)
    }
}
