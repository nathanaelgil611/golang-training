package services

import (
	"encoding/json"
	"final-project/database"
	"final-project/helper"
	"final-project/model"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func RegisterCommentAPI(r *mux.Router) {
	r.HandleFunc("/comments", getComments).Methods(http.MethodGet)
	r.HandleFunc("/comments/{id}", deleteComment).Methods(http.MethodDelete)
	r.HandleFunc("/comments", postComment).Methods(http.MethodPost)
	r.HandleFunc("/comments/{id}", updateComment).Methods(http.MethodPut)
}

func postComment(w http.ResponseWriter, r *http.Request) {
	var p struct {
		Message string `json:"message" validate:"required"`
		PhotoID int    `json:"photo_id" validate:"required"`
	}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&p); err != nil {
		w.Write([]byte("error decoding json body"))
		return
	}

	userID := helper.GetUserId(r)
	if userID == -1 {
		w.Write([]byte("user not found"))
		return
	}

	err := helper.Validate(p)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	res, err := database.PostComment(p.PhotoID, userID, p.Message)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	jsonData, _ := json.Marshal(&res)

	w.Header().Add("Content-Type", "application/json")
	w.Write(jsonData)
	fmt.Fprint(w)
}

func getComments(w http.ResponseWriter, r *http.Request) {

	userID := helper.GetUserId(r)
	if userID == -1 {
		w.Write([]byte("user not found"))
		return
	}

	res, err := database.GetComments(userID)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	jsonData, _ := json.Marshal(&res)

	w.Header().Add("Content-Type", "application/json")
	w.Write(jsonData)
	fmt.Fprint(w)
}

func updateComment(w http.ResponseWriter, r *http.Request) {
	var p struct {
		Message string `json:"message" validate:"required"`
	}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&p); err != nil {
		w.Write([]byte("error decoding json body"))
		return
	}

	err := helper.Validate(p)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	vars := mux.Vars(r)
	commentID, _ := strconv.Atoi(vars["id"])

	res, err := database.UpdateComment(commentID, p.Message)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	jsonData, _ := json.Marshal(&res)

	w.Header().Add("Content-Type", "application/json")
	w.Write(jsonData)
	fmt.Fprint(w)
}

func deleteComment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	commentID, _ := strconv.Atoi(vars["id"])

	err := database.DeleteComment(commentID)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	var res model.ResponseMessage
	res.Message = "Your comment has been successfully deleted"
	jsonData, _ := json.Marshal(&res)

	w.Header().Add("Content-Type", "application/json")
	w.Write(jsonData)
	fmt.Fprint(w)
}
