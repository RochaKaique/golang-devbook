package routes

import (
	"api/src/controllers"
	"net/http"
)

var userRoutes = []Route{
	{
		URI: "/usuarios",
		Method: http.MethodPost,
		Function: controllers.CreateUser,
		AuthRequired: false,
	},
	{
		URI: "/usuarios",
		Method: http.MethodGet,
		Function: controllers.ListUsers,
		AuthRequired: false,
	},
	{
		URI: "/usuarios/{id}",
		Method: http.MethodGet,
		Function: controllers.FindUserById,
		AuthRequired: false,
	},
	{
		URI: "/usuarios/{id}",
		Method: http.MethodPut,
		Function: controllers.UpdateUser,
		AuthRequired: false,
	},
	{
		URI: "/usuarios/{id}",
		Method: http.MethodDelete,
		Function: controllers.DeleteUser,
		AuthRequired: false,
	},
}