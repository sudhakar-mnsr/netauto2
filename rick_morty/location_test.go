finitions or references in the same repository. Learn more

package rickandmorty

import (
	"encoding/json"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestGetLocationsFirstPage(t *testing.T) {
   options := map[string]interface{}{"page": 1}
   
   locations, err := GetLocations(options)
   if err != nil {
      t.Error(err)
   }
   
   data, err := readFile("test-data/locations_first-page.json")
   if err != nil {
      t.Error(err)
   }
   
   pagedResults := new(AllLocations)
   
   json.Unmarshal(data, &pagedResults)
   
   opt := sliceEmptyNullReturnType()
   
   comparation := cmp.Equal(pagedResults, locations, opt)
   
   if !comparation {
      t.Error("The response from GetLocations was:")
      t.Error(locations)
      t.Error("The data against is being run this test is:")
      t.Error(pagedResults)
   }
}

func TestGetLocationsFifthPage(t *testing.T) {
   options := map[string]interface{}{"page": 5}
   
   locations, err := GetLocations(options)
   if err != nil {
      t.Error(err)
   }
   data, err := readFile("test-data/locations_fifth-page.json")
   if err != nil {
      t.Error(err)
   }
   
   pagedResults := new(AllLocations)
   
   json.Unmarshal(data, &pagedResults)
   
   comparation := cmp.Equal(pagedResults, locations)
   
   if !comparation {
      t.Error("The response from GetLocations was:")
      t.Error(locations)
      t.Error("The data against is being run this test is:")
      t.Error(pagedResults)
   }
}

func TestGetLocationsWithParamNil(t *testing.T) {
   locations, err := GetLocations(nil)
   if err != nil {
      t.Error(err)
   }
   
   data, err := readFile(test-data/locations_first-page.json")
   if err != nil {
      t.Error(err)
   }
   
   pagedResults := new(AllLocations)
   
   json.Unmarshal(data, &pagedResults)
   
   opt := sliceEmptyNullReturnTrue()
   
   comparation := cmp.Equal(pagedResults, locations, opt)
   
   if !comparation {
      t.Error("The response from GetLocations was:")
      t.Error(locations)
      t.Error("The data against is being run this test is:")
      t.Error(pagedResults)
   }
}

func TestGetLocationsWithFilterParams(t *testing.T) {
   options := map[string]interface{}{
      "name": "earth",
   }
   
   locations, err := GetLocations(options)
   if err != nil {
      t.Error(err)
   }
   
   data, err := readFile("test-data/locations_filter_name-earth.json")
   if err != nil {
      t.Error(err)
   }
   
   pagedResults := new(AllLocations)
   
   json.Unmarshal(data, &pagedResults)
   
   opt := sliceEmptyNullReturnTrue()
   
   comparation := cmp.Equal(pagedResults, locations, opt)
   
   if ! comparation {
      t.Error("The response from GetLocations was:")
      t.Error(locations)
      t.Error("The data against is being run this test is:")
      t.Error(pagedResults)
   }
}

func TestGetLocationsWithFilterParamsCombined(t *testing.T) {
options := map[string]interface{}{
           "name": "earth",
           "type": "planet",
}

locations, err := GetLocations(options)
if err != nil {
   t.Error(err)
}

data, err := readFile("test-data/locations_filter_name-earth_type-planet.json")
if err != nil {
   t.Error(err)
}

pagedResults := new(Allocations)

