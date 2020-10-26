package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type city struct {
	Name string
	Area uint64
}
