package repositories

import (
	"database/sql"
	guuid "github.com/google/uuid"
	"gosimplemux/models"
)

type UserRepository interface {
	FindOneById(id string) (*models.User, error)
	Create(newUser *models.User) error
	Update(id string, newUser *models.User) error
	Delete(id string) error
}

var (
	userQueries = map[string]string{
		"insertUser":     "INSERT into sys_user(id,first_name,last_name,is_active) values(?,?,?,?)",
		"insertUserAuth": "INSERT into sys_user_auth(id,user_registration_id,user_name,user_password) values(?,?,?,?)",
		"updateUser":     "UPDATE sys_user SET first_name=?,last_name=? WHERE id=?",
		"findOneUser":    "SELECT id,first_name,last_name,is_active FROM sys_user WHERE id=?",
		"deleteUser":     "UPDATE sys_user SET is_active=? WHERE id=?",
	}
)

type userRepository struct {
	db *sql.DB
	ps map[string]*sql.Stmt
}

func (u *userRepository) FindOneById(id string) (*models.User, error) {
	row := u.ps["findOneUser"].QueryRow(id)
	res := new(models.User)
	err := row.Scan(&res.Id, &res.FirstName, &res.LastName, &res.IsActive)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (u *userRepository) Create(newUser *models.User) error {
	id := guuid.New()
	newUser.Id = id.String()
	newUserAuthId := guuid.New().String()
	tx, _ := u.db.Begin()
	_, err := tx.Stmt(u.ps["insertUser"]).Exec(newUser.Id, newUser.FirstName, newUser.LastName, "A")
	if err != nil {
		tx.Rollback()
	}
	_, err = tx.Stmt(u.ps["insertUserAuth"]).Exec(newUserAuthId, newUser.Id, newUser.FirstName+"."+newUser.LastName, "random_me")
	if err != nil {
		tx.Rollback()
	}
	err = tx.Commit()
	if err != nil {
		panic(err)
	}
	return nil
}

func (u *userRepository) Update(id string, newUser *models.User) error {
	_, err := u.ps["updateUser"].Exec(newUser.FirstName, newUser.LastName, newUser.Id)
	if err != nil {
		return err
	}
	return nil
}

func (u *userRepository) Delete(id string) error {
	_, err := u.ps["deleteUser"].Exec("N", id)
	if err != nil {
		return err
	}
	return nil
}

func NewUserRepository(db *sql.DB) UserRepository {
	ps := make(map[string]*sql.Stmt, len(userQueries))

	for n, v := range userQueries {
		p, err := db.Prepare(v)
		if err != nil {
			panic(err)
		}
		ps[n] = p
	}
	return &userRepository{db, ps}
}
