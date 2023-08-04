package product

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func GetAllProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var products []Product
	if err := db.Find(&products).Error; err != nil {
		respondJSON(w, http.StatusInternalServerError, nil, "Failed to fetch products")
		return
	}

	respondJSON(w, http.StatusOK, products, "Products retrieved successfully")
	// json.NewEncoder(w).Encode(product)
}

// GetByIdProduct gets a product by ID from the database.
func GetByIdProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var product Product
	json.NewDecoder(r.Body).Decode(&product)
	result := db.Where("productname = ?", product.ProductName).First(&product)

	if result.Error != nil {
		respondJSON(w, http.StatusNotFound, nil, "Product not found")
	} else {
		respondJSON(w, http.StatusOK, product, "Product retrieved successfully")
	}

	// params := mux.Vars(r)
	// productID := params["id"]

	// var product Product
	// if err := db.First(&product, productID).Error; err != nil {
	// 	respondJSON(w, http.StatusNotFound, nil, "Product not found")
	// 	return
	// }

	// respondJSON(w, http.StatusOK, product, "Product retrieved successfully")
}

// AddProduct adds a new product to the database.
func AddProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var newProduct ProductDto
	if err := json.NewDecoder(r.Body).Decode(&newProduct); err != nil {
		http.Error(w, "Failed to decode JSON", http.StatusBadRequest)
		return
	}

	product := Product{ProductName: newProduct.ProductName, Stock: newProduct.Stock, Brand: newProduct.Brand}
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
	productID := params["id"]

	var product Product
	if err := db.First(&product, productID).Error; err != nil {
		respondJSON(w, http.StatusNotFound, nil, "Product not found")
		return
	}

	var updatedProduct ProductDto
	if err := json.NewDecoder(r.Body).Decode(&updatedProduct); err != nil {
		http.Error(w, "Failed to decode JSON", http.StatusBadRequest)
		return
	}

	product.ProductName = updatedProduct.ProductName
	product.Stock = updatedProduct.Stock
	product.Brand = updatedProduct.Brand

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
	productID := params["productid"]

	var product Product
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

// Helper function
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

func GetData(w http.ResponseWriter, r *http.Request) {

	db.AutoMigrate(&Product{})
	var products []Product
	db.Find(&products)

	// Return the products as JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)

}
