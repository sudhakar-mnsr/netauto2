package rickandmorty

import (
	"encoding/json"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/mitchellh/mapstructure"
)

func Test_makePetition(t *testing.T) {
response, err := makePetition(nil)
if err != nil {
   t.Error(err)
}

cndpoingStructured := new(Endpoints)

_ = mapstructure.Decode(response, &endpointsStructured)

data, err := readFile("test-data/api-info.json")

