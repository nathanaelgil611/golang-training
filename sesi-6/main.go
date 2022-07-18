package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"sesi-6/database"
	"sesi-6/model"
	"sesi-6/services"

	"github.com/golang-jwt/jwt/v4"
	"github.com/gorilla/mux"
)

var PORT = ":8080"
var USERNAME = "admin"
var PASSWORD = "admin"
var APPLICATION_NAME = "My Simple JWT App"
var LOGIN_EXPIRATION_DURATION = time.Duration(5) * time.Hour
var JWT_SIGNING_METHOD = jwt.SigningMethodHS256
var JWT_SIGNATURE_KEY = []byte("s3cr3t")

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

	r.HandleFunc("/login", login).Methods(http.MethodPut)

	http.Handle("/", Middleware(r))
	// r.Use(Middleware)
	http.ListenAndServe(PORT, nil)

}

func IsAuthorized(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		if r.Header["Token"] == nil {
			var err error
			err = fmt.Errorf("Token not found")
			http.Error(w, err.Error(), 500)
			return
		}

		var mySigningKey = []byte(JWT_SIGNATURE_KEY)

		token, err := jwt.Parse(r.Header["Token"][0], func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("There was an error in parsing")
			}

			return mySigningKey, nil
		})
		fmt.Println("TOKEN IS AUTHORIZED ", token)

		if err != nil {
			var err error
			err = fmt.Errorf("Your Token has been expired")
			json.NewEncoder(w).Encode(err)
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			if claims["role"] == "admin" {

				r.Header.Set("Role", "admin")
				handler.ServeHTTP(w, r)
				return

			} else if claims["role"] == "user" {

				r.Header.Set("Role", "user")
				handler.ServeHTTP(w, r)
				return
			}
		}
		var reserr error
		reserr = fmt.Errorf("Not Authorized")

		json.NewEncoder(w).Encode(reserr)
		handler.ServeHTTP(w, r)
	}
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

func GenerateJWT(email, role string) (string, error) {
	var mySigningKey = []byte(JWT_SIGNATURE_KEY)
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["email"] = email
	claims["role"] = role
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	tokenString, err := token.SignedString(mySigningKey)

	if err != nil {
		fmt.Errorf("Something Went Wrong: %s", err.Error())
		return "", err
	}
	return tokenString, nil
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
	fmt.Println(p)
	userSvc := services.NewUserService()
	res := userSvc.Login(p.Email, p.Password)
	var response struct {
		Message string `json:"message"`
		Token   string `json:"token"`
	}

	if res {
		token, err := GenerateJWT(p.Email, "admin")
		if err != nil {
			w.Write([]byte("error generating token"))
			return
		}
		response.Token = token
		response.Message = "login success"
		jsonData, _ := json.Marshal(&response)
		w.Header().Add("Content-Type", "application/json")
		w.Write(jsonData)
		fmt.Fprint(w)
	} else {
		response.Message = "Incorrect uername or password"
		jsonData, _ := json.Marshal(&response)
		http.Error(w, string(jsonData), http.StatusBadRequest)
		return
	}
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
	fmt.Println(user)
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

	userSvc.Update(&user, userID)

	w.Header().Add("Content-Type", "application/json")

	fmt.Fprint(w, "Success Update")
}
