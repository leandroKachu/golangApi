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
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func CreatePost(w http.ResponseWriter, r *http.Request) {
	userID, err := auth.ExtractIDfromToken(r)
	if err != nil {
		errorsResponse.Error(w, http.StatusBadRequest, err)
		return
	}
	var post model.Post
	post.AuthorID = userID

	bodyRequest, err := ioutil.ReadAll(r.Body)

	if err != nil {
		errorsResponse.Error(w, http.StatusBadRequest, err)
		return
	}

	if err = json.Unmarshal(bodyRequest, &post); err != nil {
		errorsResponse.Error(w, http.StatusBadRequest, err)
		return
	}
	db, err := database.Connection()
	if err != nil {
		errorsResponse.Error(w, http.StatusInternalServerError, err)
		return
	}

	repository := repositories.NewRepositoryOfPosts(db)
	infoPost, err := repository.CreatePost(post)

	if err != nil {
		errorsResponse.Error(w, http.StatusInternalServerError, err)
		return
	}
	errorsResponse.JSON(w, http.StatusOK, infoPost.Title)

}

func GetPostbyID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	postID, err := strconv.ParseUint(vars["postID"], 10, 64)
	if err != nil {
		errorsResponse.Error(w, http.StatusInternalServerError, err)
		return
	}

	db, err := database.Connection()
	if err != nil {
		errorsResponse.Error(w, http.StatusInternalServerError, err)
		return
	}
	repository := repositories.NewRepositoryOfPosts(db)

	postReturn, err := repository.FindPost(postID)

	if err != nil {
		errorsResponse.Error(w, http.StatusInternalServerError, err)
		return
	}
	errorsResponse.JSON(w, http.StatusOK, postReturn)
}

func GetPosts(w http.ResponseWriter, r *http.Request) {
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
	repository := repositories.NewRepositoryOfPosts(db)
	posts, _ := repository.FindPosts(current_userID)

	// var result []model.Post
	// for _, post := range posts {
	// 	result = append(result, *post)
	// }

	errorsResponse.JSON(w, http.StatusOK, posts)

}

func UpdatePost(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	current_userID, err := auth.ExtractIDfromToken(r)

	if err != nil {
		errorsResponse.Error(w, http.StatusInternalServerError, err)
		return
	}

	postID, _ := strconv.ParseUint(vars["postID"], 10, 64)
	db, err := database.Connection()
	if err != nil {
		errorsResponse.Error(w, http.StatusInternalServerError, err)
		return
	}
	repository := repositories.NewRepositoryOfPosts(db)

	authorID, err := repository.CheckPostAuthorID(postID)
	if err != nil {
		errorsResponse.Error(w, http.StatusInternalServerError, err)
		return

	}

	if authorID != current_userID {
		errorsResponse.Error(w, http.StatusBadRequest, errors.New("you can t update a post is not your"))
		return
	}

	var post model.Post
	bodyRequest, err := ioutil.ReadAll(r.Body)
	if err != nil {
		errorsResponse.Error(w, http.StatusBadRequest, err)
		return
	}

	if err = json.Unmarshal(bodyRequest, &post); err != nil {
		errorsResponse.Error(w, http.StatusBadRequest, err)
		return
	}

	result, _ := repository.UpdatePost(post, postID)
	errorsResponse.JSON(w, http.StatusOK, result)
}

func DeletePost(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	current_userID, err := auth.ExtractIDfromToken(r)

	if err != nil {
		errorsResponse.Error(w, http.StatusInternalServerError, err)
		return
	}

	postID, _ := strconv.ParseUint(vars["postID"], 10, 64)
	db, err := database.Connection()
	if err != nil {
		errorsResponse.Error(w, http.StatusInternalServerError, err)
		return
	}
	repository := repositories.NewRepositoryOfPosts(db)

	authorID, err := repository.CheckPostAuthorID(postID)
	if err != nil {
		errorsResponse.Error(w, http.StatusInternalServerError, err)
		return

	}

	if authorID != current_userID {
		errorsResponse.Error(w, http.StatusBadRequest, errors.New("you can t delete a post is not your"))
		return
	}

	result, err := repository.DeletePost(postID)
	if err != nil {
		errorsResponse.Error(w, http.StatusInternalServerError, err)
		return
	}
	errorsResponse.JSON(w, http.StatusOK, result)
}

func GetPostByIDuser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	authorID, _ := strconv.ParseUint(vars["userID"], 10, 64)
	db, err := database.Connection()
	if err != nil {
		errorsResponse.Error(w, http.StatusInternalServerError, err)
		return
	}
	repository := repositories.NewRepositoryOfPosts(db)

	results, err := repository.GetPostByAuthorID(authorID)

	fmt.Println(results)
	if err != nil {
		errorsResponse.Error(w, http.StatusInternalServerError, err)
		return
	}
	errorsResponse.JSON(w, http.StatusOK, results)

}

func LikePost(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	postID, _ := strconv.ParseInt(vars["postID"], 10, 64)
	db, err := database.Connection()
	if err != nil {
		errorsResponse.Error(w, http.StatusInternalServerError, err)
		return
	}
	repository := repositories.NewRepositoryOfPosts(db)

	resultxt, err := repository.LikePost(postID)
	if err != nil {
		errorsResponse.Error(w, http.StatusInternalServerError, err)
		return
	}
	errorsResponse.JSON(w, http.StatusOK, resultxt)
}

func Unlike(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	postID, _ := strconv.ParseInt(vars["postID"], 10, 64)
	db, err := database.Connection()
	if err != nil {
		errorsResponse.Error(w, http.StatusInternalServerError, err)
		return
	}
	repository := repositories.NewRepositoryOfPosts(db)

	resultxt, err := repository.UnLike(postID)
	if err != nil {
		errorsResponse.Error(w, http.StatusInternalServerError, err)
		return
	}
	errorsResponse.JSON(w, http.StatusOK, resultxt)
}
