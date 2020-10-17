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
