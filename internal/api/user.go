package api

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"game_backend/internal/db"
	"game_backend/internal/models"
)

func RegisterUser(w http.ResponseWriter, r *http.Request) {
	var user models.User

	// Decode JSON - запрос
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Неверный формат данных", http.StatusBadRequest)
		return
	}

	// Adding Users in DB
	query := `INSERT INTO users (username, password) VALUES ($1, $2) RETURNING id`
	err = db.DB.QueryRow(query, user.Username, user.Password).Scan(&user.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Не удаллось создать пользователя", http.StatusInternalServerError)
		} else {
			http.Error(w, "Ошибка баз данных", http.StatusInternalServerError)
		}
		return
	}

	//Возвращаем созданного пользователя
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}
