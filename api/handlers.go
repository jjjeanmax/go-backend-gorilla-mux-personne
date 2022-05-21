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
		a.responseWithError(w, http.StatusBadRequest, "Invalid Person Id")
    return
	}
	p := Person{PersonID: id}
	//geting error return's in model_querry
	err = p.querryGetPersonById(a.DB)

	if err != nil {
		switch err {
		case sql.ErrNoRows:
			a.responseWithError(w, http.StatusNotFound, "Person Not Found")
			return
		default:
			a.responseWithError(w, http.StatusBadGateway, err.Error())
			return
		}
	}
	a.responseWithJSON(w, http.StatusOK, p)
}

// Methods get all Person
func (a *App) getAllPersons(w http.ResponseWriter, r *http.Request) {
	countQuerry := r.URL.Query().Get("count")
	startQuerry := r.URL.Query().Get("start")

	count, err1 := strconv.Atoi(countQuerry)
	start, err2 := strconv.Atoi(startQuerry)

	if err1 != nil || err2 != nil {
		fmt.Println(err1, err2)
		a.responseWithError(w, http.StatusBadRequest, "Parameters Invalid !")
		return
	}
	p := Person{}

	//get list of person or error return in model_querries
	persons, err := p.querryGetAllPersons(a.DB, count, start)
	if err != nil {
		a.responseWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	a.responseWithJSON(w, http.StatusOK, persons)
}

//get Person Alive return liste of personne
func (a *App) getPersonAlive(w http.ResponseWriter, r *http.Request) {
	p := Person{}
	persons, err := p.querryGetPersonAlive(a.DB)
	if err != nil {
		a.responseWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	a.responseWithJSON(w, http.StatusOK, persons)
}

//get Person Alive return liste of personne
func (a *App) getPersonDeaded(w http.ResponseWriter, r *http.Request) {
	p := Person{}
	persons, err := p.querryGetPersonDeaded(a.DB)
	if err != nil {
		a.responseWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	a.responseWithJSON(w, http.StatusOK, persons)
}

//get all country
func (a *App) getAllCountry(w http.ResponseWriter, r *http.Request) {
	countQuerry := r.URL.Query().Get("count") //Querry parse row and return corr value . get recuparate vaule of key/ URL for client request
	startQuerry := r.URL.Query().Get("start")

	//convert to int
	count, err1 := strconv.Atoi(countQuerry)
	start, err2 := strconv.Atoi(startQuerry)

	//check err1 and err2 (got no int parameters)
	if err1 != nil || err2 != nil {
		fmt.Println(err1, err2)
		a.responseWithError(w, http.StatusBadRequest, "Parameters Invalid !")
		return
	}
	c := Country{}
	//pour recupere la valeur de la metodes je creer l'instance de la structure pointee par la methode
	//et cette instance est affectee a la methode
	countries, err := c.querryGetAllCountry(a.DB, count, start)
	if err != nil {
		a.responseWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	a.responseWithJSON(w, http.StatusOK, countries)
}

// get country by id
func (a *App) getCountryById(w http.ResponseWriter, r *http.Request) {
	// vars retourne la variable contenue dans la request (map[string]string)
	vars := mux.Vars(r)

	//convert to int
	country_id, err := strconv.Atoi(vars["country_id"])
	if err != nil {
		a.responseWithError(w, http.StatusBadRequest, "Parameter Invalid!")
		return
	}

	//get country by id
	c := Country{CountryID: country_id}

	//return error(id not found or badgetway)
	err = c.querryGetCountryById(a.DB)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			a.responseWithError(w, http.StatusNotFound, "Country Not Found")
			return
		default:
			a.responseWithError(w, http.StatusBadGateway, err.Error())
			return
		}
	}
	a.responseWithJSON(w, http.StatusOK, c)

}

//get person provide the same country
func (a *App) getPersonSameCountry(w http.ResponseWriter, r *http.Request) {

	//get country_id in request and cast to int
	vars := mux.Vars(r)
	country_id, err := strconv.Atoi(vars["country_id"])
	if err != nil {
		a.responseWithError(w, http.StatusBadRequest, "Parameter Invalid!")
		return
	}

	//creer une instance de la structure et filter par les country_id
	p := Person{CountryID: country_id}
	persons, err := p.querryGetPersonSameCountry(a.DB)
	if err != nil {
		a.responseWithError(w, http.StatusNotFound, "Person Not Found!")
		return
	}
	a.responseWithJSON(w, http.StatusOK, persons)
}

//create Country
func (a *App) createCountry(w http.ResponseWriter, r *http.Request) {
	var c Country
	//decod payload in body
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&c)

	//check error in payload decoder
	if err != nil {
		fmt.Println(err)
		a.responseWithError(w, http.StatusBadGateway, "Invalid request payload")
		return
	}

	//get error send in model_querries
	err = c.querryCreateCountry(a.DB)
	if err != nil {
		a.responseWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	a.responseWithJSON(w, http.StatusCreated, c)
}

//create person
func (a *App) createPerson(w http.ResponseWriter, r *http.Request) {
	var p Person
	decoder := json.NewDecoder(r.Body) //decode notre payload
	err := decoder.Decode(&p)          //decode la variable de notre piinter
	fmt.Println(*&p.BirthDay)
	if err != nil {
		fmt.Println(err)
		a.responseWithError(w, http.StatusBadRequest, "Invalid Request Payload")
		return
	}
	err = p.querryCreatePerson(a.DB)
	if err != nil {
		fmt.Println(err)
		a.responseWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	a.responseWithJSON(w, http.StatusCreated, p)
}

//update country
func (a *App) updateCountry(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	country_id, err := strconv.Atoi(vars["country_id"])
	if err != nil {
		a.responseWithError(w, http.StatusBadRequest, "Invalid Country Id")
		return
	}
	var c Country
	c.CountryID = country_id
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&c)

	if err != nil {
		a.responseWithError(w, http.StatusBadRequest, "Invalid Request Payload")
		return
	}

	err = c.querryUpdateCountry(a.DB)
	if err != nil {
		a.responseWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	a.responseWithJSON(w, http.StatusOK, c)
}

//update person
func (a *App) updatePerson(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	person_id, err := strconv.Atoi(vars["person_id"])
	if err != nil {
		a.responseWithError(w, http.StatusBadRequest, "Invalid Person Id")
		return
	}
	var p Person
	p.PersonID = person_id

	//check err in body
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&p)

	if err != nil {
		a.responseWithError(w, http.StatusBadRequest, "Invalid Request Payload")
		return
	}
	err = p.querryUpdatePerson(a.DB)
	if err != nil {
		a.responseWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	a.responseWithJSON(w, http.StatusOK, p)
}

//delete country
func (a *App) deleteCountry(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	country_id, err := strconv.Atoi(vars["country_id"])
	if err != nil {
		a.responseWithError(w, http.StatusBadRequest, "Invalid Country Id")
		return
	}
	var c Country
	c.CountryID = country_id
	err = c.querryDeleteCountry(a.DB)
	if err != nil {
		a.responseWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	a.responseWithJSON(w, http.StatusOK, delete_message(country_id))
}

//delete Person
func (a *App) deletePerson(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	person_id, err := strconv.Atoi(vars["person_id"])

	if err != nil {
		a.responseWithError(w, http.StatusBadRequest, "Invalid Person Id")
	}

	p := Person{PersonID: person_id}
	err = p.querryDeletePerson(a.DB)
	if err != nil {
		a.responseWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	a.responseWithJSON(w, http.StatusOK, delete_message(person_id))

}

// ****************
//	HELPERS METHODS
// ****************

//if error return {statuscode:message} ex {200:"ok"}
func (a *App) responseWithError(w http.ResponseWriter, statusCode int, message string) {
	var errorMessage = map[string]string{"error": message}

	a.responseWithJSON(w, statusCode, errorMessage)

	a.Logger.Printf("App error: statusCode %d, message %s", statusCode, message)
}

// return a json format response data
func responseWithJSON(w http.ResponseWriter, statusCode int, payload interface{}) {
	//encode response for the client marshall to return json encoding
	response, _ := json.Marshal(payload)
	
  w.Header().Set("Content-Type", "application/json") //to json
	w.WriteHeader(statusCode)
	w.Write(response)
}

// return message after delete
func delete_message(statusCode int) map[string]string {
	message := "id " + strconv.Itoa(statusCode) + " deleted!"
	return map[string]string{"message": message}
}
