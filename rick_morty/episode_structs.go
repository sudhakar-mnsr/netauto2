package rickandmorty

import (
"errors"
"strconv"
)

type Episode struct {
   ID int `json:"id"`
   Name string `json:"name"`
   Air_Date string `json:"air_date"`
   Episode string `json:"episode"`
   Characters []string `json:"characters"`
   URL string `json:"url"`
   Created string `json:"created"`
}

type AllEpisodes struct {
   Info Info `json:"info"`
   Results MultipleEpisodes `json:"results"`
}

type MultipleEpisodes []Episode

func (e *Episode) GetCharacters() (*MultipleCharacters, error) {
var idsCharacters []int
for _, characterURL := range e.Characters {
characterIDString := getLastElementSplitedBy(characterURL, "/")
CharacterID, err := strconv.Atoi(characterIDString)
if err != nil {
   return &MultipleCharacters{}, err
}
