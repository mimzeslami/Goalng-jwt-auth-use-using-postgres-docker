package repository

import (
	"errors"
	"log"
	"training/models"
)

func (db PostgresImpl) Login(loginPayload models.User) (models.UserResponse, error) {
	var user models.User

	stme, err := db.db.Prepare(`SELECT * FROM "user" WHERE email = $1`)
	if err != nil {
		log.Println(err)
		return models.UserResponse{}, err
	}
	defer stme.Close()
	_, err = stme.QueryOne(&user, loginPayload.Email)
	if err != nil {
		log.Println(err)
		return models.UserResponse{}, errors.New("Invalid email or password")
	}

	err = user.CheckPassword(loginPayload.Password)
	if err != nil {
		log.Println(err)
		return models.UserResponse{}, errors.New("Invalid email or password")
	}

	return models.UserResponse{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
	}, nil

}

func (db PostgresImpl) Register(registerPayload models.User) (models.UserResponse, error) {
	var user models.User

	err := user.HashPassword(registerPayload.Password)
	if err != nil {
		log.Println(err)
		return models.UserResponse{}, err
	}

	stme, err := db.db.Prepare(`INSERT INTO "user" (first_name, last_name, email, password) VALUES ($1, $2, $3, $4) RETURNING *`)
	if err != nil {
		log.Println(err)
		return models.UserResponse{}, err
	}
	defer stme.Close()
	_, err = stme.QueryOne(&user, registerPayload.FirstName, registerPayload.LastName, registerPayload.Email, user.Password)
	if err != nil {

		log.Println(err)
		return models.UserResponse{}, err
	}

	return models.UserResponse{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
	}, nil

}
