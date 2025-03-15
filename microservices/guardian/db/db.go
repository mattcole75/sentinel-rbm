package db

import (
	"database/sql"
	"fmt"
	"guardian/api/config"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

// global pointer to the instance of the database
var DB *sql.DB

// initialise the database
func InitialiseDB() {

	// open a database connection
	var err error
	DB, err = sql.Open("mysql", config.Config["dbConnection"])
	if err != nil {
		fmt.Println(err.Error())
		panic("could not initialise database")
	}

	// set db connections
	DB.SetConnMaxLifetime(time.Minute * 3) // close db connections before driver closes them
	DB.SetMaxOpenConns(10)                 // a maximum number of open connections at any one time
	DB.SetMaxIdleConns(5)                  // a minimum number of idle connections while the system is inactive

	// test the db connection credentials provided in the config file
	err = DB.Ping()
	if err != nil {
		panic(err.Error())
	}
}
