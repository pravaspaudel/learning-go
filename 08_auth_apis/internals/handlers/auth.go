package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/pravaspaudel/10_Auth_Apis/internals/ctx"
	"github.com/pravaspaudel/10_Auth_Apis/internals/models"
	"github.com/pravaspaudel/10_Auth_Apis/internals/repositories"
	"github.com/pravaspaudel/10_Auth_Apis/internals/utils"
)

type UserHandler struct {
	Repo *repositories.UserRepositories
}

func (h *UserHandler) RegisterUser(w http.ResponseWriter, r *http.Request) {
	var user models.User

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		utils.ErrorResponse(w, http.StatusBadRequest, "invalid body")
		return
	}

	var err error

	user.Password, err = utils.HashPassword(user.Password)

	if err != nil {
		utils.ErrorResponse(w, http.StatusInternalServerError, "somethinge went wrong")
		return
	}

	createdUser, err := h.Repo.CreateUser(user)

	if err != nil {
		utils.ErrorResponse(w, http.StatusInternalServerError, "error in registering user ")
		fmt.Println("error :", err)
		return
	}

	token, err := utils.GenerateJWT(createdUser.Id)
	if err != nil {
		utils.ErrorResponse(w, http.StatusInternalServerError, "failed to generat token")
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Value:    token,
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
		Secure:   false,
	})

	utils.SuccessResponse(w, http.StatusCreated, "user created successfully", createdUser)
}

func (h *UserHandler) LoginUser(w http.ResponseWriter, r *http.Request) {
	var loginUser models.LoginUser

	if err := json.NewDecoder(r.Body).Decode(&loginUser); err != nil {
		utils.ErrorResponse(w, http.StatusBadRequest, "invalid request body")
		return
	}

	user, err := h.Repo.GetUserByEmail(loginUser.Email)

	if err != nil {
		utils.ErrorResponse(w, http.StatusUnauthorized, "invalid email or password")
		return
	}

	err = utils.ComparePassword(loginUser.Password, user.Password)

	fmt.Println(loginUser.Password, user.Password)
	fmt.Println(err)

	if err != nil {
		utils.ErrorResponse(w, http.StatusUnauthorized, "failed to compare")
		return
	}

	token, err := utils.GenerateJWT(user.Id)

	if err != nil {
		utils.ErrorResponse(w, http.StatusInternalServerError, "failed to generate token")
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Value:    token,
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
		Secure:   false,
	})

	user.Password = ""

	utils.SuccessResponse(w, http.StatusOK, "user logged in successfully", user)
}

func (h *UserHandler) CheckAuth(w http.ResponseWriter, r *http.Request) {

	userId, ok := r.Context().Value(ctx.UserKey).(string)

	if !ok {
		utils.ErrorResponse(w, http.StatusUnauthorized, "missing user")
		return
	}

	fmt.Println(userId)

	utils.SuccessResponse(w, http.StatusOK, "user authenticated", userId)
}
