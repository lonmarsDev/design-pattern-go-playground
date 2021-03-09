package routes

import (
	"net/http"

	"github.com/realpamisa/RestAPI/internal/api/handler"
)

type RouteSVC struct {
	Uri     string
	Method  string
	Handler func(w http.ResponseWriter, r *http.Request)
}

var urlRoutes = []RouteSVC{
	RouteSVC{
		Uri:     "/users",
		Method:  http.MethodGet,
		Handler: handler.GetUsers,
	},
	RouteSVC{
		Uri:     "/users",
		Method:  http.MethodPost,
		Handler: handler.CreateUser,
	},
	RouteSVC{
		Uri:     "/users/{id}",
		Method:  http.MethodGet,
		Handler: handler.GetUser,
	},
	RouteSVC{
		Uri:     "/users/{id}",
		Method:  http.MethodPut,
		Handler: handler.UpdateUser,
	},
	RouteSVC{
		Uri:     "/users/{id}",
		Method:  http.MethodDelete,
		Handler: handler.DeleteUser,
	},
	RouteSVC{
		Uri:     "/users/login",
		Method:  http.MethodPost,
		Handler: handler.Login,
	},
}
