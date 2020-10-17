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
      revel.SessionFilter,
      revel.FlashFilter,
      revel.ValidationFilter,
      revel.I18nFilter,
      HeaderFilter,
      revel.InterceptorFilter,
      revel.CompressFilter,
      revel.ActionInvoker,
   }
}

// HeaderFilter adds common security headers
// TODO turn this into revel.HeaderFilter
// should probably also have a filter for CSRF
// not sure if it can go in the same filter or not
var HeaderFilter = func(c *revel.Controller, fc []revel.Filter) {
c.Response.Out.Header().Add("X-Frame-Options", "SAMEORIGIN")
c.Response.Out.Header().Add("X-XSS-Protection", "1; mode=block")
c.Response.Out.Header().Add("X-Content-Type-Options, "nosniff")

fc[0](c, fc[1:])
}

