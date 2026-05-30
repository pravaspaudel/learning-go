package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/pravaspaudel/09_crud/internals/models"
	"github.com/pravaspaudel/09_crud/internals/repositories"
	"github.com/pravaspaudel/09_crud/internals/utils"
)

type TodoHandler struct {
	Repo *repositories.TodoRepository
}

func (h *TodoHandler) CreateTodo(w http.ResponseWriter, r *http.Request) {

	var todo models.Todos
	if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
		utils.SendErrorResponse(w, "invalid request body", http.StatusBadRequest)
		return
	}

	fmt.Println("create todo method was called")

	createdTodo, err := h.Repo.CreateTodo(todo)

	if err != nil {
		utils.SendErrorResponse(w, "failed to create todo", http.StatusInternalServerError)
		return
	}
	utils.SendSuccessResponse(w, "todo created successfully", http.StatusCreated, createdTodo)
}

// get the single todo using the id
func (h *TodoHandler) GetTodo(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	todo, err := h.Repo.GetTodo(id)

	if err != nil {
		utils.SendErrorResponse(w, "todo not found", http.StatusNotFound)
		return
	}

	utils.SendSuccessResponse(w, "single todo was called", http.StatusOK, todo)
}

func (h *TodoHandler) GetAllTodos(w http.ResponseWriter, r *http.Request) {

	todos, err := h.Repo.GetAllTodos()

	if err != nil {
		utils.SendErrorResponse(w, "failed to fetch all todos", http.StatusNotFound)
	}

	utils.SendSuccessResponse(w, "all todos fetched", http.StatusOK, todos)
}

func (h *TodoHandler) UpdateTodo(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	_, err := h.Repo.GetTodo(id)

	if err != nil {
		utils.SendErrorResponse(w, "todo does not exists", http.StatusNotFound)
		return
	}

	var todo models.Todos

	if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
		utils.SendErrorResponse(w, "invalid request body", http.StatusBadRequest)
		return
	}

	todo.ID = id

	updatedOne, err := h.Repo.UpdateTodo(todo)
	if err != nil {
		utils.SendErrorResponse(w, "failed to updated todo", http.StatusInternalServerError)
		return
	}

	utils.SendSuccessResponse(w, "updated todo", http.StatusOK, updatedOne)
}

func (h *TodoHandler) DeleteTodo(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	deletedTodo, err := h.Repo.DeleteTodo(id)
	if err != nil {
		utils.SendErrorResponse(w, "todo not found or deletion failed", http.StatusOK)
		return
	}
	utils.SendSuccessResponse(w, "todo deleted successfully", http.StatusOK, deletedTodo)
}
