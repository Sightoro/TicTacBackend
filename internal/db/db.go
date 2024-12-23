package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq" //Подключаем драйвер PotgreSQL
)

var DB *sql.DB

func InitDB() {
	connStr := "host=localhost port=5432 user=game_user password=game_password dbname=game_db sslmode=disable"
	var err error
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Ошибка подключения к базе данных: %v", err)
	}

	//Провверяем соединение
	err = DB.Ping()
	if err != nil {
		log.Fatalf("Не удалось подключиться к базе данных: %v", err)
	}
	fmt.Println("Удачное подключение к базе данных!")
}
