package services

import (
	"sesi-10/database"
	"sesi-10/model"
)

// var userList = make([]model.User)

type UserServiceIface interface {
	Register(user *model.User)
	GetAllUser() []model.User
	GetUser(userID int) model.User
	Delete(userID int)
	Update(User *model.User, userID int)
	Login(email, password string) bool
}

type UserSvc struct {
	ListUser []model.User
}

func NewUserService() UserServiceIface {
	var list []model.User
	return &UserSvc{list}
}

func (u *UserSvc) Register(user *model.User) {
	// u.ListUser[user.UserID] = *user
	err := database.CreateUser(user.UserID, user.Age, user.Username, user.Email, user.Password)
	if err != nil {
		panic(err)
	}

	// fmt.Println(u)
}

func (u *UserSvc) GetUser(userID int) model.User {
	user, err := database.GetUser(userID)
	if err != nil {
		panic(err)
	}
	return user
}

func (u *UserSvc) GetAllUser() []model.User {
	listUser, err := database.GetAllUser()
	if err != nil {
		panic(err)
	}
	return listUser
}

func (u *UserSvc) Delete(userID int) {
	// delete(u.ListUser, userID)
	err := database.DeleteUser(userID)
	if err != nil {
		panic(err)
	}

}

func (u *UserSvc) Update(user *model.User, userID int) {
	err := database.UpdateUser(user.UserID, user.Age, user.Username, user.Email, user.Password)
	if err != nil {
		panic(err)
	}
}

func (u *UserSvc) Login(email, password string) bool {

	return database.Login(email, password)
}
