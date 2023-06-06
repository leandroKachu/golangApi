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

	if err := security.ValidPass(userPassHash.Password, user.Password); err != nil {
		errorsResponse.Error(w, http.StatusUnauthorized, err)
		fmt.Println("senha incorrect password")
	}

	token, _ := auth.GenToken(userPassHash.ID)
	fmt.Println(token)

	w.Write([]byte(token))

	// errorsResponse.JSON(w, http.StatusCreated, userID)

}
