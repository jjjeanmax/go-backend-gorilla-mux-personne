package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// ****************
//	QUERRIES METHODS
// ****************

func (a *App) getPersonById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)                        //get data in request return map[string]string
	id, err := strconv.Atoi(vars["person_id"]) //to int

	if err != nil {
		responseWithError(w, http.StatusBadRequest, "Invalid Person Id")
		return
	}
	p := Person{PersonID: id}
	//geting error return's in model_querry
	err = p.getPersonById(a.DB)

	if err != nil {
		switch err {
		case sql.ErrNoRows:
			responseWithError(w, http.StatusNotFound, "Person Not Found")
			return
		default:
			responseWithError(w, http.StatusBadGateway, err.Error())
			return
		}
	}
	responseWithJSON(w, http.StatusOK, p)
}

// Methods get all Person
func (a *App) getAllPersons(w http.ResponseWriter, r *http.Request) {
	countQuerry := r.URL.Query().Get("count")
	startQuerry := r.URL.Query().Get("start")

	count, err1 := strconv.Atoi(countQuerry)
	start, err2 := strconv.Atoi(startQuerry)

	if err1 != nil || err2 != nil {
		fmt.Println(err1, err2)
		responseWithError(w, http.StatusBadRequest, "Parameters Invalid !")
		return
	}
	p := Person{}

	//get list of person or error return in model_querries
	persons, err := p.getAllPersons(a.DB, count, start)
	if err != nil {
		responseWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	responseWithJSON(w, http.StatusOK, persons)
}

//get Person Alive return liste of personne
func (a *App) getPersonAlive(w http.ResponseWriter, r *http.Request) {
	p := Person{}
	persons, err := p.getPersonAlive(a.DB)
	if err != nil {
		responseWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	responseWithJSON(w, http.StatusOK, persons)
}

//get Person Alive return liste of personne
func (a *App) getPersonDeaded(w http.ResponseWriter, r *http.Request) {
	p := Person{}
	persons, err := p.getPersonDeaded(a.DB)
	if err != nil {
		responseWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	responseWithJSON(w, http.StatusOK, persons)
}

// ****************
//	HELPER METHODS
// ****************

//if error return {statuscode:message} ex {200:"ok"}
func responseWithError(w http.ResponseWriter, statusCode int, message string) {
	var errorMessage = map[string]string{"error": message}
	responseWithJSON(w, statusCode, errorMessage)
}

// return a json format response data
func responseWithJSON(w http.ResponseWriter, statusCode int, payload interface{}) {
	//encode response for the client marshall to return json encoding
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json") //to json
	w.WriteHeader(statusCode)
	w.Write(response)
}
