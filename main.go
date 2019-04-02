package main

import (
	"database/sql"
	"fmt"
	"go-webapi-poc/controllers"
	"go-webapi-poc/driver"
	"go-webapi-poc/models"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/subosito/gotenv"
)

var authusers []models.AuthUser

var db *sql.DB

func init() {
	gotenv.Load()
}

func logFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	db = driver.ConnectDB()
	router := mux.NewRouter()
	controller := controllers.Controller{}

	router.HandleFunc("/user/all", controller.GetUsers(db)).Methods("GET")
	router.HandleFunc("/user/{id}", controller.GetUser(db)).Methods("GET")
	router.HandleFunc("/user", controller.AddUser(db)).Methods("POST")
	router.HandleFunc("/user", controller.UpdateUser(db)).Methods("PUT")
	router.HandleFunc("/user/{id}", controller.RemoveUser(db)).Methods("DELETE")

	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	fmt.Println("Server has strated on port", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
