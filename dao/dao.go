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

//We're giving the IdeasDAO struct a Connect method.

func (i *IdeasDAO) Connect(){
  session, err := mgo.Dial(i.Server)
  if err != nil {
    log.Fatal(err)
  }
  db = session.DB(i.Database)
}

//Now we'll give the IdeasDAO struct methods for the rest of our queries.

//The FindAll method takes in no parameters, but outputs an array of Ideas. Note: bson.M == map

func (i *IdeasDAO) FindAll() ([]Idea, error){
  var ideas []Idea
  err := db.C(COLLECTION).Find(bson.M{}).All(&ideas)
  return ideas, err
}

//The FindById method takes in an id of type string, and outputs Idea. We use the Idea struct from our idea package in the models dir.

func (i *IdeasDAO) FindById(id string) (Idea, error){
  var idea Idea
  err := db.C(COLLECTION).FindId(bson.ObjectIdHex(id)).One(&idea)
  return idea, err
}

//Then we give the struct methods for post, delete, put

func (i *IdeasDAO) Insert(idea Idea) error{
  err := db.C(COLLECTION).Insert(&idea)
  return err
}

func (i *IdeasDAO) Delete(idea Idea) error {
  err := db.C(COLLECTION).Remove(&idea)
  return err
}

func (i *IdeasDAO) Update(idea Idea) error{
  err := db.C(COLLECTION).UpdateId(idea.ID, &idea)
  return err
}