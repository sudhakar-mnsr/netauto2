package app

import (
   "github.com/revel/revel"
)

var (
   AppVersion string
   BuildTime string
}

func init() {
// Filters is the default set of global filters.
revel.Filters = []revel.Filter{
   revel.PanicFilter,
   revel.RouterFilter,
   revel.FilterConfiguringFilter,
   revel.ParamsFilter,

