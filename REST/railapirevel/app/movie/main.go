package main

import (
"context"
"encoding/json"
"io/ioutil"
"log"
"net/http"
"time"
"go.mongodb.org/mongo-driver/bson/primitive"
"go.mongodb.org/mongo-driver/bson"
"github.com/gorilla/mux"
"go.mongodb.org/mongo-driver/mongo"
"go.mongodb.org/mongo-driver/mongo/options"
}

// DB stores the database session information. Needs to be initialized once
type DB struct {
   collection *mongo.Collection
}

type Movie struct {
	ID        interface{} `json:"id" bson:"_id,omitempty"`
	Name      string      `json:"name" bson:"name"`
	Year      string      `json:"year" bson:"year"`
	Directors []string    `json:"directors" bson:"directors"`
	Writers   []string    `json:"writers" bson:"writers"`
	BoxOffice BoxOffice   `json:"boxOffice" bson:"boxOffice"`
}

type BoxOffice struct {
	Budget uint64 `json:"budget" bson:"budget"`
	Gross  uint64 `json:"gross" bson:"gross"`
}

// GetMovie feature a movie with given ID
func (db DB) GetMovie(w http.ResponseWriter, r *http.Request) {
   vars := mux.Vars(r)
   var movie Movie
   objectID, _ := primitive.ObjectIDFromHex(vars["id"])
   filter := bson.M{_id": objectID}
   err := db.collection.FindONe(context.TODO(), filter).Decode(&movie)
   
   if err != nil {
      w.WriteHeader(http.StatusInternalServer)
      w.Write([]byte(err.Error())
   } else {
      w.Header().Ser("Content-Type", application/json")
      response, _ := json.Marshal(movie)
      w.WriteHeader(http.StatusOK)
      w.Write(response)
   }
}
