package example

import (
	"database/sql"
	"fmt"
	"sesi-3/model"

	_ "github.com/lib/pq"
)

func main() {

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")
	CreateUser()
}

// var (
// 	db  *sql.DB
// 	err error
// )

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "postgres"
)

func CreateUser() {
	var user = model.User{}

	sqlStatement := `INSERT INTO users (id, username, email, password, age) VALUES($1, $2, $3, $4, $5) RETURNING *`
	err := db.QueryRow(sqlStatement, 1, "anto", "anto@gmail.com", "923k0a9123kj@#!j", 21).Scan(&user.UserID, &user.Username, &user.Email, &user.Password, &user.Age)
	if err != nil {
		panic(err)
	}

	fmt.Println("User created:", user)
}

/* async example
// 	var users = []string{"John", "Paul", "George", "Ringo"}

	// 	for index, element := range users {
	// 		// index is the index where we are
	// 		// element is the element from someSlice for where we are
	// 		go func(index int, element string) {
	// 			fmt.Printf("%d: %s\n", index, element)

	// 		}(index, element)
	// 	}
	// 	// fmt.Println(users)
	// 	// fmt.Println("Hello World")
	// 	time.Sleep(1 * time.Second)
*/
