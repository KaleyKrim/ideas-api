package dao

import (
	"log"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"ideas-api/models"
)

type IdeasDAO struct {
	Server   string
	Database string
}

var db *mgo.Database

const {
  COLLECTION = "ideas"
}

func (i *IdeasDAO) Connect(){
  session, err := mgo.Dial(m.Server)
  if err != nil {
    log.Fatal(err)
  }
  db = session.DB(i.Database)
}