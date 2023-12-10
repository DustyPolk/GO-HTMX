package app

import (
	"database/sql"
	"html/template"
	"net/http"
)

// Handler struct holds the DB connection and templates.
type Handler struct {
	DB   *sql.DB
	Tmpl *template.Template
}

// NewHandler creates a new Handler instance.
func NewHandler(db *sql.DB) *Handler {
	return &Handler{
		DB:   db,
		Tmpl: template.Must(template.ParseGlob("templates/*")),
	}
}

// ListTasks displays all tasks from the database.
func (h *Handler) ListTasks(w http.ResponseWriter, r *http.Request) {
	rows, err := h.DB.Query("SELECT id, task, completed FROM tasks")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var tasks []Task
	for rows.Next() {
		var t Task
		if err := rows.Scan(&t.ID, &t.Task, &t.Completed); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		tasks = append(tasks, t)
	}

	h.Tmpl.ExecuteTemplate(w, "index.html", tasks)
}

// AddTask handles the addition of a new task.
func (h *Handler) AddTask(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	task := r.FormValue("task")
	_, err := h.DB.Exec("INSERT INTO tasks (task) VALUES (?)", task)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/", http.StatusFound)
}

// DeleteTask handles the deletion of a task.
func (h *Handler) DeleteTask(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	_, err := h.DB.Exec("DELETE FROM tasks WHERE id = ?", id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/", http.StatusFound)
}
