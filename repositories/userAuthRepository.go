package repositories

import (
	guuid "github.com/google/uuid"
	"gosimplemux/models"
)

var userAuth = []models.UserAuth{
	{
		Id:           "f78a0730-5e27-4faf-8ccc-21edd53306dd",
		UserRegId:    "c01d7cf6-ec3f-47f0-9556-a5d6e9009a43",
		UserName:     "ediboy93",
		UserPassword: "12345",
	},
}

type UserAuthRepository interface {
	FindOneByUserNameAndPassword(userName string, password string) *models.UserAuth
	Create(newUser *models.UserAuth) error
	UpdatePassword(id string, newPassword string) error
}

type userAuthRepository struct {
}

func (u *userAuthRepository) FindOneByUserNameAndPassword(userName string, password string) *models.UserAuth {
	var userUpdate models.UserAuth
	isFound := false
	for _, usr := range userAuth {
		if usr.UserName == userName && usr.UserPassword == password {
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

func (u *userAuthRepository) Create(newUser *models.UserAuth) error {
	id := guuid.New()
	newUser.Id = id.String()
	userAuth = append(userAuth, *newUser)
	return nil
}

func (u *userAuthRepository) UpdatePassword(id string, newPassword string) error {
	var userUpdate models.UserAuth
	var userIdx int
	for idx, usr := range userAuth {
		if usr.Id == id {
			userUpdate = usr
			userIdx = idx
			break
		}
	}
	userUpdate.UserPassword = newPassword
	userAuth[userIdx] = userUpdate
	return nil
}

func NewUserAuthRepository() UserAuthRepository {
	return &userAuthRepository{}
}
