package main

import _ "gorilla-mux-person/docs"

// @title Golang API Person Swagger
// @version 1.0
// @description This is an auto-generated API Docs Person.
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email your@mail.com
// @in header
// @BasePath /
func main() {
	//get our config settings value
	APP_DB_USERNAME := ConfigsDb()[0]
	APP_DB_PASSWORD := ConfigsDb()[1]
	APP_DB_NAME := ConfigsDb()[2]
	DB_PORT := ConfigsDb()[3]
	DB_HOST := ConfigsDb()[4]

	//create an instance App
	a := App{}
	a.InitializeDb(
		APP_DB_USERNAME,
		APP_DB_PASSWORD,
		APP_DB_NAME,
		DB_PORT,
		DB_HOST,
	)

	//run
	a.Run("localhost:8888")

}
