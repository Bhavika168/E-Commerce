package product

import (
	"GoProjects/Project2/database"
	"GoProjects/Project2/model"
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var db = database.InitialiseDb()

func GetAllProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var products model.Product
	if err := db.Find(&products).Error; err != nil {
		respondJSON(w, http.StatusInternalServerError, nil, "Failed to fetch products")
		return
	}

	respondJSON(w, http.StatusOK, products, "Products retrieved successfully")
}

func GetProductByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	productId := params["id"]

	productID, _ := strconv.ParseUint(productId, 10, 64)

	var product model.Product
	if err := db.First(&product, productID).Error; err != nil {
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}

	respondJSON(w, http.StatusOK, product, "Product found successfully")
}

func AddProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var newProduct model.ProductDto
	json.NewDecoder(r.Body).Decode(&newProduct)

	product := model.Product{
		ProductName: newProduct.ProductName,
		Price:       newProduct.Price,
		CategoryId:  newProduct.CategoryId,
		BrandId:     newProduct.BrandId,
		Stock:       newProduct.Stock,
		AvgRating:   newProduct.AvgRating,
		Location:    newProduct.Location,
		Description: newProduct.Description,
	}
	if err := db.WithContext(context.Background()).Create(&product).Error; err != nil {
		http.Error(w, "Failed to create product", http.StatusInternalServerError)
		return
	}

	respondJSON(w, http.StatusCreated, product, "Product added successfully")
}

// UpdateProduct updates an existing product in the database.
func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	productId := params["id"]
	productID, _ := strconv.ParseUint(productId, 10, 64)

	var product model.Product
	if err := db.First(&product, productID).Error; err != nil {
		respondJSON(w, http.StatusNotFound, nil, "Product not found")
		return
	}

	var updatedProduct model.ProductDto
	json.NewDecoder(r.Body).Decode(&updatedProduct)

	product.ProductName = updatedProduct.ProductName
	product.Price = updatedProduct.Price
	product.CategoryId = updatedProduct.CategoryId
	product.BrandId = updatedProduct.BrandId
	product.Stock = updatedProduct.Stock
	product.AvgRating = updatedProduct.AvgRating
	product.Location = updatedProduct.Location
	product.Description = updatedProduct.Description

	if err := db.WithContext(context.Background()).Save(&product).Error; err != nil {
		http.Error(w, "Failed to update product", http.StatusInternalServerError)
		return
	}

	respondJSON(w, http.StatusOK, product, "Product updated successfully")
}

// DeleteProduct deletes a product by ID from the database.
func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	productId := params["id"]
	productID, _ := strconv.ParseUint(productId, 10, 64)

	var product model.Product
	if err := db.First(&product, productID).Error; err != nil {
		respondJSON(w, http.StatusNotFound, nil, "Product not found")
		return
	}

	if err := db.WithContext(context.Background()).Delete(&product).Error; err != nil {
		http.Error(w, "Failed to delete product", http.StatusInternalServerError)
		return
	}

	respondJSON(w, http.StatusOK, nil, "Product deleted successfully")
}

func respondJSON(w http.ResponseWriter, statusCode int, data interface{}, message string) {
	w.WriteHeader(statusCode)

	response := map[string]interface{}{
		"data":    data,
		"message": message,
	}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Failed to encode JSON", http.StatusInternalServerError)
		return
	}
}
