package rickandmorty

import (
   "strconv"
   "github.com/mitchellh/mapstructure"
)

func GetEpisodes(options map[string]interface{}) (*AllEpisodes, error) {
   endpoint := endpointEpisode
   hasParams := false
   params := make(map[string]string)
   
   if options == nil {
      options = map[string]interface{} {
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
      return &AllEpisodes{}, err
   }
   episodes := new(AllEpisodes)
   
   if err := mapstructure.Decode(data, &episodes); err != nil {
   return &AllEpisodes{}, err
   }
   
   return episodes, nil
}
