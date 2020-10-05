package repositories

import (
	"database/sql"
	"gosimplemux/models"
)

var (
	authQueries = map[string]string{
		"updateUserPassword":                   "UPDATE sys_user_auth SET user_password=? WHERE id=?",
		"findOneUserAuthByUserNameAndPassword": "SELECT id,user_registration_id FROM sys_user_auth WHERE user_name=? AND user_password=?",
	}
)

type UserAuthRepository interface {
	FindOneByUserNameAndPassword(userName string, password string) *models.UserAuth
	UpdatePassword(id string, newPassword string) error
}

type userAuthRepository struct {
	db *sql.DB
	ps map[string]*sql.Stmt
}

func (u *userAuthRepository) FindOneByUserNameAndPassword(userName string, password string) *models.UserAuth {
	row := u.ps["findOneUserAuthByUserNameAndPassword"].QueryRow(userName, password)
	res := new(models.UserAuth)
	err := row.Scan(&res.Id, &res.UserRegId)
	if err != nil {
		return nil
	}
	return res
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
