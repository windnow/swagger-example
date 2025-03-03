# 🚀 Go Mux Swagger API  

Пример API на **Go** с `gorilla/mux` и `Swagger`.  

## 🔹 Установка зависимостей  
```sh
go mod tidy
```

## 🔹 Генерация Swagger-документации  
```sh
go install github.com/swaggo/swag/cmd/swag@latest  
swag init
```

## 🔹 Запуск сервера  
```sh
go run main.go
```

## 🔹 Доступные эндпоинты  

📌 **Swagger UI:** [`http://localhost:8080/swagger/index.html`](http://localhost:8080/swagger/index.html)  

📌 **API:**  
- `GET /api/products` — список товаров  
- `POST /api/products/decrease` — списание товара  
- `POST /api/products/increase` — оприходование товара  

## 🔹 Примеры запросов  

📌 **Получить список товаров:**  
```sh
curl -X GET "http://localhost:8080/api/products"
```

📌 **Списать товар:**  
```sh
curl -X POST "http://localhost:8080/api/products/decrease" \
     -H "Content-Type: application/json" \
     -d '{"id": 1, "quantity": 2}'
```

📌 **Оприходовать товар:**  
```sh
curl -X POST "http://localhost:8080/api/products/increase" \
     -H "Content-Type: application/json" \
     -d '{"id": 1, "quantity": 5}'
```