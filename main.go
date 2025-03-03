package main

import (
	"encoding/json"
	"net/http"

	_ "go-swagger-api/docs"

	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
)

type Product struct {
	ID    int     `json:"id" example:"1"`
	Name  string  `json:"name" example:"Телефон"`
	Price float64 `json:"price" example:"43200.0"`
	Stock int     `json:"stock" example:"15"`
}

var products = []Product{
	{ID: 1, Name: "Телефон", Price: 50000, Stock: 10},
	{ID: 2, Name: "Ноутбук", Price: 120000, Stock: 5},
	{ID: 3, Name: "Наушники", Price: 7000, Stock: 25},
}

// @title Go Mux Swagger API
// @version 1.0
// @description API для управления товарами
// host localhost:8080
func main() {
	r := mux.NewRouter()

	api := r.PathPrefix("/api").Subrouter()
	api.HandleFunc("/products", GetProductsHandler()).Methods("GET")
	// api.HandleFunc("/products/dec", DecreaseStockHandler()).Methods("POST")
	// api.HandleFunc("/products/inc", IncreaseStockHandler()).Methods("POST")

	r.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	http.ListenAndServe(":8080", r)
}

// GetProductsHandler Возвращает список товаров
// @Summary Список товаров
// @Description Получение списка товаров
// @Tags products
// @Produce json
// @Success 200 {array} Product
// @Router /api/products [get]
func GetProductsHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(products)
	}
}

type good struct {
	ID       int `json:"id" example:"34"`
	Quantity int `json:"quantity" example:"1"`
}
type resp map[string]string
type Status struct {
	M string `json:"message" example:"Успех"`
}

// DecreaseStockHandler списывает товар
// @Summary Списание товара
// @Description Уменьшение количества товара на складе
// @Tags products
// @Accept json
// @Produce json
// @Param request body good true "Данные для списания"
// @Param ignore_stock query bool false "Игнорировать контроль остатков"
// @Success 200 {object} Status "Товар списан"
// @Failure 400 {string} string "Ошибка в параметрах"
// @Router /api/products/decrease [post]
func DecreaseStockHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req good

		// Декодируем JSON-запрос
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Неверный JSON", http.StatusBadRequest)
			return
		}

		// Поиск товара
		for i, p := range products {
			if p.ID == req.ID {
				if p.Stock < req.Quantity {
					http.Error(w, "Недостаточно товара", http.StatusBadRequest)
					return
				}
				products[i].Stock -= req.Quantity
				json.NewEncoder(w).Encode(resp{"message": "Товар списан"})
				return
			}
		}

		http.Error(w, "Товар не найден", http.StatusBadRequest)
	}
}

// IncreaseStockHandler оприходует товар
// @Summary Оприходование товара
// @Description Увеличение количества товара на складе
// @Tags products
// @Accept json
// @Produce json
// @Param request body good true "Данные для оприходования"
// @Success 200 {string} string "Товар оприходован"
// @Failure 400 {string} string "Ошибка в параметрах"
// @Router /api/products/increase [post]
func IncreaseStockHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req struct {
			ID       int `json:"id"`
			Quantity int `json:"quantity"`
		}

		// Декодируем JSON-запрос
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Неверный JSON", http.StatusBadRequest)
			return
		}

		// Поиск товара
		for i, p := range products {
			if p.ID == req.ID {
				products[i].Stock += req.Quantity
				json.NewEncoder(w).Encode(resp{"message": "Товар оприходован"})
				return
			}
		}

		http.Error(w, "Товар не найден", http.StatusBadRequest)
	}
}
