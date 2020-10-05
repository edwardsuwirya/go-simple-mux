package models

type UserAuth struct {
	Id           string `json:"id"`
	UserRegId    string `json:"userRegistrationId"`
	UserName     string `json:"userName"`
	UserPassword string `json:"userPassword"`
}
