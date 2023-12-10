package main

import (
    "log"
    "net/http"
    "gohtmx/internal/app"
    "gohtmx/internal/db"
)

func main() {
    database, err := db.InitDB()
    if err != nil {
        log.Fatal(err)
    }
    defer database.Close()

    handler := app.NewHandler(database)

    http.HandleFunc("/", handler.ListTasks)
    http.HandleFunc("/add", handler.AddTask)
    http.HandleFunc("/delete", handler.DeleteTask)

    log.Println("Server started on: http://localhost:8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}

