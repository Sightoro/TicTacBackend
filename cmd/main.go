package main

import (
	"fmt"
	"net/http"

	"game_backend/internal/api"
	"game_backend/internal/db"

	"github.com/go-chi/chi/v5"
)

func main() {
	// Инициализация базы данных
	db.InitDB()

	// Создаем роутер
	r := chi.NewRouter()

	// Регистрация маршрутов
	r.Post("/room/create", api.CreateRoom)

	fmt.Println("Сервер запущен на порту 8080")
	http.ListenAndServe(":8080", r)
}
