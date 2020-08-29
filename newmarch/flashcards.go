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
   port := os.Args[1]
   
   http.HandleFunc("/", listFlashCards)
   fileServer := http.StripPrefix("/jscript/", http.FileServer(http.Dir("jscript")))
   http.Handle("/jscript/", fileServer)
   fileServer = http.StripPrefix("/html/", http.FileServer(http.Dir("html")))
   http.Handle("/html/", fileServer)
   
   http.HandleFunc("/flashcards.html", listFlashCards)
   http.HandleFunc("/flashcardSets", manageFlashCards)
   
   // deliver requests to the handlers
   err := http.ListenAndServe(port, nil)
   checkError(err)
}

func listFlashCards(rw http.ResponseWriter, req *http.Request) {
   flashCardsNames := flashcards.ListFlashCardsNames()
   t, err := template.ParseFiles("html/ListFlashcards.html")
   if err != nil {
      http.Error(rw, err.Error(), http.StatusInternalServerError)
      return
   }
   t.Execute(rw, flashCardNames)
}
