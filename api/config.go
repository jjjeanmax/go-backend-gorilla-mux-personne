package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type DataBase struct {
	Database Db `json:"DATABASE"`
}
type Db struct {
	Username string `json:"DB_USERNAME"`
	Name     string `json:"DB_NAME"`
	Password string `json:"DB_PASSWORD"`
	Port     string `json:"DB_PORT"`
	Host     string `json:"DB_HOST"`
}

//get db config
func ConfigsDb() []string {
	jsonfile, err := os.Open("config.json")
	if err != nil {
		fmt.Println("config.json Not Found")
	}

	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonfile.Close()

	// read our opened xmlFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonfile)

	// we initialize our db array
	var db DataBase

	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'db' which we defined above
	json.Unmarshal(byteValue, &db)

	var settings []string
	settings = append(settings, db.Database.Username, db.Database.Password, db.Database.Name, db.Database.Port, db.Database.Host)
	return settings
}
