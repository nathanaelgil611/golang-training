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

func RegisterSocialMediaAPI(r *mux.Router) {
	r.HandleFunc("/socialmedia", getSocialMedia).Methods(http.MethodGet)
	r.HandleFunc("/socialmedia/{id}", deleteSocialMedia).Methods(http.MethodDelete)
	r.HandleFunc("/socialmedia", postSocialMedia).Methods(http.MethodPost)
	r.HandleFunc("/socialmedia/{id}", updateSocialMedia).Methods(http.MethodPut)
}

func getSocialMedia(w http.ResponseWriter, r *http.Request) {
	userID := helper.GetUserId(r)
	if userID == -1 {
		w.Write([]byte("user not found"))
		return
	}

	res, err := database.GetSocialMedia(userID)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	jsonData, _ := json.Marshal(&res)

	w.Header().Add("Content-Type", "application/json")
	w.Write(jsonData)
	fmt.Fprint(w)
}

func deleteSocialMedia(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.Write([]byte("error parsing id"))
		return
	}

	err = database.DeleteSocialMedia(id)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	var res model.ResponseMessage
	res.Message = "Your social media has been successfully deleted"
	jsonData, _ := json.Marshal(&res)

	w.Header().Add("Content-Type", "application/json")
	w.Write(jsonData)
	fmt.Fprint(w)
}

func postSocialMedia(w http.ResponseWriter, r *http.Request) {
	var p struct {
		Name string `json:"name" validate:"required"`
		URL  string `json:"social_media_url" validate:"required,url"`
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

	userID := helper.GetUserId(r)
	if userID == -1 {
		w.Write([]byte("user not found"))
		return
	}

	res, err := database.PostSocialMedia(userID, p.URL, p.Name)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	jsonData, _ := json.Marshal(&res)

	w.Header().Add("Content-Type", "application/json")
	w.Write(jsonData)
	fmt.Fprint(w)
}

func updateSocialMedia(w http.ResponseWriter, r *http.Request) {
	var p struct {
		Name string `json:"name" validate:"required"`
		URL  string `json:"social_media_url" validate:"required"`
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
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.Write([]byte("error parsing id"))
		return
	}

	res, err := database.UpdateSocialMedia(id, p.Name, p.URL)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	jsonData, _ := json.Marshal(&res)

	w.Header().Add("Content-Type", "application/json")
	w.Write(jsonData)
	fmt.Fprint(w)
}
