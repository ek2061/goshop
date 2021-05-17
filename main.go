package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ek2061/goshop/apis/catalog_api"
	"github.com/ek2061/goshop/apis/product_api"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		j, _ := json.Marshal(map[string]string{"msg": "test ok"})
		w.Write(j)
	}).Methods("GET")

	router.HandleFunc("/api/product", product_api.FindAll).Methods("GET")
	router.HandleFunc("/api/product/{id}", product_api.Search).Methods("GET")
	router.HandleFunc("/api/product", product_api.Create).Methods("POST")
	router.HandleFunc("/api/product", product_api.Update).Methods("PUT")
	router.HandleFunc("/api/product/{id}", product_api.Delete).Methods("DELETE")

	router.HandleFunc("/api/catalog", catalog_api.FindAll).Methods("GET")
	router.HandleFunc("/api/catalog/{id}", catalog_api.Search).Methods("GET")
	router.HandleFunc("/api/catalog", catalog_api.Create).Methods("POST")
	router.HandleFunc("/api/catalog", catalog_api.Update).Methods("PUT")
	router.HandleFunc("/api/catalog/{id}", catalog_api.Delete).Methods("DELETE")

	err := http.ListenAndServe(":5000", router)
	if err != nil {
		fmt.Println(err)
	}
}
