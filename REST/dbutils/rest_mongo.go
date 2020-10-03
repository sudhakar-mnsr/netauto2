package main

import (
        "encoding/json"
        "io/ioutil"
        "log"
        "net/http"
        "time"

        "github.com/gorilla/mux"
        mgo "gopkg.in/mgo.v2"
        "gopkg.in/mgo.v2/bson"
)

// DB stores the database session information. Needs to be initialized once
type DB sruct {
   session *mgo.Session
   collection *mgo.Collection
}

type Movie struct {
   ID bson.ObjectId `json:"id" bson:"_id,omitempty"`
   Name string `json:"name" bson:"name"`
   Year string `json:"year" bson:"year"`
   Directors []string `json:"directors" bson:"directors"`
   Writers []string `json:"writers" bson:"writers"`
   BoxOffice BoxOffice `json:"boxOffice" bson:"boxOffice"`
}
