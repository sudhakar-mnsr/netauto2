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
