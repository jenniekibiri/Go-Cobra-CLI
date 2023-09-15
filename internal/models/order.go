package models

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	ID         int     `gorm:"primaryKey" json:"id"`
	CustomerID int     `json:"customer_id"`
	Date       string  `json:"date"`
	QtyOrdered int     `json:"qty_ordered"`
	TotalPrice float64 `json:"total_price"`
}
