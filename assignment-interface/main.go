package main

import (
	"fmt"
)

type User struct {
	Id       int
	Username string
	Email    string
	Password string
	Age      int
}

type UserServiceIface interface {
	Register(user *User)
}

type UserSvc struct {
	ListUser map[string]User
}

func NewUserService() UserServiceIface {
	list := make(map[string]User)
	list["anton"] = User{
		Id:       0,
		Username: "anton",
		Email:    "anton@gmail.com",
		Password: "000000",
		Age:      24,
	}
	list["rudi"] = User{
		Id:       1,
		Username: "rudi",
		Email:    "rudi@gmail.com",
		Password: "000000",
		Age:      23,
	}
	return &UserSvc{ListUser: list}
}

func (u *UserSvc) Register(user *User) {
	if _, ok := u.ListUser[user.Username]; ok {
		fmt.Println("Failed, username already exist")
	} else {
		fmt.Println("Success register user")
	}

	fmt.Println(user)
}

func main() {
	userSvc := NewUserService()
	userSvc.Register(&User{
		Id:       0,
		Username: "kevin",
		Email:    "kevin@gmail.com",
		Password: "000000",
		Age:      21,
	})

}
