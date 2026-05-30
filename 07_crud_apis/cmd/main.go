package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/pravaspaudel/09_crud/config"
	"github.com/pravaspaudel/09_crud/internals/db"
	"github.com/pravaspaudel/09_crud/internals/handlers"
	"github.com/pravaspaudel/09_crud/internals/repositories"
)

func main() {
	config.ConfigureEnv()

	r := chi.NewRouter()

	PORT := config.App.PORT

	server := &http.Server{
		Addr:    fmt.Sprintf(":%v", PORT),
		Handler: r,
	}

	db.ConnectDb()

	reposit := &repositories.TodoRepository{
		DB: db.Db,
	}

	tHandler := &handlers.TodoHandler{
		Repo: reposit,
	}

	r.Post("/todos", tHandler.CreateTodo)
	r.Get("/todos", tHandler.GetAllTodos)
	r.Get("/todos/{id}", tHandler.GetTodo)
	r.Patch("/todos/{id}", tHandler.UpdateTodo)
	r.Delete("/todos/{id}", tHandler.DeleteTodo)

	fmt.Printf("server is running on http://localhost%v\n", server.Addr)
	err := server.ListenAndServe()

	if err != nil {
		log.Fatal("Error starting the server :", err)
	}
}
