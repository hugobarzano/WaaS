package controller

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"waas/models"
)

// Index GET /
func (c *Controller) IndexV1(w http.ResponseWriter, r *http.Request) {
	businessObjects := c.Storer.GetBusinessObjects() // list of all businessObjects
	data, _ := json.Marshal(businessObjects)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
	return
}

// PushObjectV1 POST /
func (c *Controller) PushObjectV1(w http.ResponseWriter, r *http.Request) {
	var businessObject models.BusinessObject
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576)) // read the body of the request

	log.Println(body)

	if err != nil {
		log.Fatalln("Error PushObject", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := r.Body.Close(); err != nil {
		log.Fatalln("Error PushObject", err)
	}

	if err := json.Unmarshal(body, &businessObject); err != nil { // unmarshall body contents as a type Candidate
		w.WriteHeader(422) // unprocessable entity
		log.Println(err)
		if err := json.NewEncoder(w).Encode(err); err != nil {
			log.Fatalln("Error PushObject unmarshalling data", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}

	log.Println(businessObject)
	success := c.Storer.PushObject(businessObject) // adds the businessObject to the DB
	if !success {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	return
}

