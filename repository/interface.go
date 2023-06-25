package repository

import "training/models"

type DbInterface interface {
	Login(loginPayload models.User) (models.UserResponse, error)
	Register(registerPayload models.User) (models.UserResponse, error)
}
