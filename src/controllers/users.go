package controllers

import "net/http"

// CreateUser insere um usuário no banco de dados
func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Criando Usuário"))
}

// ListUsers lista os usuário cadastrados
func ListUsers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Listando Usuário"))
}

// FindUserById busca um usuário específico por id
func FindUserById(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Encontrando Usuário"))
}

// UpdateUser atualiza os dados de um usuário no banco de dados
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Atualizando Usuário"))
}

// DeleteUser apaga um usuário do banco de dados
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Deletando Usuário"))
}