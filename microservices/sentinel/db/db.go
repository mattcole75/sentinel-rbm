package db

import (
	"database/sql"
	"fmt"
	"sentinel/api/config"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

// global pointer to the instance
var DB *sql.DB

// initialise the database connection
func InitialiseDB() {
	var err error
	DB, err = sql.Open("mysql", config.Config["dbConnection"])
	if err != nil {
		fmt.Println(err.Error())
		panic("could not initialise database")
	}

	// set db connections
	DB.SetConnMaxLifetime(time.Minute * 3)
	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	// test the db connection credentials provided in the config file
	err = DB.Ping()
	if err != nil {
		panic(err.Error())
	}
}
