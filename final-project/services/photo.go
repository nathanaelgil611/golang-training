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

func RegisterPhotoAPI(r *mux.Router) {
	r.HandleFunc("/photos", getPhotos).Methods(http.MethodGet)
	// r.HandleFunc("/photos/{id}", getPhoto).Methods(http.MethodGet)
	r.HandleFunc("/photos", postPhoto).Methods(http.MethodPost)
	r.HandleFunc("/photos/{id}", editPhoto).Methods(http.MethodPut)
	r.HandleFunc("/photos/{id}", deletePhoto).Methods(http.MethodDelete)
}

func editPhoto(w http.ResponseWriter, r *http.Request) {
	var p struct {
		Title   string `json:"title" validate:"required"`
		Caption string `json:"caption" validate:"required"`
		URL     string `json:"photo_url" validate:"required,url"`
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
	photoId, _ := strconv.Atoi(vars["id"])

	res, err := database.EditPhoto(photoId, p.Caption, p.URL, p.Title)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	response := struct {
		ID        int    `json:"photo_id"`
		Title     string `json:"title"`
		Caption   string `json:"caption"`
		URL       string `json:"photo_url"`
		UserID    int    `json:"user_id"`
		UpdatedAt string `json:"updated_at"`
	}{ID: res.PhotoID, Title: res.Title, Caption: res.Caption, URL: res.URL, UserID: res.UserID, UpdatedAt: res.UpdatedAt}

	jsonData, _ := json.Marshal(&response)
	w.Header().Add("Content-Type", "application/json")
	w.Write(jsonData)
	fmt.Fprint(w)
}

func deletePhoto(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	photoId, _ := strconv.Atoi(vars["id"])

	err := database.DeletePhoto(photoId)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	var res model.ResponseMessage
	res.Message = "Photo has been successfully deleted"
	jsonData, _ := json.Marshal(&res)

	w.Header().Add("Content-Type", "application/json")
	w.Write(jsonData)
	fmt.Fprint(w)
}

func getPhotos(w http.ResponseWriter, r *http.Request) {
	token, _ := helper.GetTokenString(r)
	userId := helper.GetUserIdFromToken(token)

	photos, err := database.GetPhotos(userId)
	if err != nil {
		http.Error(w, "Error while getting photos", http.StatusBadRequest)
		return
	}

	jsonData, _ := json.Marshal(photos)
	w.Header().Add("Content-Type", "application/json")
	w.Write(jsonData)
	fmt.Fprint(w)
}

func postPhoto(w http.ResponseWriter, r *http.Request) {
	var p struct {
		Title   string `json:"title" validate:"required"`
		Caption string `json:"caption" validate:"required"`
		URL     string `json:"photo_url" validate:"required,url"`
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

	token, _ := helper.GetTokenString(r)
	userId := helper.GetUserIdFromToken(token)

	res, err := database.PostPhoto(userId, p.Caption, p.Title, p.URL)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	response := struct {
		ID        int    `json:"photo_id"`
		Title     string `json:"title"`
		URL       string `json:"photo_url"`
		UserID    int    `json:"user_id"`
		CreatedAt string `json:"created_at"`
	}{ID: res.PhotoID, Title: res.Title, URL: res.URL, UserID: res.UserID, CreatedAt: res.CreatedAt}

	jsonData, _ := json.Marshal(&response)
	w.Header().Add("Content-Type", "application/json")
	w.Write(jsonData)
	fmt.Fprint(w)
}
