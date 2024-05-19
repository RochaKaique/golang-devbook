package repository

import (
	"api/src/models"
	"database/sql"
)

type Users struct {
	db *sql.DB
}

func NewUsersRepository(db *sql.DB) *Users {
	return &Users{db}
}

func (repo Users) Create(user models.User) (string, error) {
	statement, error := repo.db.Prepare("insert into users (nome, nick, email, senha) values (?, ?, ?, ?)")
	if error != nil {
		return "", error
	}
	defer statement.Close()

	result, error := statement.Exec(user.Name, user.Nick, user.Email, user.Password)
	if error != nil {
		return "", error
	}

	_, error = result.LastInsertId();
	if error != nil {
		return "", error
	}

	return "Usuario Inserido", nil
}
