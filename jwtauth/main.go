package main
import (
    "encoding/json"
    "fmt"
    "log"
    "net/http"
    "os"
    "time"
    jwt "github.com/dgrijalva/jwt-go"
    "github.com/dgrijalva/jwt-go/request"
    "github.com/gorilla/mux"
)

var secretKey = []byte(os.Getenv("SESSION_SECRET"))
var users = map[string]string{"naren": "passme", "admin": "password"}
// Response is a representation of JSON response for JWT

type Response struct {
   Token string `json: "token"`
   Status string `json: "status"`
}

// HEalthcheckHandler returns the date and time
func HealthcheckHandler(w http.ResponseWriter, r *http.Request) {
   tokenString, err := request.HeaderExtractor{"access_token"}.ExtractToken(r)
   token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
      // Dont forget to validate the alg is what you expect:
