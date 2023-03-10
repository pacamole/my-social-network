package api

import (
	"fmt"
	"log"
	"my-social-network/config"
	repos "my-social-network/repositories"
	srv "my-social-network/services"
	"net/http"
)

type API struct {
	users srv.IUsers
	auth  srv.IAuth
}

func StartAPI() {
	if err := config.EnvironmentConfig(); err != nil {
		log.Panicf("ERROR on configuring environment: %s", err)
	}
	fmt.Println("1/4 - Enviroment set up!")

	db, err := repos.ConnectGorm()
	if err != nil {
		log.Panicf("ERROR initializing GORM Client: %v", err)
	}

	if err = repos.MigrateGormDB(db); err != nil {
		log.Panicf("ERROR on migrating DB: %v", err)
	}
	fmt.Println("2/4 - Migration to DB completed!")

	client := repos.NewGormClient(db)
	fmt.Println("3/4 - GORM connected to DB!")

	userConn, err := srv.NewUserConn(client)
	if err != nil {
		log.Panicf("ERROR connecting user client: %v", err)
	}

	api := &API{userConn, nil}

	router := api.InicializeRouter()
	fmt.Println("4/4 - Routes are set up!")

	log.Fatal(http.ListenAndServe(":"+config.ApiPort, router))
}
