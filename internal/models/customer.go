package models
import "gorm.io/gorm"

type Customer struct {
	gorm.Model
    ID    int    `gorm:"primaryKey" json:"id"`
    Name  string `json:"name"`
    Email string `json:"email" gorm:"unique"`
}