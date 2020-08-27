/* GET
   Retrieve representation of a resource rather than just information
   Content of response is in response field Body (type io.ReadCloser).
 */

package main

import (
        "fmt"
        "net/http"
        "net/http/httputil"
        "os"
        "strings"
)
