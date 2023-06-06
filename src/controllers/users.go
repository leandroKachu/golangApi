package controllers

import (
	"api/src/database"
	"api/src/errorsResponse"
	"api/src/model"
	"api/src/repositories"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"reflect"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	bodyRequest, err := ioutil.ReadAll(r.Body)
	if err != nil {
		errorsResponse.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	var user model.User
	if err = json.Unmarshal(bodyRequest, &user); err != nil {
		errorsResponse.Error(w, http.StatusBadRequest, err)
		return
	}

	if err = user.Run("cadastro"); err != nil {
		errorsResponse.Error(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connection()
	if err != nil {
		errorsResponse.Error(w, http.StatusInternalServerError, err)
		return
	}

	repository := repositories.NewRepositoryOfUsers(db)
	userID, err := repository.Create(user)
	if err != nil {
		log.Fatal(err)
	}

	errorsResponse.JSON(w, http.StatusCreated, userID)

}

func FindUsers(w http.ResponseWriter, r *http.Request) {

	nameOrNick := strings.ToLower(r.URL.Query().Get("user"))
	db, err := database.Connection()
	if err != nil {
		errorsResponse.Error(w, http.StatusInternalServerError, err)
		return
	}

	repository := repositories.NewRepositoryOfUsers(db)
	users, err := repository.FindUsersByNameOrNick(nameOrNick)

	if err != nil {
		errorsResponse.Error(w, http.StatusInternalServerError, err)
		return
	}
	errorsResponse.JSON(w, http.StatusCreated, users)

	w.Write([]byte("find users"))
}

func FindUserByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	userID, _ := strconv.ParseUint(vars["userid"], 10, 64)

	fmt.Println(userID)

	db, err := database.Connection()
	if err != nil {
		errorsResponse.Error(w, http.StatusInternalServerError, err)
		return
	}

	repository := repositories.NewRepositoryOfUsers(db)
	users, err := repository.FindUserByID(userID)

	if err != nil {
		errorsResponse.Error(w, http.StatusInternalServerError, err)
		return
	}
	errorsResponse.JSON(w, http.StatusOK, users)

	w.Write([]byte("find user"))
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {

	bodyRequest, err := ioutil.ReadAll(r.Body)
	vars := mux.Vars(r)
	ID, _ := strconv.ParseUint(vars["userid"], 10, 64)

	fmt.Println("var1 = ", reflect.TypeOf(bodyRequest))
	fmt.Println(string(bodyRequest))

	var user model.User

	if err = json.Unmarshal(bodyRequest, &user); err != nil {
		errorsResponse.Error(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connection()
	if err != nil {
		errorsResponse.Error(w, http.StatusInternalServerError, err)
		return
	}

	repository := repositories.NewRepositoryOfUsers(db)
	user, err = repository.UpdateUser(user, ID)

	if err != nil {
		errorsResponse.Error(w, http.StatusInternalServerError, err)
		return
	}
	errorsResponse.JSON(w, http.StatusOK, user)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	ID, _ := strconv.ParseUint(vars["userid"], 10, 64)
	fmt.Println(ID)

	db, err := database.Connection()
	if err != nil {
		errorsResponse.Error(w, http.StatusInternalServerError, err)
		return
	}

	repository := repositories.NewRepositoryOfUsers(db)
	resultDeleted := repository.DeletedUser(ID)

	if err != nil {
		errorsResponse.Error(w, http.StatusInternalServerError, err)
		return
	}
	errorsResponse.JSON(w, http.StatusOK, resultDeleted)

	w.Write([]byte("delete user"))
}
