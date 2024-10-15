package services

import (
    "restapi/models"
    "restapi/utils"
    "gorm.io/gorm"
    "golang.org/x/crypto/bcrypt"
)

func CreateUser(db *gorm.DB, user *models.User) error {
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
    if err != nil {
        return err
    }
    user.Password = string(hashedPassword)
    return db.Create(user).Error
}

func LoginUser(db *gorm.DB, username, password string) (string, error) {
    var user models.User
    if err := db.Where("username = ?", username).First(&user).Error; err != nil {
        return "", err
    }

    if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
        return "", err
    }

    token, err := utils.GenerateToken(user.ID)
    if err != nil {
        return "", err
    }

    return token, nil
}

func BlacklistToken(token string) error {
    // Implement blacklist logic using Redis
    return nil
}
