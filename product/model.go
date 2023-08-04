package product

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	ID          uint   `json:"id" gorm:"primaryKey"`
	ProductName string `json:"productname" gorm:"not null"`
	Stock       string `json:"stock"`
	Brand       string `json:"brand"`
}

type ProductDto struct {
	ProductName string `json:"productname"`
	Stock       string `json:"stock"`
	Brand       string `json:"brand"`
}

type Productget struct {
	ProductName string `json:"productname"`
}
