package main

import (
	"GoProjects/Project2/product"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	// database.InitialiseDb()

	router := mux.NewRouter()

	productRouter := router.PathPrefix("/product").Subrouter()
	productRouter.HandleFunc("", product.GetAllProduct).Methods(http.MethodGet)
	productRouter.HandleFunc("/{id}", product.GetProductByID).Methods(http.MethodGet)
	productRouter.HandleFunc("/add", product.AddProduct).Methods(http.MethodPost)
	productRouter.HandleFunc("/update/{id}", product.UpdateProduct).Methods(http.MethodPut)
	productRouter.HandleFunc("/delete/{id}", product.DeleteProduct).Methods(http.MethodDelete)

	// homeRouter := router.PathPrefix("/search").Subrouter()
	// homeRouter.HandleFunc("", search.Search).Methods(http.MethodGet)

	log.Fatal(http.ListenAndServe(":8080", router))
}
