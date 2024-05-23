package repository

import (
	"api/src/models"
	"database/sql"
	"fmt"
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

	_, error = result.LastInsertId()
	if error != nil {
		return "", error
	}

	return "Usuario Inserido", nil
}

func (repo Users) Find(nameOrNick string) ([]models.User, error) {
	nameOrNick = fmt.Sprintf("%%%s%%", nameOrNick) // %nameOrNick%

	lines, error := repo.db.Query("SELECT id, nome, nick, email, criado_em FROM users WHERE nome LIKE ? OR nick LIKE ?", nameOrNick, nameOrNick)
	if error != nil {
		return nil, error
	}
	defer lines.Close()

	var users []models.User
	for lines.Next() {
		var user models.User

		if error = lines.Scan(
			&user.ID,
			&user.Name,
			&user.Nick,
			&user.Email,
			&user.CreatedAt,
		); error != nil {
			return nil, error
		}
		users = append(users, user)
	}

	return users, nil
}

func (repo Users) FindById(id string) (models.User, error) {
	line, error := repo.db.Query("SELECT id, nome, nick, email, criado_em FROM users WHERE id = ?", id)
	if error != nil {
		return models.User{}, error
	}
	defer line.Close()

	var user models.User
	if line.Next() {
		if error = line.Scan(
			&user.ID,
			&user.Name,
			&user.Nick,
			&user.Email,
			&user.CreatedAt,
		); error != nil {
			return models.User{}, error
		}
	}
	return user, nil
}

func (repo Users) Update(id string, user models.User) error {
	statement, error := repo.db.Prepare("UPDATE users SET nome = ?, nick = ?, email = ? WHERE id = ?")
	if error != nil {
		return error
	}
	defer statement.Close()

	if _, error = statement.Exec(user.Name, user.Nick, user.Email, id); error != nil {
		return error
	}

	return nil
}

func (repo Users) Delete(id string) error {
	statement, error := repo.db.Prepare("DELETE from users WHERE id = ?")
	if error != nil {
		return error
	}
	defer statement.Close()

	if _, error = statement.Exec(id); error != nil {
		return error
	}

	return nil
}

func (repo Users) FindByEmailForLogin(email string) (models.User, error) {
	line, error := repo.db.Query("SELECT id, senha FROM users WHERE email = ?", email)
	if error != nil {
		return models.User{}, error
	}
	defer line.Close()

	var user models.User
	if line.Next() {
		if error = line.Scan(
			&user.ID,
			&user.Password,
		); error != nil {
			return models.User{}, error
		}
	}

	return user, nil
}
