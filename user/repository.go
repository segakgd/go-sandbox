package user

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

func ConnectAndCreateUser() {
	db, err := sql.Open("sqlite3", "test.db")

	if err != nil {
		panic(err)
	}

	defer db.Close()

	createTableIfNotExist(db)
	createUser(db)

	var rows = getUsers(db)

	defer rows.Close()

	for rows.Next() {
		var id int
		var name string
		var age int
		var height float64
		err = rows.Scan(&id, &name, &age, &height)
		if err != nil {
			panic(err)
		}
		fmt.Printf("ID: %d | %s, %d лет, рост %.2f\n", id, name, age, height)
	}
}

func getUsers(db *sql.DB) *sql.Rows {
	// Чтение данных
	rows, err := db.Query("SELECT id, name, age, height FROM users")

	if err != nil {
		panic(err)
	}

	return rows
}

func createUser(db *sql.DB) {
	_, err := db.Exec("INSERT INTO users (name, age, height) VALUES (?, ?, ?)", "User name", 25, 180.5)

	if err != nil {
		panic(err)
	}
}

func createTableIfNotExist(db *sql.DB) {
	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS users (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        name TEXT,
        age INTEGER,
        height REAL
    )`)

	if err != nil {
		panic(err)
	}
}
