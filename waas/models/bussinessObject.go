package models

//
type BusinessObject struct {
	ID     string     	 `bson:"_id" json:"_id,omitempty"`
}

//
type BusinessObjects []BusinessObject


