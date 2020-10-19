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
