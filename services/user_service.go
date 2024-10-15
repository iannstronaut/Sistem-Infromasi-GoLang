package services

import (
    "restapi/models"
    "gorm.io/gorm"
)

func GetUserByID(db *gorm.DB, userID uint) (*models.User, error) {
    var user models.User
    if err := db.First(&user, userID).Error; err != nil {
        return nil, err
    }
    return &user, nil
}

func UpdateUser(db *gorm.DB, userID uint, user *models.User) error {
    var existingUser models.User
    if err := db.First(&existingUser, userID).Error; err != nil {
        return err
    }
    return db.Model(&existingUser).Updates(user).Error
}

func DeleteUser(db *gorm.DB, userID uint) error {
    return db.Delete(&models.User{}, userID).Error
}
