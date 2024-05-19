package models

import "time"

// User representa um usuário utilizando a rede social
type User struct {
	ID        string    `json:"id,omitempty"`
	Name      string    `json:"nome,omitempty"`
	Nick      string    `json:"nick,omitempty"`
	Email     string    `json:"email,omitempty"`
	Password  string    `json:"senha,omitempty"`
	CreatedAt time.Time `json:"criado_em,omitempty"`
}
