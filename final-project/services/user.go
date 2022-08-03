package services

import (
	"encoding/json"
	"final-project/database"
	"final-project/helper"
	"final-project/model"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

func RegisterUserAPI(r *mux.Router) {
	r.HandleFunc("/users", getUser).Methods(http.MethodGet)
	r.HandleFunc("/users", deleteUser).Methods(http.MethodDelete)
	r.HandleFunc("/users/login", login).Methods(http.MethodPost)
	r.HandleFunc("/users/register", registerUser).Methods(http.MethodPost)
}

func login(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var p struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := decoder.Decode(&p); err != nil {
		w.Write([]byte("error decoding json body"))
		return
	}

	res, err := database.Login(p.Email, p.Password)
	if err != nil {
		http.Error(w, "Error while logging in", http.StatusBadRequest)
		return
	}

	var response struct {
		Token string `json:"token"`
	}

	if res {
		token, err := helper.GenerateJWT(p.Email, "admin")
		if err != nil {
			w.Write([]byte("error generating token"))
			return
		}
		response.Token = token
		jsonData, _ := json.Marshal(&response)
		w.Header().Add("Content-Type", "application/json")
		w.Write(jsonData)
		fmt.Fprint(w)
	} else {
		jsonData, _ := json.Marshal(&response)
		http.Error(w, string(jsonData), http.StatusBadRequest)
		return
	}
}

func getUser(w http.ResponseWriter, r *http.Request) {
	token, err := helper.GetTokenString(r)
	if err != nil {
		http.Error(w, "Error while getting token", http.StatusBadRequest)
		return
	}
	userId := helper.GetUserIdFromToken(token)

	newUser, err := database.GetUser(userId)
	if err != nil {
		http.Error(w, "Error while getting user", http.StatusBadRequest)
		return
	}

	jsonData, _ := json.Marshal(&newUser)
	w.Header().Add("Content-Type", "application/json")
	w.Write(jsonData)
	fmt.Fprint(w)
}

func registerUser(w http.ResponseWriter, r *http.Request) {
	validate := validator.New()
	decoder := json.NewDecoder(r.Body)
	var user model.User
	if err := decoder.Decode(&user); err != nil {
		w.Write([]byte("error decoding json body"))
		return
	}

	err := validate.Struct(user)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	err = database.CreateUser(user.Age, user.Username, user.Email, user.Password)
	if err != nil {
		http.Error(w, "Error while registering user", http.StatusBadRequest)
		return
	}

	var response struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		Age      int    `json:"age"`
	}

	js, _ := json.Marshal(&user)

	json.Unmarshal(js, &response)

	jsonData, _ := json.Marshal(&response)

	w.Header().Add("Content-Type", "application/json")
	w.Write(jsonData)
	fmt.Fprint(w)
}

func deleteUser(w http.ResponseWriter, r *http.Request) {

	token, err := helper.GetTokenString(r)
	if err != nil {
		http.Error(w, "Error while getting token", http.StatusBadRequest)
		return
	}
	userId := helper.GetUserIdFromToken(token)
	if userId == -1 {
		http.Error(w, "Invalid token", http.StatusBadRequest)
		return
	}

	err = database.DeleteUser(userId)
	if err != nil {
		http.Error(w, "Error while deleting user", http.StatusBadRequest)
		return
	}

	var res model.ResponseMessage
	res.Message = "Account has been successfully deleted"
	jsonData, _ := json.Marshal(&res)

	w.Header().Add("Content-Type", "application/json")
	w.Write(jsonData)
	fmt.Fprint(w)
}
