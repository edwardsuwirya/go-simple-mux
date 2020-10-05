package repositories

import (
	"database/sql"
	guuid "github.com/google/uuid"
	"gosimplemux/models"
)

var (
	authQueries = map[string]string{
		"insertUserAuth":               "INSERT into user_auth(id,user_registration_id,user_name,user_password) values(?,?,?,?)",
		"updateUserPassword":           "UPDATE user_auth SET user_password=? WHERE id=?",
		"findOneByUserNameAndPassword": "SELECT id,user_registration_id FROM user_auth WHERE user_name=? AND user_password=?",
	}
)

type UserAuthRepository interface {
	FindOneByUserNameAndPassword(userName string, password string) *models.UserAuth
	Create(newUser *models.UserAuth) error
	UpdatePassword(id string, newPassword string) error
}

type userAuthRepository struct {
	db *sql.DB
	ps map[string]*sql.Stmt
}

func (u *userAuthRepository) FindOneByUserNameAndPassword(userName string, password string) *models.UserAuth {
	row := u.ps["findOneByUserNameAndPassword"].QueryRow(userName, password)
	res := new(models.UserAuth)
	err := row.Scan(&res.Id, &res.UserRegId)
	if err != nil {
		return nil
	}
	return res
}

func (u *userAuthRepository) Create(newUser *models.UserAuth) error {
	id := guuid.New()
	newUser.Id = id.String()
	_, err := u.ps["insertUserAuth"].Exec(newUser.Id, newUser.UserRegId, newUser.UserName, newUser.UserPassword)
	if err != nil {
		return err
	}
	return nil
}

func (u *userAuthRepository) UpdatePassword(id string, newPassword string) error {
	_, err := u.ps["updateUserPassword"].Exec(newPassword, id)
	if err != nil {
		return err
	}
	return nil
}

func NewUserAuthRepository(db *sql.DB) UserAuthRepository {
	ps := make(map[string]*sql.Stmt, len(authQueries))

	for n, v := range authQueries {
		p, err := db.Prepare(v)
		if err != nil {
			panic(err)
		}
		ps[n] = p
	}

	return &userAuthRepository{
		db, ps,
	}
}
