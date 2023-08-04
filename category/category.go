package category

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

func GetAllCategory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var categories model.Category
	if err := db.Find(&categories).Error; err != nil {
		respondJSON(w, http.StatusInternalServerError, nil, "Failed to fetch categories")
		return
	}

	respondJSON(w, http.StatusOK, categories, "Categories retrieved successfully")
}

func GetCategoryByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	categoryID := params["id"]

	categoryIDUint, _ := strconv.ParseUint(categoryID, 10, 64)

	var category model.Category
	if err := db.First(&category, categoryIDUint).Error; err != nil {
		http.Error(w, "Category not found", http.StatusNotFound)
		return
	}

	respondJSON(w, http.StatusOK, category, "Category found successfully")
}

func AddCategory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var newCategory model.CategoryDto
	json.NewDecoder(r.Body).Decode(&newCategory)

	category := model.Category{
		CategoryName: newCategory.CategoryName,
	}

	if err := db.WithContext(context.Background()).Create(&category).Error; err != nil {
		http.Error(w, "Failed to create category", http.StatusInternalServerError)
		return
	}

	respondJSON(w, http.StatusCreated, newCategory, "Category added successfully")
}

func UpdateCategory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	categoryID := params["id"]
	categoryIDUint, _ := strconv.ParseUint(categoryID, 10, 64)

	var category model.Category
	if err := db.First(&category, categoryIDUint).Error; err != nil {
		respondJSON(w, http.StatusNotFound, nil, "Category not found")
		return
	}

	var updatedCategory model.CategoryDto
	json.NewDecoder(r.Body).Decode(&updatedCategory)

	category.CategoryName = updatedCategory.CategoryName

	if err := db.WithContext(context.Background()).Save(&category).Error; err != nil {
		http.Error(w, "Failed to update category", http.StatusInternalServerError)
		return
	}

	respondJSON(w, http.StatusOK, updatedCategory, "Category updated successfully")
}

// DeleteCategory deletes a category by ID from the database.
func DeleteCategory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	categoryID := params["id"]
	categoryIDUint, _ := strconv.ParseUint(categoryID, 10, 64)

	var category model.Category
	if err := db.First(&category, categoryIDUint).Error; err != nil {
		http.Error(w, "Category not found", http.StatusNotFound)
		return
	}

	if err := db.WithContext(context.Background()).Delete(&category).Error; err != nil {
		http.Error(w, "Failed to delete category", http.StatusInternalServerError)
		return
	}

	respondJSON(w, http.StatusOK, nil, "Category deleted successfully")
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
