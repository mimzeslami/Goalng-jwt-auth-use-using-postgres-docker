package models

import "testing"

func TestHashPassword(t *testing.T) {
	user := User{
		Password: "password",
	}
	err := user.HashPassword(user.Password)
	if err != nil {
		t.Error("Error hashing password")
	}
	if user.Password == "password" {
		t.Error("Password not hashed")
	}

}

func TestCheckPassword(t *testing.T) {
	user := User{
		Password: "password",
	}
	err := user.HashPassword(user.Password)
	if err != nil {
		t.Error("Error hashing password")
	}
	err = user.CheckPassword("password")
	if err != nil {
		t.Error("Error checking password")
	}
	err = user.CheckPassword("password1")
	if err == nil {
		t.Error("Password not checked")
	}

}
