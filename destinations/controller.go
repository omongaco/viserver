package destinations

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

//Controller is the thing
type Controller struct {
	Repository Repository
}

//Index GET
func (c *Controller) Index(w http.ResponseWriter, r *http.Request) {
	destinations := c.Repository.GetDestinations()
	log.Println(destinations)
	data, _ := json.Marshal(destinations)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	w.Write(data)

	return
}

//AddDestination POST
func (c *Controller) AddDestination(w http.ResponseWriter, r *http.Request) {
	var destination Destination
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))

	if err != nil {
		log.Fatalln("Error AddDestination", err)
		w.WriteHeader(http.StatusInternalServerError)

		return
	}

	if err := r.Body.Close(); err != nil {
		log.Fatalln("Error AddDestination", err)
	}

	if err := json.Unmarshal(body, &destination); err != nil {
		w.WriteHeader(422)
		if err := json.NewEncoder(w).Encode(err); err != nil {
			log.Fatalln("Error AddDestination unmarshalling data", err)
			w.WriteHeader(http.StatusInternalServerError)
		}
	}

	success := c.Repository.AddDestination(destination)
	if !success {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)

	return
}

//UpdateDestination PUT
func (c *Controller) UpdateDestination(w http.ResponseWriter, r *http.Request) {
	var destination Destination
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))

	if err != nil {
		log.Fatalln("Error UpdateDestination", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := r.Body.Close(); err != nil {
		log.Fatalln("Error UpdateDestination", err)
	}

	if err := json.Unmarshal(body, &destination); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422)
		if err := json.NewEncoder(w).Encode(err); err != nil {
			log.Fatalln("Error UpdateDestination unmashalling data", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}

	success := c.Repository.UpdateDestination(destination)
	if !success {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)

	return
}

//DeleteDestination ...
func (c *Controller) DeleteDestination(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	if err := c.Repository.DeleteDestination(id); err != "" {
		if strings.Contains(err, "404") {
			w.WriteHeader(http.StatusNotFound)
		} else if strings.Contains(err, "500") {
			w.WriteHeader(http.StatusInternalServerError)
		}

		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)

	return
}
