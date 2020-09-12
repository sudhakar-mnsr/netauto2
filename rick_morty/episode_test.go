package rickandmorty

import (
   "endcoding/json"
   "testing"
   "github.com/google/go-cmp/cmp"
)

func TestGetEpisodesFirstPage(t *testing.T) {
options := map[string]interface{}{"page": 1}

episodes, err := GetEpisodes(options)
if err != nil {
   t.Error(err)
}

data, err := readFile("test-data/episodes_first-page.json")
if err != nil {
   t.Error(err)
}

