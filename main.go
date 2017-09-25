package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2"
)

type Island struct {
	ID    string `json:"id, omitempty"`
	Name  string `json:"name,omitempty"`
	Image string `json:"image,omitempty"`
}

type Province struct {
	ID          string    `json:"id,omitempty"`
	Name        string    `json:"name,omitempty"`
	Intro       string    `json:"intro,omitempty"`
	Description string    `json:"description,omitempty"`
	Image       string    `json:"image,omitempty"`
	Location    *Location `json:"location,omitempty"`
	Island      *Island   `json:"island,omitempty"`
}

type City struct {
	ID          string    `json:"id,omitempty"`
	Name        string    `json:"name,omitempty"`
	Intro       string    `json:"intro,omitempty"`
	Description string    `json:"description,omitempty"`
	Image       string    `json:"image,omitempty"`
	Location    *Location `json:"location,omitempty"`
	Province    *Province `json:"province,omitempty"`
}

type Destination struct {
	ID          string    `json:"id,omitempty"`
	Name        string    `json:"name,omitempty"`
	Intro       string    `json:"intro,omitempty"`
	Description string    `json:"description,omitempty"`
	Image       string    `json:"image,omitempty"`
	Location    *Location `json:"location,omitempty"`
	City        *City     `json:"city,omitempty"`
}

type Category struct {
	ID          string `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	Intro       string `json:"intro,omitempty"`
	Description string `json:"description,omitempty"`
	Image       string `json:"image,omitempty"`
}

type Type struct {
	ID          string    `json:"id,omitempty"`
	Name        string    `json:"name,omitempty"`
	Intro       string    `json:"intro,omitempty"`
	Description string    `json:"description,omitempty"`
	Image       string    `json:"image,omitempty"`
	Category    *Category `json:"category,omitempty"`
}

type Location struct {
	Latitude  string `json:"latitude, omitempty"`
	Longitude string `json:"longitude, omitempty"`
}

var provinces []Province

func GetProvinceList(w http.ResponseWriter, req *http.Request) {
	json.NewEncoder(w).Encode(provinces)
}

func GetProvince(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	for _, item := range provinces {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Province{})
}

func CreateProvince(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	var province Province
	_ = json.NewDecoder(req.Body).Decode(&province)
	province.ID = params["id"]
	provinces = append(provinces, province)
	json.NewEncoder(w).Encode(provinces)
}

func DeleteProvince(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	for index, item := range provinces {
		if item.ID == params["id"] {
			provinces = append(provinces[:index], provinces[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(provinces)
}

func main() {
	router := mux.NewRouter()
	provinces = append(provinces, Province{ID: "1", Name: "Aceh", Intro: "Something about Aceh", Description: "More words about Aceh for the description"})
	provinces = append(provinces, Province{ID: "2", Name: "Medan", Intro: "Something about Medan", Description: "More words about Medan for the description"})
	router.HandleFunc("/provinces", GetProvinceList).Methods("GET")
	router.HandleFunc("/provinces/{id}", GetProvince).Methods("GET")
	router.HandleFunc("/provinces/{id}", CreateProvince).Methods("POST")
	router.HandleFunc("/provinces/{id}", DeleteProvince).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8118", router))

	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
}
