package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/pravaspaudel/10_Auth_Apis/config"
	"github.com/pravaspaudel/10_Auth_Apis/internals/handlers"
	"github.com/pravaspaudel/10_Auth_Apis/internals/repositories"
	"github.com/pravaspaudel/10_Auth_Apis/internals/routes"
)

func main() {

	config.LoadEnv()
	Db := config.ConfigDB()

	userRepo := &repositories.UserRepositories{
		Db: Db,
	}

	userHandler := &handlers.UserHandler{
		Repo: userRepo,
	}

	r := chi.NewRouter()

	routes.UserRouters(r, userHandler)

	server := &http.Server{
		Addr:    ":" + config.AppConfig.PORT,
		Handler: r,
	}

	fmt.Printf("server is running on http://localhost%v\n", server.Addr)

	if err := server.ListenAndServe(); err != nil {
		fmt.Println("error while starting the server")
	}

}
