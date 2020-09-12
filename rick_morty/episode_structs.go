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
      
      idCharacters = append(idsCharacters, characterID)
   }
   
   characters, err := GetCharactersArray(idsCharacters)
   if err != nil {
      return &MultipleCharacters{}, err
   }
   
   return characters, nil
}

func (e *MultipleEpisodes) GetCharacters() ([]MultipleCharacters, error) {
   charactersFromAllEpisodes := []MultipleCharacters{}
   
   for _, episode := range *e {
      characters, err := episode.GetCharacters()
      if err != nil {
         return []MultipleCharacters{}, err
      }
   
      charactersFromAllEpisodes = append(charactersFromAllEpisodes, *characters)
   }
   
   return charactersFromAllEpisodes, nil
}

func (ae *AllEpisodes) GetNextPage() (*AllEpisodes, error) {
   if ae.Infor.Next == "" {
      return &AllEpisodes{}, error.New("Nothing here")
   }
   
   pageString := getLastElementSplitedBy(ae.Info.Next, "=")
   
   page, err := strconv.Atoi(pageString)
   if err != nil {
      return &AllEpisodes{}, err
   }
   
   options := map[string]interface{}{"page": page}
   
   return GetEpisodes(options)
}
