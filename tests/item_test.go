package tests

import (
	"restapi/models"
	"restapi/services"
	"testing"

	"gorm.io/gorm"
)

func TestCreateItem(t *testing.T) {
    db, _ := gorm.Open("mysql", "user:password@tcp(localhost:3306)/dbname")
    item := models.Item{Title: "Item 1", Content: "This is item 1", UserID: 1}
    err := services.CreateItem(db, &item)
    if err != nil {
        t.Errorf("Failed to create item: %v", err)
    }
}
