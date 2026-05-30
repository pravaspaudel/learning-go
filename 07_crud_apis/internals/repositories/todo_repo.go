package repositories

import (
	"database/sql"
	"log"

	"github.com/pravaspaudel/09_crud/internals/models"
)

type TodoRepository struct {
	DB *sql.DB
}

func (r *TodoRepository) CreateTodo(todo models.Todos) (models.Todos, error) {
	log.Println("create todo method was called")
	query := `INSERT INTO todos(title,done) VALUES ($1,$2) RETURNING id,title,done`

	var created models.Todos

	err := r.DB.QueryRow(query, todo.Title, todo.Done).Scan(&created.ID, &created.Title, &created.Done)

	if err != nil {
		return models.Todos{}, err
	}

	return created, nil
}

func (r *TodoRepository) GetTodo(id string) (models.Todos, error) {
	var todo models.Todos
	query := `SELECT id,title,done FROM todos WHERE id=$1`

	err := r.DB.QueryRow(query, id).Scan(&todo.ID, &todo.Title, &todo.Done)

	return todo, err
}

func (r *TodoRepository) GetAllTodos() ([]models.Todos, error) {

	query := `SELECT id,title,done FROM todos`

	rows, err := r.DB.Query(query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var todos []models.Todos

	for rows.Next() {
		var t models.Todos
		err := rows.Scan(&t.ID, &t.Title, &t.Done)

		if err != nil {
			return nil, err
		}

		todos = append(todos, t)
	}

	return todos, nil
}

func (r *TodoRepository) UpdateTodo(todo models.Todos) (models.Todos, error) {
	query := `UPDATE todos SET title=$1, done=$2 WHERE id=$3 RETURNING id,title,done`

	var updatedTodo models.Todos

	err := r.DB.QueryRow(query, todo.Title, todo.Done, todo.ID).Scan(&updatedTodo.ID, &updatedTodo.Title, &updatedTodo.Done)

	if err != nil {
		if err == sql.ErrNoRows {
			return models.Todos{}, sql.ErrNoRows
		}
		return models.Todos{}, err
	}
	return updatedTodo, nil
}

func (r *TodoRepository) DeleteTodo(id string) (models.Todos, error) {
	query := `DELETE FROM todos WHERE id=$1 RETURNING id,title,done`

	var deletedTodo models.Todos
	err := r.DB.QueryRow(query, id).Scan(&deletedTodo.ID, &deletedTodo.Title, &deletedTodo.Done)

	if err != nil {
		return models.Todos{}, err
	}

	return deletedTodo, nil
}
