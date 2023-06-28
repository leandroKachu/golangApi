package controllers

import (
	"api/src/auth"
	"api/src/database"
	"api/src/errorsResponse"
	"api/src/model"
	"api/src/repositories"
	"api/src/security"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

func Login(w http.ResponseWriter, r *http.Request) {
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

	// if err = user.Run("cadastro"); err != nil {
	// 	errorsResponse.Error(w, http.StatusBadRequest, err)
	// 	return
	// }

	db, err := database.Connection()
	if err != nil {
		errorsResponse.Error(w, http.StatusInternalServerError, err)
		return
	}

	repository := repositories.NewRepositoryOfUsers(db)
	userPassHash, err := repository.CheckEmail(user.Email)

	if err != nil {
		errorsResponse.Error(w, http.StatusInternalServerError, err)
		return

	}
	// valid if the pass in body  is the same as the repository
	if err := security.ValidPass(userPassHash.Password, user.Password); err != nil {
		errorsResponse.Error(w, http.StatusUnauthorized, err)
		fmt.Println("senha incorrect password")
	}

	//Create a token by id of user
	token, err := auth.GenToken(userPassHash.ID)
	if err != nil {
		errorsResponse.Error(w, http.StatusInternalServerError, err)
		return
	}

	userID := strconv.FormatInt(userPassHash.ID, 10)

	errorsResponse.JSON(w, http.StatusOK, model.DataAuth{ID: userID, Token: token})

}
