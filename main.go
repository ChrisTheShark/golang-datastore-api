package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"cloud.google.com/go/datastore"

	"github.com/ChrisTheShark/golang-datastore-api/controllers"
	"github.com/ChrisTheShark/golang-datastore-api/repository"
	"github.com/julienschmidt/httprouter"
)

func main() {
	r := httprouter.New()

	ur := repository.NewUserRepository(getClient())
	uc := controllers.NewUserController(ur)

	r.GET("/users", uc.GetUsers)
	r.POST("/users", uc.AddUser)
	r.GET("/users/:id", uc.GetUserByID)
	r.DELETE("/users/:id", uc.DeleteUser)
	log.Fatal(http.ListenAndServe(":8080", r))
}

func getClient() *datastore.Client {
	// Passing an empty string as project identifier looks up the
	// project iddentifier from the same ENV value specified here. I
	// feel making this explicit is better practice.
	client, err := datastore.NewClient(
		context.Background(), os.Getenv("DATASTORE_PROJECT_ID"))
	if err != nil {
		log.Panic(err)
	}
	return client
}
