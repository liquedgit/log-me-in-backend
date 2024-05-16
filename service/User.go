package service

import (
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"log-me-in/database"
	"log-me-in/model"
)

func GetUserByUsername(username string) (*model.User, error) {
	db := database.GetConnection()
	var user model.User
	result := db.Where("username = ?", username).First(&user)
	if result.Error != nil {
		return nil, result.Error
	} else {
		return &user, nil
	}
}

func CreateNewUser(username string, password string) (*model.User, error) {
	db := database.GetConnection()

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		panic(err)
	}

	user := model.User{
		Id:             uuid.NewString(),
		Username:       username,
		HashedPassword: string(hashedPassword),
		Role:           "User",
	}
	result := db.Create(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}
