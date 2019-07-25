package mongo

import (
	"fmt"
	"gopkg.in/mgo.v2/bson"
	"log"
	"waas/models"
)

//Repository ...
type Storer struct{
	Sesion Session
}

// DBNAME the name of the DB instance
const DBNAME = "waasdb"

// COLLECTION is the name of the collection in DB
const COLLECTION = "default"


// GetBusinessObjects returns the list of BusinessObjects
func (s Storer) GetBusinessObjects() models.BusinessObjects {
	session, err := NewSession()

	if err != nil {
		fmt.Println("Failed to establish connection to Mongo server:", err)
	}

	//defer session.Close()

	c := session.Copy().session.DB(DBNAME).C(COLLECTION)
	results := models.BusinessObjects{}

	if err := c.Find(nil).All(&results); err != nil {
		fmt.Println("Failed to write results:", err)
	}
	session.session.Close()
	return results
}


// PushObject adds a BusinessObject in the DB
func (s Storer) PushObject(businessObject models.BusinessObject) bool {
	session, err := NewSession()
	//defer session.Close()

	i := bson.NewObjectId()
	businessObject.ID=i.Hex()
	session.Copy().session.DB(DBNAME).C(COLLECTION).Insert(&businessObject)

	if err != nil {
		log.Fatal(err)
		return false
	}

	fmt.Println("Added New  BusinessObject Title- ", businessObject.ID)
	session.Close()
	return true
}


