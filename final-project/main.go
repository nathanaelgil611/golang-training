package main

import (
	"net/http"

	"final-project/config"
	"final-project/database"
	"final-project/helper"
	"final-project/services"

	"github.com/gorilla/mux"
	"github.com/ilyakaznacheev/cleanenv"
)

// GET /users (untuk get all users)
// GET /users/id (untuk get user by id)
// POST /users (untuk create user)
// PUT /users/id (untuk update user)
// DELETE /users/id (untuk delete user by id)
var cfgPORT config.ConfigPort

func main() {

	err := cleanenv.ReadConfig(".env", &cfgPORT)
	if err != nil {
		panic(err)
	}

	database.DatabaseInit()
	defer database.CloseDatabase()

	r := mux.NewRouter()

	services.RegisterUserAPI(r)
	services.RegisterPhotoAPI(r)
	services.RegisterCommentAPI(r)
	services.RegisterSocialMediaAPI(r)

	http.Handle("/", helper.Middleware(r))
	http.ListenAndServe(cfgPORT.Port, nil)

}

// POST /photos (untuk create photo)
// GET /photos (untuk get all photos)
// GET /photos/id (untuk get photo by id)
// PUT /photos/id (untuk update photo)
// DELETE /photos/id (untuk delete photo by id)
