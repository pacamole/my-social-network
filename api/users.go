package api

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func (api *API) GetAllUsers(w http.ResponseWriter, r *http.Request) {

	fmt.Println("METHOD NOT IMPLEMENTED!")
	return
}

func (api *API) GetUserByID(w http.ResponseWriter, r *http.Request) {
	userID := mux.Vars(r)["user_id"]

	fmt.Println("UserID:", userID)

	fmt.Println("Working until now")
	return
}
