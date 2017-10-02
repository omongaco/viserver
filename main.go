package main

import (
	"log"
	"net/http"

	"github.com/omongaco/viserver/destinations"

	"github.com/gorilla/handlers"
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
	Type        *Type     `json:"type, omitempty"`
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

func main() {
	//	Creating the routers
	mRouter := destinations.NewRouter()

	//	Allow access from the front-end side to the methods
	allowedOrigins := handlers.AllowedOrigins([]string{"*"})
	allowedMethods := handlers.AllowedMethods([]string{"GET", "POST", "DELETE", "PUT"})

	//	Launch server with CORS validation
	log.Fatal(http.ListenAndServe(":9000", handlers.CORS(allowedOrigins, allowedMethods)(mRouter)))
}
