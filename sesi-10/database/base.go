package database

import (
	"database/sql"
	"fmt"
	"sesi-10/config"

	"github.com/ilyakaznacheev/cleanenv"
	_ "github.com/lib/pq"
)

var (
	db  *sql.DB
	err error
)

func DatabaseInit() {
	var cfg config.ConfigDatabase

	err = cleanenv.ReadConfig(".env", &cfg)
	if err != nil {
		panic(err)
	}
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DBName)
	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	// defer db.Close()
	err = db.Ping()
	if err != nil {
		panic(err)
	}
}

func CloseDatabase() {
	db.Close()
}
