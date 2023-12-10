package db

import (
    "database/sql"
    _ "github.com/mattn/go-sqlite3"
)

func InitDB() (*sql.DB, error) {
    db, err := sql.Open("sqlite3", "./todo.db")
    if err != nil {
        return nil, err
    }

    createTableSQL := `CREATE TABLE IF NOT EXISTS tasks (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        task TEXT NOT NULL,
        completed BOOLEAN NOT NULL DEFAULT 0
    );`
    _, err = db.Exec(createTableSQL)
    if err != nil {
        return nil, err
    }

    return db, nil
}
