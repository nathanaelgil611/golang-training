package services

import (
	"fmt"
	"sesi-2/model"
)

var userList = make(map[int]model.User)

type UserServiceIface interface {
	Register(user *model.User)
	GetAllUser() map[int]model.User
	GetUser(userID int) model.User
	Delete(userID int)
	Update(User *model.User, userID int)
}

type UserSvc struct {
	ListUser map[int]model.User
}

func NewUserService() UserServiceIface {
	list := &userList
	return &UserSvc{*list}
}

func (u *UserSvc) Register(user *model.User) {
	u.ListUser[user.UserID] = *user

	fmt.Println(u)
}

func (u *UserSvc) GetUser(userID int) model.User {
	return u.ListUser[userID]
}

func (u *UserSvc) GetAllUser() map[int]model.User {
	return u.ListUser
}

func (u *UserSvc) Delete(userID int) {
	delete(u.ListUser, userID)
}

func (u *UserSvc) Update(user *model.User, userID int) {
	u.ListUser[userID] = *user
}
