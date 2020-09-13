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
   
   for k, v := range options {
      switch v.(type) {
      case int:
      if k == "page" {
         hasParams = true
         params[k] = strconv.FormatInt(int64(v.(int)), 10)
      }
      delete(options, k)
      case string:
      if k == "endpoint" {
         continue
      }
      // Valid parameters to be passed to the parameters map
      validParams := []string{"name", "status", "species", "type", "gender"}
      exists := containsString(validParams, k)
      if exists {
         hasParams = true
         params[k] = v.(string)
      }
      delete(options, k)
      default:
      delete(options, k)
      options["endpoint"] = endpoint
      }
   }

   if hasParams {
      options["endpoint"] = endpoint
      options["params"] = params
   }
   
   data, err := makePetition(options)
   if err != nil {
      return &AllLocations{}, err
   }
   
   locations := new(AllLocations)
   
   if err := mapstructure.Decode(data, &locations); err != nil {
      return &AllLocations{}, err
   }
   
   return locations, nil
}

func GetLocation(integer int) (*Location, error) {
   endpoint := endpointLocation
   
   options := map[string]interface{}{
      "endpoint": endpoint,
      "params": map[string]int{
                "integer": integer,
      },
   }
   
   data, err := makePetition(options)
   if err != nil {
      return &Location{}, err
   }
   
   location := new(Location)
   
   if err := mapstructure.Decode(data, &location); err != nil {
      return &Location{}, err
   }
   
   return location, nil
}
