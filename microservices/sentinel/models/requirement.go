package models

import (
	"fmt"
	"sentinel/api/db"
	"time"
)

type Requirement struct {
	ID              int64     `json:"id"`
	Title           string    `json:"title" validate:"required,min=3,max=64"`
	Statement       string    `json:"statement" validate:"required,min=10,max=512"`
	Reference       string    `json:"reference" validate:"required,min=3,max=64"`
	ReferenceSource string    `json:"referenceSource" validate:"required,min=3,max=16"`
	Updated         time.Time `json:"updated"`
	Created         time.Time `json:"created"`
}

func (r *Requirement) Create() error {
	query := "INSERT INTO functional_requirement (title, statement, reference, reference_source) values (?, ?, ?, ?)"
	// prepare the query for execution
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	// set up the database connection close deferral
	defer stmt.Close()
	res, err := stmt.Exec(r.Title, r.Statement, r.Reference, r.ReferenceSource)
	if err != nil {
		return err
	}
	// update model with the last inserted id
	reqId, err := res.LastInsertId()
	r.ID = reqId

	return err
}

func GetRequirements() ([]Requirement, error) {
	query := "select * from functional_requirement"

	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var requirements []Requirement

	for rows.Next() {
		var requirement Requirement
		err := rows.Scan(&requirement.ID, &requirement.Title, &requirement.Statement, &requirement.Reference, &requirement.ReferenceSource, &requirement.Updated, &requirement.Created)
		if err != nil {
			return nil, err
		}

		requirements = append(requirements, requirement)
	}

	return requirements, nil
}

func GetRequirementById(id int64) (*Requirement, error) {
	query := "select * from functional_requirement where id = ?"
	row := db.DB.QueryRow(query, id)

	var requirement Requirement

	err := row.Scan(&requirement.ID, &requirement.Title, &requirement.Statement, &requirement.Reference, &requirement.ReferenceSource, &requirement.Updated, &requirement.Created)
	if err != nil {
		return nil, err
	}

	return &requirement, nil
}

func (r *Requirement) Update() error {
	query := "update functional_requirement set title = ?, statement = ?, reference = ?, reference_source = ? where id = ?"
	// prepare the query for execution
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	// defer db connection closure until after execution
	defer stmt.Close()
	// execute the sql statement
	_, err = stmt.Exec(r.Title, r.Statement, r.Reference, r.ReferenceSource, r.ID)

	return err
}

func (r *Requirement) Delete() error {
	query := "delete from functional_requirement where id = ?"

	// prepare the query for execution
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	// defer db connection closure until after execution
	defer stmt.Close()
	// execute the sql statement
	fmt.Println(r)
	_, err = stmt.Exec(r.ID)

	return err
}
