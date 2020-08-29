/* server
*/

package main

import (
"fmt"
"net/http"
"os"
"html/template"
)

import (
"dictionary"
"flashcards"
"templatefuncs"
)

func main() {
if len(os.Args) != 2 {
   fmt.Fprint(os.Stderr, "Usage: ", os.Args[0], ":port\n")
   os.Exit(1)
}
