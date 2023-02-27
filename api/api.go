package api

import (
	"log"
	"net/http"
)

type API struct {
}

func StartAPI() *API {
	api := &API{}

	router := api.InicializeRouter()

	log.Fatal(http.ListenAndServe(":5000", router))

	return api
}
