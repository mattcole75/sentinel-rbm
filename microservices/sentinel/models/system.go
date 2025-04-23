package models

import (
	"sentinel/api/db"
	"time"
)

type System struct {
	ID          int64     `json:"id"`
	Name        string    `json:"name" validate:"required,min=1,max=64"`
	Description string    `json:"description" validate:"required,min=1,max=512"`
	Updated     time.Time `json:"updated"`
	Created     time.Time `json:"created"`
}

// create new system
func (s *System) Create() error {
	query := "insert into functional_system (name, description) values (?, ?)"
	// prepare the query
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	// defer db connection close to completion
	defer stmt.Close()
	// exec statement
	res, err := stmt.Exec(s.Name, s.Description)
	if err != nil {
		return err
	}
	// update model with last inserted id
	id, err := res.LastInsertId()
	s.ID = id

	return err
}

// return all systems that are in use
func GetSystems() ([]System, error) {
	query := "select * from functional_system"
	// query the database
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	// defer db connection close to completion
	defer rows.Close()
	// process rows
	var systems []System
	for rows.Next() {
		var system System
		err := rows.Scan(&system.ID, &system.Name, &system.Description, &system.Updated, &system.Created)
		if err != nil {
			return nil, err
		}
		systems = append(systems, system)
	}

	return systems, nil
}

// return system by id
func GetSystemById(id int64) (*System, error) {
	query := "select * from functional_system where id = ?"
	// query the database
	row := db.DB.QueryRow(query, id)
	//process row
	var system System
	err := row.Scan(&system.ID, &system.Name, &system.Description, &system.Updated, &system.Created)
	if err != nil {
		return nil, err
	}

	return &system, nil
}

// update system
func (s *System) Update() error {
	query := "update functional_system set name = ?, description = ? where id = ?"
	// prepare the query for execution
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	// defer db connection close to completion
	defer stmt.Close()
	// execute the query
	_, err = stmt.Exec(s.Name, s.Description, s.ID)

	return err
}

// delete system
func (s *System) Delete() error {
	query := "delete from functional_system where id = ?"
	// prepare query for execution
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	// defer db connection close to completion
	defer stmt.Close()
	//execute the query
	_, err = stmt.Exec(s.ID)

	return err
}
