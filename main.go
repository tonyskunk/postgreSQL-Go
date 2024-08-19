package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "1234"
	dbname   = "mydb"
)

func main() {
	// Формирование строки подключения
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	// Подключение к базе данных
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatalf("Ошибка подключения: %v", err)
	}
	defer db.Close()

	// Проверка подключения
	err = db.Ping()
	if err != nil {
		log.Fatalf("Ошибка при проверке подключения: %v", err)
	}
	fmt.Println("Успешное подключение к базе данных!")

	// Вставка данных
	sqlStatement := `INSERT INTO users (name, age) VALUES ($1, $2)`
	_, err = db.Exec(sqlStatement, "Alice", 30)
	if err != nil {
		log.Fatalf("Ошибка вставки данных: %v", err)
	}

	fmt.Println("Данные успешно вставлены!")

	// Чтение данных
	var (
		id   int
		name string
		age  int
	)
	query := `SELECT id, name, age FROM users WHERE name=$1`
	row := db.QueryRow(query, "Alice")
	err = row.Scan(&id, &name, &age)
	if err != nil {
		log.Fatalf("Ошибка при чтении данных: %v", err)
	}
	fmt.Printf("Результат: id=%d, name=%s, age=%d\n", id, name, age)
}
