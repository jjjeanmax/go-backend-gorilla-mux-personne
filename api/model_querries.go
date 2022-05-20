package main

import (
	"database/sql"
)

type Country struct {
	CountryID int    `json:"country_id"`
	Name      string `json:"name"`
	Continent string `json:"Continent"`
	Capital   string `json:"capital"`
}

type Person struct {
	PersonID  int    `json:"person_id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	BirthDay  string `json:"birth_day"`
	InLife    bool   `json:"in_life"`
	CountryID int    `json:"country_id"`
	Registry  string `json:"registry"`
}

// *************************************************************
// all this Method are using in the handler (doing a sql querry)
// *************************************************************

//method to get person by id return err
func (p *Person) getPersonById(db *sql.DB) error {
	return db.QueryRow(
		"SELECT first_name,last_name,birth_day,in_life,country_id,registry from person WHERE person_id=$1",
		p.PersonID,
	).Scan(&p.FirstName, &p.LastName, &p.BirthDay, &p.InLife, &p.CountryID, &p.Registry)
	// scan method only works on method that return row(s)
}

//method to get all person return err () ou list de person
func (p *Person) getAllPersons(db *sql.DB, count, start int) ([]Person, error) {
	rows, err := db.Query(
		"SELECT person_id,first_name,last_name,birth_day,in_life,country_id,registry from person LIMIT $1 OFFSET $2",
		count, start,
	)
	if err != nil {
		return nil, err
	}

	// Close closes the database and prevents new queries from starting.
	// Close then waits for all queries that have started processing on
	// the server to finish.
	defer rows.Close()

	persons := []Person{} //creer une instance de Person

	for rows.Next() { //bool
		var p Person

		// can copies the columns from the matched row into the values pointed at by dest
		err := rows.Scan(&p.PersonID, &p.FirstName, &p.LastName, &p.BirthDay, &p.InLife, &p.CountryID, &p.Registry)
		if err != nil {
			return nil, err
		}
		persons = append(persons, p)
	}
	return persons, nil
}

//methods get person in live
func (p *Person) getPersonAlive(db *sql.DB) ([]Person, error) {
	rows, err := db.Query(
		"SELECT person_id,first_name,last_name,birth_day,in_life,country_id,registry from person WHERE in_life=TRUE",
	)
	if err != nil {
		return nil, err
	}
	//Close closes the database and prevents new queries from starting.
	defer rows.Close()

	persons := []Person{}

	for rows.Next() { //bool
		var p Person

		// can copies the columns from the matched row into the values pointed at by dest
		err := rows.Scan(&p.PersonID, &p.FirstName, &p.LastName, &p.BirthDay, &p.InLife, &p.CountryID, &p.Registry)
		if err != nil {
			return nil, err
		}
		persons = append(persons, p)
	}
	return persons, nil

}

//methos to get person deaded
func (p *Person) getPersonDeaded(db *sql.DB) ([]Person, error) {
	rows, err := db.Query(
		"SELECT person_id,first_name,last_name,birth_day,in_life,country_id,registry from person WHERE in_life=TRUE",
	)
	if err != nil {
		return nil, err
	}
	//Close closes the database and prevents new queries from starting.
	defer rows.Close()

	persons := []Person{}
	for rows.Next() {
		var p Person
		err := rows.Scan(&p.PersonID, &p.FirstName, &p.LastName, &p.BirthDay, &p.InLife, &p.CountryID, &p.Registry)
		if err != nil {
			return nil, err
		}
		persons = append(persons, p)
	}
	return persons, err
}
