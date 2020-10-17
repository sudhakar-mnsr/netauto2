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
