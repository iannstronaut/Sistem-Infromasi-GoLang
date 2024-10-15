package main

import (
	"fmt"
	"log"
	"os"
	"restapi/controllers"
	"restapi/middlewares"
	"restapi/models"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbUser := os.Getenv("DB_USERNAME")
	dbPass := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_DBNAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", 
		dbUser, dbPass, dbHost, dbPort, dbName)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database")
	}

	db.AutoMigrate(&models.User{}, &models.Item{})

	app := fiber.New()

	app.Post("/register", controllers.Register(db))
	app.Post("/login", controllers.Login(db))

	api := app.Group("/api", middlewares.AuthMiddleware)
	api.Get("/user", controllers.GetUser(db))
	api.Put("/user", controllers.UpdateUser(db))
	api.Delete("/user", controllers.DeleteUser(db))
	api.Post("/logout", controllers.Logout(db))
	api.Get("/item", controllers.GetItems(db))
	api.Put("/item", controllers.AddItem(db))
	api.Put("/item/:id", controllers.UpdateItem(db))
	api.Delete("/item/:id", controllers.DeleteItem(db))

	log.Fatal(app.Listen(":8080"))
}
