package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"

	"sesi-5/database"
	"sesi-5/model"
	"sesi-5/services"

	"github.com/gorilla/mux"
)

var PORT = ":8080"
var USERNAME = "admin"
var PASSWORD = "admin"

// GET /users (untuk get all users)
// GET /users/id (untuk get user by id)
// POST /users (untuk create user)
// PUT /users/id (untuk update user)
// DELETE /users/id (untuk delete user by id)

func main() {

	database.DatabaseInit()
	defer database.CloseDatabase()

	defer func() {
		if err := recover(); err != nil {
			log.Println("panic occurred:", err)
		}
	}()

	r := mux.NewRouter()
	r.HandleFunc("/users", getAllUser).Methods(http.MethodGet)
	r.HandleFunc("/users/{id}", getUser).Methods(http.MethodGet)
	r.HandleFunc("/users", registerUser).Methods(http.MethodPost)
	r.HandleFunc("/users/{id}", updateUser).Methods(http.MethodPut)
	r.HandleFunc("/users/{id}", deleteUser).Methods(http.MethodDelete)
	r.HandleFunc("/users-url", getUserURL).Methods(http.MethodGet)

	http.Handle("/", Middleware(r))
	http.ListenAndServe(PORT, nil)

}

func Middleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		username, password, ok := r.BasicAuth()
		if !ok {
			w.Write([]byte(`something went wrong`))
			return
		}

		isValid := (username == USERNAME) && (password == PASSWORD)
		if !isValid {
			w.Write([]byte(`wrong username/password`))
			return
		}

		h.ServeHTTP(w, r)
	})
}

func getUserURL(w http.ResponseWriter, r *http.Request) {
	response, err := http.Get("https://random-data-api.com/api/users/random_user?size=10")

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var responseObject []model.UserURL
	json.Unmarshal(responseData, &responseObject)

	jsonData, _ := json.Marshal(&responseObject)

	w.Header().Add("Content-Type", "application/json")
	w.Write(jsonData)
	fmt.Fprint(w)
}

func getAllUser(w http.ResponseWriter, r *http.Request) {
	user := services.NewUserService()
	newUser := user.GetAllUser()

	jsonData, _ := json.Marshal(&newUser)
	w.Header().Add("Content-Type", "application/json")
	w.Write(jsonData)
	fmt.Fprint(w)
}

func getUser(w http.ResponseWriter, r *http.Request) {
	user := services.NewUserService()

	vars := mux.Vars(r)
	userID, _ := strconv.Atoi(vars["id"])
	fmt.Println(userID)

	fmt.Println(user.GetUser(userID))
	newUser := user.GetUser(userID)

	jsonData, _ := json.Marshal(&newUser)
	w.Header().Add("Content-Type", "application/json")
	w.Write(jsonData)
	fmt.Fprint(w)
}

func registerUser(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)
	var user model.User
	if err := decoder.Decode(&user); err != nil {
		w.Write([]byte("error decoding json body"))
		return
	}
	userSvc := services.NewUserService()

	userSvc.Register(&user)

	w.Header().Add("Content-Type", "application/json")

	fmt.Fprint(w, "Success Register")
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	user := services.NewUserService()

	vars := mux.Vars(r)
	userID, _ := strconv.Atoi(vars["id"])
	fmt.Println(userID)

	user.Delete(userID)

	w.Header().Add("Content-Type", "application/json")
	fmt.Fprint(w, "Success Delete")
}

func updateUser(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)
	var user model.User
	if err := decoder.Decode(&user); err != nil {
		w.Write([]byte("error decoding json body"))
		return
	}
	userSvc := services.NewUserService()

	vars := mux.Vars(r)
	userID, _ := strconv.Atoi(vars["id"])
	fmt.Println(userID)

	userSvc.Update(&user, userID)

	w.Header().Add("Content-Type", "application/json")

	fmt.Fprint(w, "Success Update")
}
