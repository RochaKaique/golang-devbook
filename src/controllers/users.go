package controllers

import (
	"api/src/authentication"
	database "api/src/db"
	"api/src/models"
	"api/src/repository"
	"api/src/responses"
	"api/src/security"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

// CreateUser insere um usuário no banco de dados
func CreateUser(w http.ResponseWriter, r *http.Request) {

	requestBody, error := io.ReadAll(r.Body)
	if error != nil {
		responses.Error(w, http.StatusUnprocessableEntity, error)
		return
	}

	var user models.User
	if error = json.Unmarshal(requestBody, &user); error != nil {
		responses.Error(w, http.StatusBadRequest, error)
		return
	}

	if error = user.Prepare("cadastro"); error != nil {
		responses.Error(w, http.StatusBadRequest, error)
		return
	}

	db, error := database.Connect()
	if error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}
	defer db.Close()

	repository := repository.NewUsersRepository(db)
	_, error = repository.Create(user)
	if error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}

	responses.JSON(w, http.StatusCreated, user)

}

// ListUsers lista os usuário cadastrados
func ListUsers(w http.ResponseWriter, r *http.Request) {
	nameOrNick := strings.ToLower(r.URL.Query().Get("usuario"))

	db, error := database.Connect()
	if error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}
	defer db.Close()

	repository := repository.NewUsersRepository(db)

	users, error := repository.Find(nameOrNick)
	if error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}

	responses.JSON(w, http.StatusOK, users)
}

// FindUserById busca um usuário específico por id
func FindUserById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	ID := params["id"]

	db, error := database.Connect()
	if error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}
	defer db.Close()

	repository := repository.NewUsersRepository(db)
	user, error := repository.FindById(ID)
	if error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
	}

	responses.JSON(w, http.StatusOK, user)
}

// UpdateUser atualiza os dados de um usuário no banco de dados
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	ID := params["id"]

	requestBody, error := io.ReadAll(r.Body)
	if error != nil {
		responses.Error(w, http.StatusUnprocessableEntity, error)
		return
	}

	var user models.User
	if error = json.Unmarshal(requestBody, &user); error != nil {
		responses.Error(w, http.StatusBadRequest, error)
		return
	}

	if error = user.Prepare("atualizacao"); error != nil {
		responses.Error(w, http.StatusBadRequest, error)
		return
	}

	userIDFromToken, error := authentication.ExtractUserID(r)
	if error != nil {
		responses.Error(w, http.StatusUnauthorized, error)
		return
	}

	if ID != userIDFromToken {
		responses.Error(w, http.StatusForbidden, errors.New("Não é possivel atualizar um usuário que não seja o seu"))
		return
	}

	db, error := database.Connect()
	if error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}
	defer db.Close()

	repository := repository.NewUsersRepository(db)
	if error = repository.Update(ID, user); error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)
}

// DeleteUser apaga um usuário do banco de dados
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	ID := params["id"]

	userIDFromToken, error := authentication.ExtractUserID(r)
	if error != nil {
		responses.Error(w, http.StatusUnauthorized, error)
		return
	}

	if ID != userIDFromToken {
		responses.Error(w, http.StatusForbidden, errors.New("Não é possivel deletar um usuário que não seja o seu"))
		return
	}

	db, error := database.Connect()
	if error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}
	defer db.Close()

	repository := repository.NewUsersRepository(db)

	if error = repository.Delete(ID); error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)
}

// FollowUser permie que um usuário siga outro
func FollowUser (w http.ResponseWriter, r *http.Request) {
	followerID, error := authentication.ExtractUserID(r)
	if error != nil {
		responses.Error(w, http.StatusUnauthorized, error)
		return
	}

	params := mux.Vars(r)
	userID := params["userId"]

	if userID == followerID {
		responses.Error(w, http.StatusForbidden, errors.New("Não é possível serguir voce mesmo"))
		return
	}

	db, error := database.Connect()
	if error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}
	defer db.Close()

	repository := repository.NewUsersRepository(db)
	if error = repository.Follow(userID, followerID); error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)
}

// UnfollowUser permite que um usuário de seguir outro usuário
func UnfollowUser (w http.ResponseWriter, r *http.Request) {
	followerID, error := authentication.ExtractUserID(r)
	if error != nil {
		responses.Error(w, http.StatusUnauthorized, error)
		return
	}

	params := mux.Vars(r)
	userID := params["userId"]

	if userID == followerID {
		responses.Error(w, http.StatusForbidden, errors.New("Não é possivel parar de segur você mesmo"))
		return
	}

	db, error := database.Connect()
	if error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}
	defer db.Close()

	repository := repository.NewUsersRepository(db)
	if error = repository.Unfollow(userID, followerID); error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)
}

// FindFollowers tras todos os seguidores de um usuário
func FindFollowers(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userId := params["userID"]

	db, error := database.Connect()
	if error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}
	defer db.Close()

	repository := repository.NewUsersRepository(db)
	followers, error := repository.FindFollowers(userId)

	if error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}

	responses.JSON(w, http.StatusOK, followers)
}

// FindFollowing trás todos as pessoas que um usuário segue
func FindFollowig(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userId := params["userID"]

	db, error := database.Connect()
	if error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}
	defer db.Close()

	repository := repository.NewUsersRepository(db)
	following, error := repository.FindFollowing(userId)

	if error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}

	responses.JSON(w, http.StatusOK, following)
}

// UpdatePassword permite o usuário atualizar sua senha
func UpdatePassword(w http.ResponseWriter, r *http.Request) {
	loggedUserId, error := authentication.ExtractUserID(r)
	if error != nil {
		responses.Error(w, http.StatusUnauthorized, error)
		return
	}

	params := mux.Vars(r)
	paramUserId := params["userID"]

	if loggedUserId != paramUserId {
		responses.Error(w, http.StatusForbidden, errors.New("Não é possivel atualizar a senha de outro usuário"))
		return
	}

	body, error := io.ReadAll(r.Body)

	var password models.Password
	if error = json.Unmarshal(body, &password); error != nil {
		responses.Error(w, http.StatusBadRequest, error)
		return
	}

	db, error := database.Connect()
	if error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}
	defer db.Close()

	repository := repository.NewUsersRepository(db)
	passwordInDb, error := repository.FindPassword(loggedUserId)
	if error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}

	if error = security.CheckPassword(passwordInDb, password.ActualPassword); error != nil {
		responses.Error(w, http.StatusUnauthorized, errors.New("A senha atual não condiz com a senha que está salva no banco"))
		return
	}

	hashedPassword, error := security.Hash(password.NewPassword)
	if error != nil {
		responses.Error(w, http.StatusBadRequest, error)
		return
	}

	if error = repository.UpdatePassword(loggedUserId, string(hashedPassword)); error != nil {
		responses.Error(w, http.StatusInternalServerError, error)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)
}

