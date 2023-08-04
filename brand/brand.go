package brand

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

func GetAllBrand(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var brands model.Brand
	if err := db.Find(&brands).Error; err != nil {
		respondJSON(w, http.StatusInternalServerError, nil, "Failed to fetch brands")
		return
	}

	respondJSON(w, http.StatusOK, brands, "Brands retrieved successfully")
}

func GetBrandByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	brandID := params["id"]

	brandIDUint, _ := strconv.ParseUint(brandID, 10, 64)

	var brand model.Brand
	if err := db.First(&brand, brandIDUint).Error; err != nil {
		http.Error(w, "Brand not found", http.StatusNotFound)
		return
	}

	respondJSON(w, http.StatusOK, brand, "Brand found successfully")
}

func AddBrand(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var newBrand model.BrandDto
	json.NewDecoder(r.Body).Decode(&newBrand)

	brand := model.Brand{
		BrandName: newBrand.BrandName,
	}

	if err := db.WithContext(context.Background()).Create(&brand).Error; err != nil {
		http.Error(w, "Failed to create brand", http.StatusInternalServerError)
		return
	}

	respondJSON(w, http.StatusCreated, newBrand, "Brand added successfully")
}

func UpdateBrand(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	brandID := params["id"]
	brandIDUint, _ := strconv.ParseUint(brandID, 10, 64)

	var brand model.Brand
	if err := db.First(&brand, brandIDUint).Error; err != nil {
		respondJSON(w, http.StatusNotFound, nil, "Brand not found")
		return
	}

	var updatedBrand model.BrandDto
	json.NewDecoder(r.Body).Decode(&updatedBrand)

	brand.BrandName = updatedBrand.BrandName

	if err := db.WithContext(context.Background()).Save(&brand).Error; err != nil {
		http.Error(w, "Failed to update brand", http.StatusInternalServerError)
		return
	}

	respondJSON(w, http.StatusOK, updatedBrand, "Brand updated successfully")
}

func DeleteBrand(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	brandID := params["id"]
	brandIDUint, _ := strconv.ParseUint(brandID, 10, 64)

	var brand model.Brand
	if err := db.First(&brand, brandIDUint).Error; err != nil {
		http.Error(w, "Brand not found", http.StatusNotFound)
		return
	}

	if err := db.WithContext(context.Background()).Delete(&brand).Error; err != nil {
		http.Error(w, "Failed to delete brand", http.StatusInternalServerError)
		return
	}

	respondJSON(w, http.StatusOK, nil, "Brand deleted successfully")
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
