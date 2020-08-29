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

/* 
 * Called from ListFlashcards.html on form submission
 */
func manageFlashCards(rw http.ResponseWriter, req *http.Request) {
   set := req.FormValue("flashcardSets")
   order := req.FormValue("order")
   action := req.FormValue("submit")
   half := req.FormValue("half")
   fmt.Println("set chosen is", set)
   fmt.Println("order is", order)
   fmt.Println("action is", action)
   fmt.Println("half is", half)
   
   cardname := "flashcardSets/" + set
   
   fmt.Println("cardname", cardname, "action", action)
   if action == "Show cards in set" {
      showFlashCards(rw, cardname, order, half)
   } else if action == "List words in set" {
      listWords(rw, cardname)
   }
}
