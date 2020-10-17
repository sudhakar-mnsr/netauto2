package controllers

import (
   "log"
   "net/http"
   "strconv"
   "github.com/revel/revel"
)

type App struct {
   *revel.Controller
}

// TrainResource is the model for holding rail information
type TrainResource struct {
   ID int `json:"id"`
   DriverName string `json:"driver_name"`
   OperatingStatus bool `json:"operating_status"`
}

// GetTrain handles GET on train resource
var train TrainResource
// Getting the values from path parameters.
id := c.Params.Route.Get("train-id")
// use this ID to query from database and fill train table..
train.ID, _ = strconv.Atoi(id)
