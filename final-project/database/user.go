package database

import (
	"final-project/model"
	"fmt"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func CreateUser(age int, name, email, password string) error {
	// var user = model.User{}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		fmt.Errorf("Error creating user: %v", err)
	}
	sqlStatement := `INSERT INTO "user" (username, email, password, age, created_at, updated_at) VALUES($1, $2, $3, $4, $5, $6)`
	_, err = db.Exec(sqlStatement, name, email, string(hashedPassword), age, time.Now(), time.Now())
	if err != nil {
		fmt.Errorf("Error creating user: %v", err)
	}
	return err
}

func GetUser(id int) (model.User, error) {
	var user = model.User{}

	sqlStatement := `SELECT user_id, username, email, age, updated_at FROM "user" WHERE user_id = $1`
	rows, err := db.Query(sqlStatement, id)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&user.UserID, &user.Username, &user.Email, &user.Age, &user.UpdatedAt)
	}

	return user, err
}

func DeleteUser(user_id int) error {
	sqlStatement := `DELETE FROM "user" WHERE user_id = $1`
	_, err := db.Exec(sqlStatement, user_id)
	return err
}

func Login(email, password string) (bool, error) {
	var pass string

	sqlStatement := `SELECT password FROM "user" WHERE email = $1`
	rows, err := db.Query(sqlStatement, email)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&pass)
	}

	err = bcrypt.CompareHashAndPassword([]byte(pass), []byte(password))

	return err == nil, err
}

func GetUserIdFromEmail(email string) (int, error) {
	var userId int

	sqlStatement := `SELECT user_id FROM "user" WHERE email = $1`
	rows, err := db.Query(sqlStatement, email)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&userId)
	}

	return userId, err
}
