package controllers

import (
	"api/src/auth"
	"api/src/database"
	"api/src/errorsResponse"
	"api/src/model"
	"api/src/repositories"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
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

	if err != nil {
		errorsResponse.Error(w, http.StatusInternalServerError, err)
		return
	}

	vars := mux.Vars(r)
	ID, _ := strconv.ParseUint(vars["userid"], 10, 64)

	ExtractIDfromToken, err := auth.ExtractIDfromToken(r)
	if err != nil {
		errorsResponse.Error(w, http.StatusInternalServerError, err)
		return
	}

	fmt.Println(ID, ExtractIDfromToken, "token no controller")

	if ID != ExtractIDfromToken && ExtractIDfromToken != ID {
		errorsResponse.Error(w, http.StatusBadRequest, errors.New("user cannot update that user"))
		return
	}

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

	current_userID, err := auth.ExtractIDfromToken(r)
	if err != nil {
		errorsResponse.Error(w, http.StatusInternalServerError, err)
		return
	}

	db, err := database.Connection()
	if err != nil {
		errorsResponse.Error(w, http.StatusInternalServerError, err)
		return
	}

	fmt.Println(ID, current_userID, "token no controller")

	if ID != current_userID && current_userID != ID {
		errorsResponse.Error(w, http.StatusBadRequest, errors.New("user cannot update that user"))
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

func Follow(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	follow_user_id, err := strconv.ParseUint(vars["userid"], 10, 64)

	if err != nil {
		errorsResponse.Error(w, http.StatusInternalServerError, err)
		return
	}

	current_userID, err := auth.ExtractIDfromToken(r)
	if err != nil {
		errorsResponse.Error(w, http.StatusInternalServerError, err)
		return
	}

	if follow_user_id == current_userID {
		errorsResponse.Error(w, http.StatusBadRequest, errors.New("you cant follow yourself "))
		return
	}
	db, err := database.Connection()
	if err != nil {
		errorsResponse.Error(w, http.StatusInternalServerError, err)
		return
	}
	repository := repositories.NewRepositoryOfUsers(db)
	response := repository.FollowUser(current_userID, follow_user_id)

	if response != nil {
		errorsResponse.Error(w, http.StatusBadRequest, err)
		return
	}

	errorsResponse.JSON(w, http.StatusOK, "OK")
}

func Unfollow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	follow_user_id, err := strconv.ParseUint(vars["userid"], 10, 64)

	if err != nil {
		errorsResponse.Error(w, http.StatusInternalServerError, err)
		return
	}

	current_userID, err := auth.ExtractIDfromToken(r)
	if err != nil {
		errorsResponse.Error(w, http.StatusInternalServerError, err)
		return
	}

	if follow_user_id == current_userID {
		errorsResponse.Error(w, http.StatusBadRequest, errors.New("you cant unfollow yourself "))
		return
	}

	db, err := database.Connection()
	if err != nil {
		errorsResponse.Error(w, http.StatusInternalServerError, err)
		return
	}

	repository := repositories.NewRepositoryOfUsers(db)
	response := repository.UnFollowUser(current_userID, follow_user_id)

	if response != nil {
		errorsResponse.Error(w, http.StatusBadRequest, response)
		return
	}

	errorsResponse.JSON(w, http.StatusOK, "OK")
}

func FindFollowers(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	find_followers_byID, err := strconv.ParseUint(vars["userid"], 10, 64)

	if err != nil {
		errorsResponse.Error(w, http.StatusBadRequest, err)
		return
	}
	db, err := database.Connection()
	if err != nil {
		errorsResponse.Error(w, http.StatusInternalServerError, err)
		return
	}

	repository := repositories.NewRepositoryOfUsers(db)
	followers, err := repository.FindFollowers(find_followers_byID)

	if err != nil {
		errorsResponse.Error(w, http.StatusInternalServerError, err)
		return
	}
	errorsResponse.JSON(w, http.StatusOK, followers)
}

func WhoFollowme(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	userID, err := strconv.ParseUint(vars["userid"], 10, 64)

	if err != nil {
		errorsResponse.Error(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connection()
	if err != nil {
		errorsResponse.Error(w, http.StatusInternalServerError, err)
		return
	}

	repository := repositories.NewRepositoryOfUsers(db)
	followers, err := repository.Following(userID)

	if err != nil {
		errorsResponse.Error(w, http.StatusInternalServerError, err)
		return
	}

	errorsResponse.JSON(w, http.StatusOK, followers)
}
