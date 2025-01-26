// handlers/product.go
package main

import (
	"encoding/json"
	"net/http"
)

type Product struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
	Image string  `json:"image"`
}

func GetProducts(w http.ResponseWriter, r *http.Request) {
	products := []Product{
		{ID: 1, Name: "Dumbbell", Price: 20.00, Image: "/images/dumbbell.jpg"},
		{ID: 2, Name: "Yoga Mat", Price: 15.00, Image: "/images/yogamat.jpg"},
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}
