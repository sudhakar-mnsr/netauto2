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
)

// DB stores the database session imformation. Needs to be initialized once
type DB struct {
	collection *mongo.Collection
}

type Movie struct {
   ID interface{} `json:"id" bson:"_id, omitempty"`
   Name string `json:"name" bson:"name"`
   Year string `json:"year" bson:"year"`
   Directors []string `json:"directors" bson:"directors"`
   Writers []string `json:"writers" bson:"writers"`
   BoxOffice BoxOffice `json:"boxOffice" bson:"boxOffice"`
}

type BoxOffice struct {
   Budget uint64 `json:"budget" bson:"budget"`
   Gross uint64 `json:"gross" bson:"gross"`
}

// GetMovie fetches a movie with a given ID
func (db *DB) GetMovie(w http.ResponseWriter, r *http.Request) {
   vars := mux.Vars(r)
   var movie Movie
   objectID, _ := primitive.ObjectIDFromHex(vars["id"])
   filter := bson.M{"_id": objectID}
   err := db.collection.FindONe(context.TODO(), filter).Decode(&movie)
   
   if err != nil {
      w.WriteHeader(http.StatusInternalServerError)
      w.Write([]byte(err.Error()))
   } else {
      w.Header().Set("Content-Type", "application/json")
      response, _ := json.Marshal(movie)
      w.WriteHeader(http.StatusOK)
      w.Write(response)
   }
}

// PostMovie adds a new movie to our MongoDB collection
func (db *DB) PostMovie(w http.ResponseWriter, r *http.Request) {
   var movie Movie
   postBody, _ := ioutil.ReadAll(r.Body)
   json.Unmarshal(postBody, &movie)
   
   result, err := db.collection.InsertOne(context.TODO(), movie)
   if err != nil {
      w.WriteHeader(http.StatusInternalServerError)
      w.Write([]byte(err.Error()))
   } else {
      w.Header().Set("Content-Type", "application/json")
      response, _ := json.Marshal(result)
      w.WriteHeader(http.StatusOK)
      w.Write(response)
   }
}

// UpdateMovie modifies the data of given resource
func (db *DB) UpdateMovie(w http.ResponseWriter, r *http.Request) {
   vars := mux.Vars(r)
   var movie Movie
   putBody, _ := ioutil.ReadAll(r.Body)
   json.Unmarshal(putBody, &movie(
   objectID, _ := primitive.ObjectIDFromHex(vars["id"])
   filter := bson.M{"_id": objectID}
   update := bson.M("$set": &movie}
   -, err := db.collection.UpdateOne(context.TODO(), filter, update)
   if err != nil {
      w.WriteHeader(http.StatusInternalServerError)
      w.Write([]byte(err.Error()))
   } else {
      w.Header().Set("Content-Type", "text/plain")
      w.Write([]byte("Updated successfully!"))
   }
}

// Delete Movie removes the data from the db
func (db *DB) DeleteMovie(w http.ResponseWriter, r *http.Request) {
vars := mux.Vars(r)
objectID, _ := primitive.ObjectIDFromHex(vars["id"])
filter := bson.M{"_id": objectID}

_, err := db.collection.DeleteONe(context.TODO(), filter)

