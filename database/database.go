package database

import (
	"auth/model"
	"database/sql"
	"errors"

	_ "github.com/mattn/go-sqlite3"
)

const (
	sqlCreateUser = ` CREATE TABLE IF NOT EXISTS user(
        Id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
        username VARCHAR NOT NULL
    );
	`
	sqlInsertUserprofile = `
	INSERT INTO user 
		(username) VALUES ('testdata');
    `
	sqlFindUserByUsername = `
	SELECT username FROM user
		WHERE username = ?
	`
)

func InitDB(filepath string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", filepath)
	if err != nil {
		return nil, err
	}
	if db == nil {
		return nil, errors.New("no database found")
	}
	err = migrate(db)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func migrate(db *sql.DB) error {

	_, err := db.Exec(sqlCreateUser)
	if err != nil {
		return err
	}
	err = insertTest(db)
	if err != nil {
		return err
	}
	return nil
}
func insertTest(db *sql.DB) error {
	_, err := db.Exec(sqlInsertUserprofile)
	if err != nil {
		return err
	}
	return nil
}

func FindUserByUsername(username string, db *sql.DB) (*model.User, error) {
	getUser := model.User{}
	err := db.QueryRow(sqlFindUserByUsername, username).Scan(&getUser.Username)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	return &getUser, nil
}
