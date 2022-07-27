package database

import (
	"fmt"
	"sesi-10/model"

	"golang.org/x/crypto/bcrypt"
)

func CreateUser(id, age int, name, email, password string) error {
	// var user = model.User{}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		panic(err)
	}
	sqlStatement := `INSERT INTO users (id, username, email, password, age) VALUES($1, $2, $3, $4, $5)`

	_, err = db.Exec(sqlStatement, id, name, email, string(hashedPassword), age)
	if err != nil {
		panic(err)
	}
	return err
}

func GetAllUser() ([]model.User, error) {
	var userList = []model.User{}

	sqlStatement := `SELECT * FROM users`
	rows, err := db.Query(sqlStatement)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var user = model.User{}
		err := rows.Scan(&user.UserID, &user.Username, &user.Email, &user.Password, &user.Age)
		if err != nil {
			panic(err)
		}
		userList = append(userList, user)
	}

	return userList, err
}

func GetUser(id int) (model.User, error) {
	var user = model.User{}

	sqlStatement := `SELECT * FROM users WHERE id = $1`
	rows, err := db.Query(sqlStatement, id)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&user.UserID, &user.Username, &user.Email, &user.Password, &user.Age)
	}

	return user, err
}

func UpdateUser(id int, age int, name, email, password string) error {
	sqlStatement := `UPDATE users SET username = $1, email = $2, password = $3, age = $4 WHERE id = $5`
	_, err := db.Exec(sqlStatement, name, email, password, age, id)
	return err
}

func DeleteUser(id int) error {
	sqlStatement := `DELETE FROM users WHERE id = $1`
	_, err := db.Exec(sqlStatement, id)
	return err
}

func Login(email, password string) bool {
	var pass string

	sqlStatement := `SELECT password FROM users WHERE email = $1`
	rows, err := db.Query(sqlStatement, email)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&pass)
		fmt.Println("pass ", pass)
	}

	err = bcrypt.CompareHashAndPassword([]byte(pass), []byte(password))
	// fmt.Println(bcrypt.CompareHashAndPassword([]byte(pass), []byte(password)))

	if err != nil {
		return false
	}

	return true
}
