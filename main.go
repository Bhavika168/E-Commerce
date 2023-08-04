package main

import (
	"GoProjects/Project2/product"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	product.InitialiseDb()

	router := mux.NewRouter()

	productRouter := router.PathPrefix("/product").Subrouter()
	productRouter.HandleFunc("", product.GetAllProduct).Methods(http.MethodGet)
	productRouter.HandleFunc("/get", product.GetByIdProduct).Methods(http.MethodGet)
	productRouter.HandleFunc("/add", product.AddProduct).Methods(http.MethodPost)
	productRouter.HandleFunc("/update", product.UpdateProduct).Methods(http.MethodPut)
	productRouter.HandleFunc("/delete", product.DeleteProduct).Methods(http.MethodDelete)

	productRouter.HandleFunc("/data", product.GetData).Methods(http.MethodGet)

	// homeRouter := router.PathPrefix("/search").Subrouter()
	// homeRouter.HandleFunc("", search.Search).Methods(http.MethodGet)

	log.Fatal(http.ListenAndServe(":8080", router))
}
