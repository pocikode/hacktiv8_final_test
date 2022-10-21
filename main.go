package main

import (
	"MyGram/database"
	"MyGram/server"
	"MyGram/server/controllers"
	"MyGram/server/repositories/gorm"
	"github.com/joho/godotenv"
	"log"
	"os"
)

// @description    MyGram - Hacktiv8 Final Test
// @termsOfService http://swagger.io/terms/
// @BasePath       /
// @contact.name   Agus Supriyatna
// @contact.email  supriyatnaagus@outlook.com
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization

func init() {
	if _, err := os.Stat(".env"); err == nil {
		err = godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}
}

func main() {
	DB := database.ConnectDB()

	photoRepo := gorm.NewPhotoRepo(DB)
	photoController := controllers.NewPhotoController(photoRepo)

	socialMediaRepo := gorm.NewSocialMediaRepo(DB)
	socialMediaController := controllers.NewSocialMediaController(socialMediaRepo, photoRepo)

	commentRepo := gorm.NewCommentRepo(DB)
	commentController := controllers.NewCommentController(commentRepo)

	userRepo := gorm.NewUserRepo(DB)
	userController := controllers.NewUserController(userRepo, photoRepo, commentRepo, socialMediaRepo)

	router := server.NewRouter(userController, photoController, socialMediaController, commentController)
	router.Start(":" + os.Getenv("PORT"))
}
