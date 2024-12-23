package api

import (
	"encoding/json"
	"net/http"

	"game_backend/internal/db"
	"game_backend/internal/models"
)

// CreateRoom - обработчик для создания комнаты
func CreateRoom(w http.ResponseWriter, r *http.Request) {
	var room models.Room

	// Проверяем метод запроса
	if r.Method != http.MethodPost {
		http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
		return
	}

	// Декодируем JSON из тела запроса
	err := json.NewDecoder(r.Body).Decode(&room)
	if err != nil {
		http.Error(w, "Неверный формат JSON", http.StatusBadRequest)
		return
	}

	// Логика для добавления комнаты в базу данных (добавим позже)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Комната успешно создана"))
	query := `INSERT INTO rooms (name, host_id) VALUES ($1, $2) RETURNING id`
	err = db.DB.QueryRow(query, room.Name, room.HostID).Scan(&room.ID)
	if err != nil {
		http.Error(w, "Ошибка базы данных", http.StatusInternalServerError)
		return
	}
}
