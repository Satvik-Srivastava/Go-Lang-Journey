package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// ------------------ Models ------------------

type Todo struct {
	ID    string `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

var todos = []Todo{} // In-memory DB

// ------------------ Handlers ------------------

// Home page

func HomePage(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`<h1>This is the home page</h1>`))
}

// Create a new Todo
func CreateTodo(w http.ResponseWriter, r *http.Request) {
	var todo Todo
	_ = json.NewDecoder(r.Body).Decode(&todo)

	todos = append(todos, todo)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todo)
}

// Get all Todos
func GetTodos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todos)
}

// Get single Todo by ID
func GetTodoByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	for _, item := range todos {
		if item.ID == id {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(item)
			return
		}
	}

	http.NotFound(w, r)
}

// Update Todo by ID (PUT)
func UpdateTodo(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	for i, item := range todos {
		if item.ID == id {
			var updated Todo
			_ = json.NewDecoder(r.Body).Decode(&updated)

			todos[i] = updated

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(updated)
			return
		}
	}

	http.NotFound(w, r)
}

// Delete Todo by ID
func DeleteTodo(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	for i, item := range todos {
		if item.ID == id {
			todos = append(todos[:i], todos[i+1:]...)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}

	http.NotFound(w, r)
}

// ------------------ Main ------------------

func main() {
	r := mux.NewRouter()

	// CRUD ROUTES
	r.HandleFunc("/", HomePage).Methods("GET")
	r.HandleFunc("/todos", CreateTodo).Methods("POST")
	r.HandleFunc("/todos", GetTodos).Methods("GET")
	r.HandleFunc("/todos/{id}", GetTodoByID).Methods("GET")
	r.HandleFunc("/todos/{id}", UpdateTodo).Methods("PUT")
	r.HandleFunc("/todos/{id}", DeleteTodo).Methods("DELETE")

	log.Println("Server started at :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
