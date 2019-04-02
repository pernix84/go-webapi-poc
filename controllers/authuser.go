package controllers

import (
	"database/sql"
	"encoding/json"
	"go-webapi-poc/models"
	authuserrepository "go-webapi-poc/repository/authuser"
	"go-webapi-poc/utils"
	"log"
	"net/http"
	"strconv"

	"github.com/davecgh/go-spew/spew"
	"github.com/gorilla/mux"
)

var users []models.AuthUser

type Controller struct{}

func logFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func (c Controller) GetUsers(db *sql.DB) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		var user models.AuthUser
		var error models.Error

		users = []models.AuthUser{}
		userRepo := authuserrepository.AuthUserRepository{}
		users, err := userRepo.GetUsers(db, user, users)

		if err != nil {
			error.Message = "Server error"
			utils.SendError(w, http.StatusInternalServerError, error)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		utils.SendSuccess(w, users)
	}
}

func (c Controller) GetUser(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user models.AuthUser
		var error models.Error

		params := mux.Vars(r)

		users = []models.AuthUser{}
		userRepo := authuserrepository.AuthUserRepository{}

		id, err := strconv.Atoi(params["id"])

		if err != nil {
			error.Message = "Incorrect id."
			utils.SendError(w, http.StatusBadRequest, error)
			return
		}

		user, err = userRepo.GetUser(db, user, id)

		if err != nil {
			error.Message = "Server error."
			utils.SendError(w, http.StatusInternalServerError, error)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		utils.SendSuccess(w, user)
	}
}

func (c Controller) AddUser(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user models.AuthUser
		var userID int
		var error models.Error

		json.NewDecoder(r.Body).Decode(&user)

		if user.FullName == "" || user.OAuthId == "" {
			error.Message = "Enter missing fields."
			utils.SendError(w, http.StatusBadRequest, error)
			return
		}

		userRepo := authuserrepository.AuthUserRepository{}
		userID, err := userRepo.AddUser(db, user)

		if err != nil {
			error.Message = "Server error"
			utils.SendError(w, http.StatusInternalServerError, error)
			return
		}

		w.Header().Set("Content-Type", "text/plain")
		utils.SendSuccess(w, userID)
	}
}

func (c Controller) UpdateUser(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user models.AuthUser
		var error models.Error

		json.NewDecoder(r.Body).Decode(&user)

		if user.FullName == "" || user.OAuthId == "" {
			error.Message = "Enter all fields."
			utils.SendError(w, http.StatusBadRequest, error)
			return
		}

		userRepo := authuserrepository.AuthUserRepository{}
		rowsUpdated, err := userRepo.UpdateUser(db, user)

		spew.Dump(err)

		if err != nil {
			error.Message = "Server error"
			utils.SendError(w, http.StatusInternalServerError, error)
			return
		}

		w.Header().Set("Content-Type", "text/plain")
		utils.SendSuccess(w, rowsUpdated)
	}
}

func (c Controller) RemoveUser(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var error models.Error
		params := mux.Vars(r)

		userRepo := authuserrepository.AuthUserRepository{}

		id, err := strconv.Atoi(params["id"])

		if err != nil {
			error.Message = "Incorrect id."
			utils.SendError(w, http.StatusBadRequest, error)
			return
		}

		rowsDeleted, err := userRepo.RemoveUser(db, id)

		w.Header().Set("Content-Type", "text/plain")
		utils.SendSuccess(w, rowsDeleted)
	}
}
