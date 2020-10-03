package useCases

import (
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

type IUserUseCase interface {
	Create(newUser *models.User)
	GetAll() []models.User
	Delete(id string)
	Update(id string, newUser *models.User) *models.User
}

type UserUseCase struct {
}

func NewUserUseCase() IUserUseCase {
	return &UserUseCase{}
}

func (uc *UserUseCase) Create(newUser *models.User) {
	id := guuid.New()
	newUser.Id = id.String()
	users = append(users, *newUser)
}

func (uc *UserUseCase) GetAll() []models.User {
	return users
}

func (uc *UserUseCase) Update(id string, newUser *models.User) *models.User {
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
	return &userUpdate
}

func (uc *UserUseCase) Delete(id string) {
	var newUsers = make([]models.User, 0)
	for _, usr := range users {
		if usr.Id == id {
			continue
		}
		newUsers = append(newUsers, usr)
	}
	users = newUsers
}
