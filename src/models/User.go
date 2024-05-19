package models

import (
	"errors"
	"strings"
	"time"
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

func (user *User) Prepare() error {
	if error := user.validate(); error != nil {
		return error
	}

	user.format()
	return nil
}

func (user *User) validate() error {
	if user.Name == "" {
		return errors.New("O nome é obrigatorio e não pode estar em branco")
	}

	if user.Nick == "" {
		return errors.New("O nick é obrigatorio e não pode estar em branco")
	}

	if user.Email == "" {
		return errors.New("O e-mail é obrigatorio e não pode estar em branco")
	}

	if user.Password == "" {
		return errors.New("A senha é obrigatorio e não pode estar em branco")
	}

	return nil
}

func (user *User) format() {
	user.Name = strings.TrimSpace(user.Name)
	user.Email = strings.TrimSpace(user.Email)
	user.Nick = strings.TrimSpace(user.Nick)
}
