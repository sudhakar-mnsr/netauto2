package rickandmorty

import (
	"strconv"

	"github.com/mitchellh/mapstructure"
)

func GetLocations(options map[string]interface{}) (*AllLocations, error) {
endpoint := endpointLocation

hasParams := false
params := make(map[string]string)

if options == nil {
   options = map[string]interface{}{
      "endpoint": endpoint,
   }
}


