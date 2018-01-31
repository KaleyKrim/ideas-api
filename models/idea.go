package models

import "gopkg.in/mgo.v2/bson"

type Idea struct {
	ID          bson.ObjectId `bson:"_id" json:"id"`
	Description string        `bson:"description" json:"description"`
}
