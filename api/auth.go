package api

import (
	"encoding/json"
	"fmt"
	"io"
	m "my-social-network/models"
	"net/http"
)

func (api *API) RegisterNewUser(w http.ResponseWriter, r *http.Request) {

	content, err := io.ReadAll(r.Body)
	if err != nil {

		fmt.Println("error reading body")
		return
	}

	var body m.RegisterStruct
	if err := json.Unmarshal(content, &body); err != nil {
		fmt.Println("Error validating body")
		return
	}

	fmt.Println("Working until now")
	return
}

func (api *API) LoginUser(w http.ResponseWriter, r *http.Request) {
	content, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println("error reading body")
		return
	}

	var body m.LoginStruct
	if err = json.Unmarshal(content, &body); err != nil {
		fmt.Println("error validating body")
		return
	}

	fmt.Println("Working until now")
	return
}

func (api *API) RefreshToken(w http.ResponseWriter, r *http.Request) {

	fmt.Println("METHOD NOT IMPLEMENTED!")
}
