package repositories

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

func DbPrep() *sql.DB {
	database, _ := sql.Open("sqlite3", ":memory:")

	database.Exec("CREATE TABLE IF NOT EXISTS sys_user(id TEXT PRIMARY KEY,first_name TEXT,last_name TEXT,is_active TEXT)")
	database.Exec("CREATE TABLE IF NOT EXISTS sys_user_auth(id TEXT PRIMARY KEY,user_registration_id TEXT,user_name TEXT,user_password TEXT)")

	database.Exec("INSERT INTO sys_user VALUES ('1','Dummy First Name 1','Dummy Last Name 1','A')")
	database.Exec("INSERT INTO sys_user_auth VALUES ('1','1','user.name','secret')")

	return database
}
