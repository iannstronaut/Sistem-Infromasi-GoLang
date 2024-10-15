package services

import (
    "restapi/models"
    "gorm.io/gorm"
)

func GetItemsByUserID(db *gorm.DB, userID uint) ([]models.Item, error) {
    var items []models.Item
    err := db.Where("user_id = ?", userID).Find(&items).Error
    return items, err
}

func CreateItem(db *gorm.DB, item *models.Item) error {
    return db.Create(item).Error
}

func UpdateItem(db *gorm.DB, itemID string, userID uint, item *models.Item) error {
    var existingItem models.Item
    if err := db.Where("id = ? AND user_id = ?", itemID, userID).First(&existingItem).Error; err != nil {
        return err
    }
    return db.Model(&existingItem).Updates(item).Error
}

func DeleteItem(db *gorm.DB, itemID string, userID uint) error {
    var item models.Item
    if err := db.Where("id = ? AND user_id = ?", itemID, userID).First(&item).Error; err != nil {
        return err
    }
    return db.Delete(&item).Error
}
