package main

import (
   "database/sql"
   "encoding/json"
   "log"
   "net/http"
   "time"
   "github.com/emicklei/go-restful"
   _ "github.com/mattn/go-sqlite3"
   "github.com/sudhakar-mnsr/dbutils"
)

var DB *sql.DB
// TrainResource is model for holding rail
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

// ScheduleResource links both trains and stations
type ScheduleResource struct {
   ID int
   TrainID int
   StationID int
   ArrivalTime time.Time
}

// Register adds paths and routes to container
func (t *TrainResource) Register(container *restful.Container) {
   ws := new(restful.WebService)
   ws.Path("/v1/trains").Consumes(restful.MIME_JSON).Produces(restful.MIME_JSON)
   ws.Route(ws.GET("/{train-id}").To(t.getTrain))
   ws.Route(ws.POST("").To(t.CreateTrain))
   ws.Route(ws.Delete("/{train-id}").To(t.removeTrain))
   container.Add(ws)
}

func main() {
   db, err :=sql.Open("sqlite3", "./railapi.db")
   if err != nil {
      log.Println("Driver creation failed!")
   }
   dbutils.Initialize(db)
}

