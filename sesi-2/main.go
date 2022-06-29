package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type User struct {
	UserID   int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Age      int    `json:"age"`
}

var userList = make(map[int]User)

type UserServiceIface interface {
	Register(user *User)
	GetAllUser() map[int]User
	GetUser(userID int) User
	Delete(userID int)
	Update(User *User, userID int)
}

type UserSvc struct {
	ListUser map[int]User
}

func NewUserService() UserServiceIface {
	list := &userList
	return &UserSvc{*list}
}

func (u *UserSvc) Register(user *User) {
	// if _, ok := u.ListUser[user.UserID]; ok {
	// 	fmt.Println("Failed, username already exist")
	// } else {
	// 	fmt.Println("Success register user")
	// }
	u.ListUser[user.UserID] = *user

	fmt.Println(u)
}

func (u *UserSvc) GetUser(userID int) User {
	return u.ListUser[userID]
}

func (u *UserSvc) GetAllUser() map[int]User {
	return u.ListUser
}

func (u *UserSvc) Delete(userID int) {
	delete(u.ListUser, userID)
}

func (u *UserSvc) Update(user *User, userID int) {
	u.ListUser[userID] = *user
}

var PORT = ":8080"

// GET /users (untuk get all users)
// GET /users/id (untuk get user by id)
// POST /users (untuk create user)
// PUT /users/id (untuk update user)
// DELETE /users/id (untuk delete user by id)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/users", getAllUser).Methods(http.MethodGet)
	r.HandleFunc("/users/{id}", getUser).Methods(http.MethodGet)
	r.HandleFunc("/users", registerUser).Methods(http.MethodPost)
	r.HandleFunc("/users/{id}", updateUser).Methods(http.MethodPut)
	r.HandleFunc("/users/{id}", deleteUser).Methods(http.MethodDelete)

	http.Handle("/", r)
	http.ListenAndServe(PORT, nil)
}

func getAllUser(w http.ResponseWriter, r *http.Request) {
	user := NewUserService()
	newUser := user.GetAllUser()

	jsonData, _ := json.Marshal(&newUser)
	w.Header().Add("Content-Type", "application/json")
	w.Write(jsonData)
	fmt.Fprint(w)
}

func getUser(w http.ResponseWriter, r *http.Request) {
	user := NewUserService()

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
	var user User
	if err := decoder.Decode(&user); err != nil {
		w.Write([]byte("error decoding json body"))
		return
	}
	userSvc := NewUserService()

	userSvc.Register(&user)

	w.Header().Add("Content-Type", "application/json")

	fmt.Fprint(w, "Success Register")
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	user := NewUserService()

	vars := mux.Vars(r)
	userID, _ := strconv.Atoi(vars["id"])
	fmt.Println(userID)

	user.Delete(userID)

	w.Header().Add("Content-Type", "application/json")
	fmt.Fprint(w, "Success Delete")
}

func updateUser(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)
	var user User
	if err := decoder.Decode(&user); err != nil {
		w.Write([]byte("error decoding json body"))
		return
	}
	userSvc := NewUserService()

	vars := mux.Vars(r)
	userID, _ := strconv.Atoi(vars["id"])
	fmt.Println(userID)

	userSvc.Update(&user, userID)

	w.Header().Add("Content-Type", "application/json")

	fmt.Fprint(w, "Success Update")
}
