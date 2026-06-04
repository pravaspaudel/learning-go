package routes

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/pravaspaudel/10_Auth_Apis/internals/handlers"
	"github.com/pravaspaudel/10_Auth_Apis/internals/middlewares"
)

func UserRouters(r chi.Router, userHandler *handlers.UserHandler) {
	r.Post("/auth/register", userHandler.RegisterUser)
	r.Post("/auth/login", userHandler.LoginUser)
	// r.Get("/auth/check",userHandler.)

	r.Get("/auth/check",
		middlewares.ProtectMiddleware(
			http.HandlerFunc(userHandler.CheckAuth),
		),
	)
}
