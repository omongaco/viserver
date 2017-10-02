package destinations

import "gopkg.in/mgo.v2/bson"

//Destination represent the destination article
type Destination struct {
	ID          bson.ObjectId `bson:"_id"`
	Name        string        `json:"name"`
	Description string        `json:"description"`
	Image       string        `json:"image"`
}

//Destinations is the array of Destination
type Destinations []Destination
