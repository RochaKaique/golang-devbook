package routes

import (
	"api/src/controllers"
	"net/http"
)

var userRoutes = []Route{
	{
		URI:          "/usuarios",
		Method:       http.MethodPost,
		Function:     controllers.CreateUser,
		AuthRequired: false,
	},
	{
		URI:          "/usuarios",
		Method:       http.MethodGet,
		Function:     controllers.ListUsers,
		AuthRequired: true,
	},
	{
		URI:          "/usuarios/{id}",
		Method:       http.MethodGet,
		Function:     controllers.FindUserById,
		AuthRequired: true,
	},
	{
		URI:          "/usuarios/{id}",
		Method:       http.MethodPut,
		Function:     controllers.UpdateUser,
		AuthRequired: true,
	},
	{
		URI:          "/usuarios/{id}",
		Method:       http.MethodDelete,
		Function:     controllers.DeleteUser,
		AuthRequired: true,
	},
	{
		URI:          "/usuarios/{userId}/seguir",
		Method:       http.MethodPost,
		Function:     controllers.FollowUser,
		AuthRequired: true,
	},
	{
		URI:          "/usuarios/{userId}/parar-de-seguir",
		Method:       http.MethodPost,
		Function:     controllers.UnfollowUser,
		AuthRequired: true,
	},
	{
		URI: "/usuarios/{userID}/seguidores",
		Method: http.MethodGet,
		Function: controllers.FindFollowers,
		AuthRequired: true,
	},
	{
		URI: "/usuarios/{userID}/seguindo",
		Method: http.MethodGet,
		Function: controllers.FindFollowig,
		AuthRequired: true,
	},
	{
		URI: "/usuarios/{userID}/atualizar-senha",
		Method: http.MethodPost,
		Function: controllers.UpdatePassword,
		AuthRequired: true,
	},
}
