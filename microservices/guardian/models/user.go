package models

import (
	"errors"
	"fmt"
	"guardian/api/db"
	"guardian/api/utils"
	"time"
)

type User struct {
	ID               int64     `json:"id"`
	DisplayName      string    `json:"displayName" validate:"omitempty,required,min=3,max=32"`
	Email            string    `binding:"required" json:"email" validate:"required,email"`
	Password         string    `binding:"required" json:"password" validate:"required,min=5,max=32"`
	Enabled          int64     `json:"enabled"`
	Roles            string    `json:"roles"`
	LastTimeLoggedIn time.Time `json:"lastTimeLoggedIn"`
	LogInCount       int64     `json:"logInCount"`
	Updated          time.Time `json:"updated"`
	Created          time.Time `json:"created"`
}

type DisplayName struct {
	DisplayName string `json:"displayName" validate:"required,min=3,max=32"`
}

type Email struct {
	Email string `binding:"required" json:"email" validate:"required,email"`
}

type Password struct {
	Password string `binding:"required" json:"password" validate:"required,min=5,max=32"`
}

func (u *User) Create() error {
	query := "INSERT INTO user (display_name, email, password) values (?, ?, ?)"
	// prepare the database query for execution
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		fmt.Println("MC1", err)
		return err
	}

	fmt.Println("db prep")

	// set up the database close defer
	defer stmt.Close()
	//hash password
	hashedPassword, err := utils.HashPassword(u.Password)

	if err != nil {
		return err
	}
	// update the user instance with the hashed password
	u.Password = hashedPassword
	//execute the statement against the database
	res, err := stmt.Exec(u.DisplayName, u.Email, u.Password)

	if err != nil {
		fmt.Println("MC2", err)
		return err
	}

	fmt.Println("db exec")

	// update model with last inserted id
	userId, err := res.LastInsertId()
	u.ID = userId

	return err
}

func (u *User) ValidateCredentials() error {
	query := "SELECT id, password, roles, enabled FROM user WHERE email = ?"
	// execute query
	row := db.DB.QueryRow(query, u.Email)
	// retrieve the password
	var retrievedPassword string
	err := row.Scan(&u.ID, &retrievedPassword, &u.Roles, &u.Enabled)

	if err != nil {
		fmt.Println(err.Error())
		return err
		// if err.Error() == "sql: no rows in result set" { // checking to see if user exists
		// 	return err
		// } else {
		// 	return errors.New("credentials invalid")
		// }
	}
	// check password
	isPasswordValid := utils.CheckPasswordHash(u.Password, retrievedPassword)
	if !isPasswordValid {
		return errors.New("credentials invalid")
	}
	// credentials passed
	return nil
}

func GetUserById(id int64) (*User, error) {
	var user User
	// query the database for the specific user
	query := "SELECT * FROM user WHERE id = ?"
	row := db.DB.QueryRow(query, id)
	// extract values from row
	err := row.Scan(&user.ID, &user.DisplayName, &user.Email, &user.Password, &user.Enabled, &user.Roles, &user.LastTimeLoggedIn, &user.LogInCount, &user.Updated, &user.Created)

	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	// return user
	return &user, nil
}

func (dn *DisplayName) UpdateDisplayNameById(userId int64) error {
	query := "UPDATE user SET display_name = ? WHERE id = ?"
	// prepare the database query for execution
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}
	// defer the database connection close
	defer stmt.Close()
	// execute the statement against the database
	_, err = stmt.Exec(dn.DisplayName, userId)

	// update model with last inserted id
	// userId, err := res.LastInsertId()
	// dn.ID = userId

	return err
}

func (e *Email) UpdateEmailById(userId int64) error {
	query := "UPDATE user SET email = ? WHERE id = ?"
	// prepare the database query for execution
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}
	// defer the database connection close
	defer stmt.Close()
	// execute the statement against the database
	_, err = stmt.Exec(e.Email, userId)

	return err
}

func (p *Password) UpdatePasswordById(userId int64) error {
	query := "UPDATE user SET password = ? WHERE id = ?"
	// prepare the database query for execution
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}
	// defer the database connection close
	defer stmt.Close()
	//hash password
	hashedPassword, err := utils.HashPassword(p.Password)

	if err != nil {
		return err
	}
	// update the user instance with the hashed password
	p.Password = hashedPassword
	// execute the statement against the database
	_, err = stmt.Exec(p.Password, userId)

	return err
}
