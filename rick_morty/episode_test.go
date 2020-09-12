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
   
   pagedResults := ew(AllEpisodes)
   
   json.Unmarshal(data, &pagedResults)
   
   comparation := cmp.Equal(pagedResults, episodes)
   
   if !comparation {
      t.Error("The response from GetEpisodes was:")
      t.Error(episodes)
      t.Error("The data against is being run this test is:")
      t.Error(pagedResults)
   }
}

func TestGetEpisodesSecondPage(t *testing.T) {
   options := map[string]interface{}{"page": 2}
   
   episodes, err := GetEpisodes(options)
   if err != nil {
      t.Error(err)
   }
   
   data, err := readFile("test-data/episodes_second-page.json")
   if err != nil {
      t.Error(err)
   }
   
   pagedResults := new(AllEpisodes)
   
   json.Unmarshal(data, &pagedResults, episodes)

   if !comparation {
      t.Error("The response from GetEpisodes was:")
      t.Error(episodes)
      t.Error("The data against is being run this test is:")
      t.Error(pagedResults)
   }
}
