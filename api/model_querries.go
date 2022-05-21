package main

import (
	"database/sql"
	"time"
)

/*"TODO":si on supprime le pays on supprime les habitants du pays (delete cascade)*/

type Country struct {
	CountryID   int    `json:"country_id"`
	CountryName string `json:"name_country"`
	Continent   string `json:"continent"`
	Capital     string `json:"capital"`
}

type Person struct {
	PersonID  int       `json:"person_id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	BirthDay  time.Time `json:"birth_day"`
	InLife    bool      `json:"in_life"`
	CountryID int       `json:"country_id"`
	Registry  time.Time `json:"registry"`
}

const (
	YYYYMMDD       = "2006-01-02"
	DDMMYYYYhhmmss = "2006-01-02 15:04:05"
)

var now = time.Now().UTC()

// *************************************************************
// all this Method are using in the handler (doing a sql querry)
// *************************************************************

//method to get person by id return err
func (p *Person) querryGetPersonById(db *sql.DB) error {
	return db.QueryRow(
		"SELECT first_name,last_name,birth_day,in_life,country_id,registry from person WHERE person_id=$1",
		p.PersonID,
	).Scan(&p.FirstName, &p.LastName, &p.BirthDay, &p.InLife, &p.CountryID, &p.Registry)
	// scan method only works on method that return row(s)
}

//method to get all person return err () ou list de person
func (p *Person) querryGetAllPersons(db *sql.DB, count, start int) ([]Person, error) {
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
func (p *Person) querryGetPersonAlive(db *sql.DB) ([]Person, error) {
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
func (p *Person) querryGetPersonDeaded(db *sql.DB) ([]Person, error) {
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

//get all country
func (c *Country) querryGetAllCountry(db *sql.DB, count, start int) ([]Country, error) {
	rows, err := db.Query(
		"SELECT country_id,name_country,continent,capital from country LIMIT $1 OFFSET $2", count, start,
	)
	if err != nil {
		return nil, err
	}
	//Close closes the database and prevents new queries from starting.
	defer rows.Close()

	countries := []Country{}

	for rows.Next() { //rows.new() type bool (true/false)
		var c Country
		err := rows.Scan(&c.CountryID, &c.CountryName, &c.Capital, &c.Continent) //scan to copy values in row pointed return err

		if err != nil {
			return nil, err
		}
		countries = append(countries, c)
	}
	return countries, err
}

//get country by country_id
func (c *Country) querryGetCountryById(db *sql.DB) error {
	//QuerrRow return row.sql // i use scan to copy values of pointer and return err
	return db.QueryRow(
		"SELECT country_id,name_country,continent,capital from country WHERE country_id=$1", c.CountryID,
	).Scan(&c.CountryID, &c.CountryName, &c.Continent, &c.Capital)
}

//get persons in the same country
func (p *Person) querryGetPersonSameCountry(db *sql.DB) ([]Person, error) {
	rows, err := db.Query(
		"SELECT person_id,first_name,last_name,birth_day,in_life,country_id,registry from person WHERE country_id=$1", p.CountryID,
	)
	if err != nil {
		return nil, err
	}
	//Close closes the database and prevents new queries from starting.
	defer rows.Close()

	persons := []Person{} //pour append les valeurs des rows

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

//create country
func (p *Person) querryCreatePerson(db *sql.DB) error {
	return db.QueryRow(
		"INSERT INTO person(person_id,first_name,last_name,birth_day,in_life,country_id,registry) VALUES(nextval('persone_sequence'),$1,$2,$3,$4,$5,$6) RETURNING person_id,first_name,last_name,birth_day,in_life,country_id,registry",
		p.FirstName,
		p.LastName,
		p.BirthDay.Format(YYYYMMDD),
		p.InLife,
		p.CountryID,
		now.Format(DDMMYYYYhhmmss),
	).Scan(
		&p.PersonID,
		&p.FirstName,
		&p.LastName,
		&p.BirthDay,
		&p.InLife,
		&p.CountryID,
		&p.Registry,
	)
}

//create country
func (c *Country) querryCreateCountry(db *sql.DB) error {
	return db.QueryRow(
		"INSERT INTO country(country_id,name_country,continent,capital) VALUES(nextval('country_sequence'),$1,$2,$3) RETURNING country_id,name_country,continent,capital",
		c.CountryName,
		c.Continent,
		c.Capital,
	).Scan(
		&c.CountryID,
		&c.CountryName,
		&c.Continent,
		&c.Capital,
	)
}

//update country
func (c *Country) querryUpdateCountry(db *sql.DB) error {
	//db.exc return sql resul or err
	_, err := db.Exec(
		"UPDATE country SET name_country=$1,continent=$2,capital=$3 WHERE country_id=$4",
		c.CountryName, c.Continent, c.Capital, c.CountryID,
	)
	return err
}

//update person
func (p *Person) querryUpdatePerson(db *sql.DB) error {
	_, err := db.Exec(
		"UPDATE person SET first_name=$1,last_name=$2,birth_day=$3,in_life=$4,country_id=$5,registry=$6 WHERE person_id=$7",
		p.FirstName, p.LastName, p.BirthDay.Format(YYYYMMDD), p.InLife, p.CountryID, now.Format(DDMMYYYYhhmmss), p.PersonID,
	)
	return err
}

//delete country
func (c *Country) querryDeleteCountry(db *sql.DB) error {
	_, err := db.Exec(
		"DELETE FROM country WHERE country_id=$1", c.CountryID,
	)
	return err
}

//delete person
func (p *Person) querryDeletePerson(db *sql.DB) error {
	_, err := db.Exec(
		"DELETE FROM person WHERE person_id=$1", p.PersonID,
	)
	return err
}
