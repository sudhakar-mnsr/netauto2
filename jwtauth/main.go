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
      if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
         return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
      }
      // hmacSampleSecret is a []byte containing your secret, 
      // e.g []byte("my_secret_key")
      return secretKey, nil
   })
   if err != nil {
      w.WriteHeader(http.StatusForbidden)
      w.Write([]byte("Access Denied; Please check the access token"))
      return
   }
   if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
      // if token is valid
      response := make(map[string]string)
      // response["user"] = claims["username"]
      response["time"] = time.Now().String()
      response["user"] = claims["username"].(string)
      responseJSON, _ := json.Marshal(response)
      w.Write(responseJSON)
   } else {
      w.WriteHeader(http.StatusForbidden)
      w.Write([]byte(err.Error()))
   }
}
