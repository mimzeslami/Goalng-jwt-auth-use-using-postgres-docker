package controllers

import (
	"encoding/json"
	"net/http"
	"training/models"
	"training/repository"
	"training/utils"
)

type UserController struct {
	repo repository.DbInterface
}

func NewUserController(repo repository.DbInterface) *UserController {
	return &UserController{repo}
}

func (u *UserController) Login(w http.ResponseWriter, r *http.Request) {
	var loginPayload models.User
	json.NewDecoder(r.Body).Decode(&loginPayload)

	user, err := u.repo.Login(loginPayload)
	if err != nil {
		utils.ErrorJSON(w, err, 400)
		return
	}

	token, err := utils.GenerateJWT(user.ID)
	if err != nil {
		utils.ErrorJSON(w, err, 400)
		return
	}

	user.Token = token

	utils.SuccessJSON(w, "Successfull login", user)
}

func (u *UserController) Register(w http.ResponseWriter, r *http.Request) {
	var registerPayload models.User
	json.NewDecoder(r.Body).Decode(&registerPayload)

	user, err := u.repo.Register(registerPayload)
	if err != nil {
		utils.ErrorJSON(w, err, 500)
		return
	}
	token, err := utils.GenerateJWT(user.ID)
	if err != nil {
		utils.ErrorJSON(w, err, 500)
		return
	}

	user.Token = token

	utils.SuccessJSON(w, "Successfull register", user)
}
