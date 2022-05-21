package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"sync/atomic"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

type App struct {
	Router      *mux.Router
	DB          *sql.DB
	Logger      *log.Logger
	shutdownReq chan bool
	reqCount    uint32
}

// method to initialise our db
func (a *App) InitializeDb(user, password, dbname, port, host string) {

	a.Logger = log.New(os.Stdout, "", log.LstdFlags)

	//create parameters to connection
	connection_to_db := fmt.Sprintf("user=%s password=%s dbname=%s port=%s host=%s sslmode=disable", user, password, dbname, port, host)

	var err error
	//connection to postgresql
	a.DB, err = sql.Open("postgres", connection_to_db)
	if err != nil {
		fmt.Println("Cannot connect to the database right now")
		a.Logger.Fatal(err)
	}

	//initialise router
	a.Router = mux.NewRouter()
	a.initialiseRouters()

}

// methods to initialize all the routes (all methos in handlers)
func (a *App) initialiseRouters() {
	a.Router.HandleFunc("/shutdown", a.ShutdownHandler)
	a.Router.HandleFunc("/api/person/{person_id:[0-9]+}", a.getPersonById).Methods("GET")
	a.Router.HandleFunc("/api/persons", a.getAllPersons).Methods("GET")
	a.Router.HandleFunc("/api/person/alive", a.getPersonAlive).Methods("GET")
	a.Router.HandleFunc("/api/person/dead", a.getPersonDeaded).Methods("GET")
	a.Router.HandleFunc("/api/countries", a.getAllCountry).Methods("GET")
	a.Router.HandleFunc("/api/country/{country_id:[0-9]+}", a.getCountryById).Methods("GET")
	a.Router.HandleFunc("/api/person/country/{country_id:[0-9]+}", a.getPersonSameCountry).Methods("GET")
	a.Router.HandleFunc("/api/create/country", a.createCountry).Methods("POST")
	a.Router.HandleFunc("/api/create/person", a.createPerson).Methods("POST")
	a.Router.HandleFunc("/api/update/country/{country_id:[0-9]+}", a.updateCountry).Methods("PUT")
	a.Router.HandleFunc("/api/update/person/{person_id:[0-9]+}", a.updatePerson).Methods("PUT")
	a.Router.HandleFunc("/api/delete/country/{country_id:[0-9]+}", a.deleteCountry).Methods("DELETE")
	a.Router.HandleFunc("/api/delete/person/{person_id:[0-9]+}", a.deletePerson).Methods("DELETE")
}

//method run to start our app (une addresse et un port addr est le port)
func (a *App) Run(addr string) {
	done := make(chan bool)

	loggedRouter := a.createLoggingRouter(a.Logger.Writer())
	a.Logger.Fatal(http.ListenAndServe(addr, loggedRouter))

	done <- true
}

// method to shutdown server
func (a *App) ShutdownHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Shutdown server"))

	//Do nothing if shutdown request already issued
	//if a.reqCount == 0 then set to 1, return true otherwise false
	if !atomic.CompareAndSwapUint32(&a.reqCount, 0, 1) {
		log.Printf("Shutdown through API call in progress...")
		return
	}

	go func() {
		a.shutdownReq <- true
	}()
}
