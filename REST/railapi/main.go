package main

import (
   "database/sql"
   "encoding/json"
   "log"
   "net/http"
   "time"
   "github.com/Hands-On-Restful-Web-services-with-Go/chapter4/dbutils"
   "github.com/emicklei/go-restful"
   _ "github.com/mattn/go-sqlite3"
)

// DB Driver visible to whole program
var DB *sql.DB

// TrainResource is the model for holding rail information
type TrainResource struct {
   ID int
   DriverName string
   OperatingStatus bool
}

// StationResource holds information about locations
type StationResource struct {
   ID int
   Name string
   OpeningTime time.Time
   ClosingTime time.Time
}

// Register adds paths and routes to container
func (t *TrainResource) Register(container *restful.Container) {
   ws := new(restful.WebService)
   ws.Path("/v1/trains").Consumes(restful.MIME_JSON).Produces(restful.MIME_JSON) // you can specify this per route as well
   ws.Route(ws.Get("/{train-id}").To(t.getTrain))
   ws.Route(ws.POST("").To(t.CreateTrain))
   ws.Route(ws.DELETE("/{train-id}").To(t.removeTrain))
   container.Add(ws)
}
