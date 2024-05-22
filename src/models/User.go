package models

import (
	"api/src/security"
	"errors"
	"strings"
	"time"

	"github.com/badoux/checkmail"
)

// User representa um usuário utilizando a rede social
type User struct {
	ID        string    `json:"id,omitempty"`
	Name      string    `json:"nome,omitempty"`
	Nick      string    `json:"nick,omitempty"`
	Email     string    `json:"email,omitempty"`
	Password  string    `json:"senha,omitempty"`
	CreatedAt time.Time `json:"criado_em,omitempty"`
}

func (user *User) Prepare(step string) error {
	if error := user.validate(step); error != nil {
		return error
	}

	if error := user.format(step); error != nil {
		return error
	}
	return nil
}

func (user *User) validate(step string) error {
	if user.Name == "" {
		return errors.New("O nome é obrigatorio e não pode estar em branco")
	}

	if user.Nick == "" {
		return errors.New("O nick é obrigatorio e não pode estar em branco")
	}

	if user.Email == "" {
		return errors.New("O e-mail é obrigatorio e não pode estar em branco")
	}

	if error := checkmail.ValidateFormat(user.Email); error != nil {
		return errors.New("Formato de e-mail inválido")
	}

	if step == "cadastro" && user.Password == "" {
		return errors.New("A senha é obrigatorio e não pode estar em branco")
	}

	return nil
}

func (user *User) format(step string) error {
	user.Name = strings.TrimSpace(user.Name)
	user.Email = strings.TrimSpace(user.Email)
	user.Nick = strings.TrimSpace(user.Nick)

	if step == "cadastro" {
		passwordHash, error := security.Hash(user.Password)
		if error != nil {
			return errors.New("Erro ao criptografar senha")
		}

		user.Password = string(passwordHash)
	}

	return nil
}
