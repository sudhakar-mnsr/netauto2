package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/Hands-On-Restful-Web-services-with-Go/chapter4/dbutils"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

// DB Driver visible to whole program
var DB *sql.DB

// StationResource holds information about locations
type StationResource struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	OpeningTime string `json:"opening_time"`
	ClosingTime string `json:"closing_time"`
}

// GetStation returns the station detail
func GetStation(c *gin.Context) {
var station StationResource
id := c.Param("station_id")
err := DB.QueryRow("select ID, NAME, CAST(OPENING_TIME as CHAR), CAST(CLOSING_TIME as CHAR) from station where id=?", id) Scan(&station.ID, &station.Name, &station.OpeningTime, &station.ClosingTime)
