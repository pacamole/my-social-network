package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func (api *API) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := api.users.GetAllUsers()
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	response, err := json.Marshal(users)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	w.Write(response)
}

func (api *API) GetUserByID(w http.ResponseWriter, r *http.Request) {
	userIDs := mux.Vars(r)["user_id"]
	fmt.Println("UserID:", userIDs)

	userID, err := strconv.ParseUint(userIDs, 10, 16)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	user, err := api.users.GetUserByID(uint16(userID))
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	response, err := json.Marshal(user)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	w.Write(response)
}
