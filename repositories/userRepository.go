package repositories

import (
	"database/sql"
	guuid "github.com/google/uuid"
	"gosimplemux/models"
)

var users = []models.User{
	{
		Id:        "c01d7cf6-ec3f-47f0-9556-a5d6e9009a43",
		FirstName: "Edi",
		LastName:  "Uchida",
	},
}

type UserRepository interface {
	FindOneById(id string) *models.User
	Create(newUser *models.User) error
	Update(id string, newUser *models.User) error
	Delete(id string) error
}

type userRepository struct {
	db *sql.DB
}

func (u *userRepository) FindOneById(id string) *models.User {
	var userUpdate models.User
	isFound := false
	for _, usr := range users {
		if usr.Id == id {
			userUpdate = usr
			isFound = true
			break
		}
	}
	if isFound {
		return &userUpdate
	} else {
		return nil
	}
}

func (u *userRepository) Create(newUser *models.User) error {
	id := guuid.New()
	newUser.Id = id.String()
	users = append(users, *newUser)
	return nil
}

func (u *userRepository) Update(id string, newUser *models.User) error {
	var userUpdate models.User
	var userIdx int
	for idx, usr := range users {
		if usr.Id == id {
			userUpdate = usr
			userIdx = idx
			break
		}
	}
	userUpdate.FirstName = newUser.FirstName
	userUpdate.LastName = newUser.LastName
	users[userIdx] = userUpdate
	return nil
}

func (u *userRepository) Delete(id string) error {
	var newUsers = make([]models.User, 0)
	for _, usr := range users {
		if usr.Id == id {
			continue
		}
		newUsers = append(newUsers, usr)
	}
	users = newUsers
	return nil
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{db}
}
