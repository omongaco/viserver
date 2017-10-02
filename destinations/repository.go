package destinations

import (
	"fmt"
	"log"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

//Repository ...
type Repository struct{}

//SERVER the DB server
const SERVER = "localhost:27017"

//DBNAME the name of DB instance
const DBNAME = "visitindonesiadb"

//DOCNAME the name of the Document
const DOCNAME = "destination"

//GetDestinations return the list of destinations
func (r Repository) GetDestinations() Destinations {
	session, err := mgo.Dial(SERVER)
	if err != nil {
		fmt.Println("Failed to establish connection to Mongo server:", err)
	}
	defer session.Close()

	c := session.DB(DBNAME).C(DOCNAME)
	results := Destinations{}
	if err := c.Find(nil).All(&results); err != nil {
		fmt.Println("Failed to write results:", err)
	}

	return results
}

//AddDestination insert Destination to DB
func (r Repository) AddDestination(destination Destination) bool {
	session, err := mgo.Dial(SERVER)
	defer session.Close()

	destination.ID = bson.NewObjectId()
	session.DB(DBNAME).C(DOCNAME).Insert(destination)

	if err != nil {
		log.Fatal(err)
		return false
	}

	return true
}

//UpdateDestination updates a Destination in the DB
func (r Repository) UpdateDestination(destination Destination) bool {
	session, err := mgo.Dial(SERVER)
	defer session.Close()

	destination.ID = bson.NewObjectId()
	session.DB(DBNAME).C(DOCNAME).UpdateId(destination.ID, destination)

	if err != nil {
		log.Fatal(err)
		return false
	}

	return true
}

//DeleteDestination deletes a Destination
func (r Repository) DeleteDestination(id string) string {
	session, err := mgo.Dial(SERVER)
	defer session.Close()

	//  Verify the id is an ObjectId, otherwise bail
	if !bson.IsObjectIdHex(id) {
		return "NOT FOUND"
	}

	//  Grab the ID
	oid := bson.ObjectIdHex(id)

	//  Remove User
	if err = session.DB(DBNAME).C(DOCNAME).RemoveId(oid); err != nil {
		log.Fatal(err)
		return "INTERNAL ERROR"
	}

	//  Write status "OK"
	return "OK"
}
