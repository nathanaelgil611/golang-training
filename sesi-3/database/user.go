package database

import "sesi-3/model"

func CreateUser(id, age int, name, email, password string) error {
	// var user = model.User{}

	sqlStatement := `INSERT INTO users (id, username, email, password, age) VALUES($1, $2, $3, $4, $5)`
	_, err := db.Exec(sqlStatement, id, name, email, password, age)
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
