package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

// ****************
//	QUERRIES METHODS
// ****************

//params for create and update person
type PersonCreateUpdateParam struct {
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	BirthDay  time.Time `json:"birth_day"`
	InLife    bool      `json:"in_life"`
	CountryID int       `json:"country_id"`
}

//params for create and update Country
type CountryCreateUpdateParam struct {
	CountryName string `json:"name_country"`
	Continent   string `json:"continent"`
	Capital     string `json:"capital"`
}

// Methods Get person by id
// Get Person by id
// GetPersonById godoc
// @Summary Get person by id
// @Description Get person by id
// @Tags Person
// @Accept json
// @Produce json
// @Param        person_id   path      int  true  "PersonID"
// @Success 200
// @Router /api/person/{person_id} [get]
func (a *App) GetPersonById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)                        //Get data in request return map[string]string
	id, err := strconv.Atoi(vars["person_id"]) //to int
	if err != nil {
		a.responseWithError(w, http.StatusBadRequest, "Invalid Person Id")
		return
	}
	p := Person{PersonID: id}
	//Geting error return's in model_querry
	err = p.QuerryGetPersonById(a.DB)

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

// Methods Get All Persons
// GetAllPersons godoc
// @Summary Get all persons
// @Description Get all persons
// @Tags Person
// @Accept json
// @Produce json
// @Param        count   query  int   true  "Count" Format(count)
// @Param        start   query  int   true  "Start" Format(start)
// @Success 200
// @Router /api/persons [get]
func (a *App) GetAllPersons(w http.ResponseWriter, r *http.Request) {
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
	persons, err := p.QuerryGetAllPersons(a.DB, count, start)
	if err != nil {
		a.responseWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	a.responseWithJSON(w, http.StatusOK, persons)
}

// Get Person Alive return liste of personne
// GetPersonAlive godoc
// @Summary get Person Alive
// @Description get Person Alive
// @Tags Person
// @Accept json
// @Produce json
// @Success 200
// @Router /api/person/alive [get]
func (a *App) GetPersonAlive(w http.ResponseWriter, r *http.Request) {
	p := Person{}
	persons, err := p.QuerryGetPersonAlive(a.DB)
	if err != nil {
		a.responseWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	a.responseWithJSON(w, http.StatusOK, persons)
}

//get Person deaded return liste of personne
// GetPersonDeaded godoc
// @Summary get Person deaded
// @Description get Person Deaded
// @Tags Person
// @Accept json
// @Produce json
// @Success 200
// @Router /api/person/dead [get]
func (a *App) GetPersonDeaded(w http.ResponseWriter, r *http.Request) {
	p := Person{}
	persons, err := p.QuerryGetPersonDeaded(a.DB)
	if err != nil {
		a.responseWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	a.responseWithJSON(w, http.StatusOK, persons)
}

//get all country
// GetAllCountry godoc
// @Summary Get all country
// @Description Get all country
// @Tags Country
// @Accept json
// @Produce json
// @Param        count   query  int   true  "Count" Format(count)
// @Param        start   query  int   true  "Start" Format(start)
// @Success 200
// @Router /api/countries [get]
func (a *App) GetAllCountry(w http.ResponseWriter, r *http.Request) {
	countQuerry := r.URL.Query().Get("count") //Querry parse row and return corr value . Get recuparate vaule of key/ URL for client request
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
	countries, err := c.QuerryGetAllCountry(a.DB, count, start)
	if err != nil {
		a.responseWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	a.responseWithJSON(w, http.StatusOK, countries)
}

//get Country by id
// GetCountryById godoc
// @Summary Get country by id
// @Description Get country by id
// @Tags Country
// @Accept json
// @Produce json
// @Param        country_id   path      int  true  "CountryID"
// @Success 200
// @Router /api/country/{country_id} [get]
func (a *App) GetCountryById(w http.ResponseWriter, r *http.Request) {
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
	err = c.QuerryGetCountryById(a.DB)
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
// GetPersonSameCountry godoc
// @Summary Get person provide the same country
// @Description Get person provide the same country
// @Tags Person
// @Accept json
// @Produce json
// @Param        country_id   path      int  true  "CountryID"
// @Success 200
// @Router /api/person/country/{country_id} [get]
func (a *App) GetPersonSameCountry(w http.ResponseWriter, r *http.Request) {

	//get country_id in request and cast to int
	vars := mux.Vars(r)
	country_id, err := strconv.Atoi(vars["country_id"])
	if err != nil {
		a.responseWithError(w, http.StatusBadRequest, "Parameter Invalid!")
		return
	}

	//creer une instance de la structure et filter par les country_id
	p := Person{CountryID: country_id}
	persons, err := p.QuerryGetPersonSameCountry(a.DB)
	if err != nil {
		a.responseWithError(w, http.StatusNotFound, "Person Not Found!")
		return
	}
	a.responseWithJSON(w, http.StatusOK, persons)
}

//create Country
// CreateCountry godoc
// @Summary Create a new country
// @Description Create a new Country
// @Tags Country
// @Accept json
// @Produce json
// @Param        person   body    CountryCreateUpdateParam true "Country Data"
// @Success 201
// @Router /api/create/country [post]
func (a *App) CreateCountry(w http.ResponseWriter, r *http.Request) {
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
	err = c.QuerryCreateCountry(a.DB)
	if err != nil {
		a.responseWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	a.responseWithJSON(w, http.StatusCreated, c)
}

//create Person
// CreatePerson godoc
// @Summary Create a new Person
// @Description Create a new Person
// @Tags Person
// @Accept json
// @Produce json
// @Param        person   body    PersonCreateUpdateParam true "Person Data"
// @Success 201
// @Router /api/create/person [post]
func (a *App) CreatePerson(w http.ResponseWriter, r *http.Request) {
	var p Person
	decoder := json.NewDecoder(r.Body) //decode notre payload
	err := decoder.Decode(&p)          //decode la variable de notre piinter
	fmt.Println(p.BirthDay)
	if err != nil {
		fmt.Println(err)
		a.responseWithError(w, http.StatusBadRequest, "Invalid Request Payload")
		return
	}
	err = p.QuerryCreatePerson(a.DB)
	if err != nil {
		fmt.Println(err)
		a.responseWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	a.responseWithJSON(w, http.StatusCreated, p)
}

// Update country
// UpdateCountry godoc
// @Summary Update Country
// @Description update country by id
// @Tags Country
// @Accept json
// @Produce json
// @Param        country_id   path      int  true  "CountryID"
// @Param        country   body    CountryCreateUpdateParam true "Country Data"
// @Success 200
// @Router /api/update/country/{country_id} [put]
func (a *App) UpdateCountry(w http.ResponseWriter, r *http.Request) {
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

	//check if country exist
	err = c.QuerryGetCountryById(a.DB)
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

	err = c.QuerryUpdateCountry(a.DB)
	if err != nil {
		a.responseWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	a.responseWithJSON(w, http.StatusOK, c)
}

// Update person
// UpdatePerson godoc
// @Summary Update Person
// @Description update person by id
// @Tags Person
// @Accept json
// @Produce json
// @Param        person_id   path      int  true  "PersonID"
// @Param        person   body    PersonCreateUpdateParam true "Person Data"
// @Success 200
// @Router /api/update/person/{person_id} [put]
func (a *App) UpdatePerson(w http.ResponseWriter, r *http.Request) {
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

	//check if person exist
	err = p.QuerryGetPersonById(a.DB)
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
	err = p.QuerryUpdatePerson(a.DB)
	if err != nil {
		a.responseWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	a.responseWithJSON(w, http.StatusOK, p)
}

// Delete country
// DeleteCountry godoc
// @Summary Delete Country
// @Description Delete Country by id
// @Tags Country
// @Accept json
// @Produce json
// @Param        country_id   path      int  true  "CountryID"
// @Success 200
// @Router /api/delete/country/{country_id} [delete]
func (a *App) DeleteCountry(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	country_id, err := strconv.Atoi(vars["country_id"])
	if err != nil {
		a.responseWithError(w, http.StatusBadRequest, "Invalid Country Id")
		return
	}
	var c Country
	c.CountryID = country_id

	//check if country exist
	err = c.QuerryGetCountryById(a.DB)
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

	err = c.QuerryDeleteCountry(a.DB)
	if err != nil {
		a.responseWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	a.responseWithJSON(w, http.StatusOK, delete_message(country_id))
}

// Delete Person
// DeletePerson godoc
// @Summary Delete Person
// @Description Delete Person by id
// @Tags Person
// @Accept json
// @Produce json
// @Param        person_id   path      int  true  "PersonID"
// @Success 200
// @Router /api/delete/person/{person_id} [delete]
func (a *App) DeletePerson(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	person_id, err := strconv.Atoi(vars["person_id"])

	if err != nil {
		a.responseWithError(w, http.StatusBadRequest, "Invalid Person Id")
	}

	p := Person{PersonID: person_id}

	//check if person exist
	err = p.QuerryGetPersonById(a.DB)
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

	err = p.QuerryDeletePerson(a.DB)
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
func (a *App) responseWithJSON(w http.ResponseWriter, statusCode int, payload interface{}) {
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
