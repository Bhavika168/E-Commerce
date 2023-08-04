package model

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	ProductId   uint     `json:"id" gorm:"primaryKey"`
	ProductName string   `json:"productname" gorm:"not null"`
	Price       float64  `json:"price"`
	CategoryId  uint     `json:"categoryid"`
	Category    Category `gorm:"foreignKey:CategoryId"`
	BrandId     uint     `json:"brandid"`
	Brand       Brand    `gorm:"foreignKey:BrandId"`
	Stock       string   `json:"stock"`
	AvgRating   float64  `json:"avgrating"`
	Location    string   `json:"location"`
	Description string   `json:"description"`
}

type Category struct {
	gorm.Model
	CategoryId   uint   `json:"categoryid" gorm:"primaryKey"`
	CategoryName string `json:"categoryname"`
}

type Brand struct {
	gorm.Model
	BrandId   uint   `json:"brandid" gorm:"primaryKey"`
	BrandName string `json:"brandname"`
}

type ProductDto struct {
	ProductName string  `json:"productname"`
	Price       float64 `json:"price"`
	CategoryId  uint    `json:"categoryid"`
	BrandId     uint    `json:"brandid"`
	Stock       string  `json:"stock"`
	AvgRating   float64 `json:"avgrating"`
	Location    string  `json:"location"`
	Description string  `json:"description"`
}

type CategoryDto struct {
	CategoryName string `json:"categoryname"`
}

type BrandDto struct {
	BrandName string `json:"brandname"`
}
