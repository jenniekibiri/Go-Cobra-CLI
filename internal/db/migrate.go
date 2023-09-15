package db
import (
	"flourish-coding-challenge/internal/models"

)

func Migrate() {
    DB.AutoMigrate(&models.Customer{}, &models.Order{})
}