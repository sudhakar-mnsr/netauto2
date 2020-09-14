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
