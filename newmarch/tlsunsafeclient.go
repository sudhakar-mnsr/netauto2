/* TLSUnsafeClientGet
   Go presently bails out when it encounters certificate errors
   However we are configuring to ignore certificate errors
   Transport configuration flag InsecureSkipVerify is set
   (NOT A GOOD PRACTISE)
   With self-signed certificates go will generate error of below
   x509: certificate signed by unknown authority
*/
package main

import (
   "fmt"
   "net/http"
   "net/url"
   "os"
   "strings"
   "crypto/tls"
)

