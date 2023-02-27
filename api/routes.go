package api

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	URI         string
	Method      string
	Func        func(http.ResponseWriter, *http.Request)
	RequireAuth bool
}

func (api *API) InicializeRouter() *mux.Router {
	routes := []Route{
		// Authorization
		{
			URI: "/auth/login", Method: "POST",
			Func: api.LoginUser, RequireAuth: false,
		},
		{
			URI: "/auth/register", Method: "POST",
			Func: api.RegisterNewUser, RequireAuth: false,
		},
		{
			URI: "/auth/refresh", Method: "POST",
			Func: api.RefreshToken, RequireAuth: false,
		},
		// -----------
		// Users
		{
			URI: "/users", Method: "GET",
			Func: api.GetAllUsers, RequireAuth: false,
		},
		{
			URI: "/users/{user_id}", Method: "GET",
			Func: api.GetUserByID, RequireAuth: false,
		},
		// -----------
	}

	router := mux.NewRouter()

	for _, route := range routes {
		router.HandleFunc(route.URI, route.Func).Methods(route.Method)
	}

	return router
}
