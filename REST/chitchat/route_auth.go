package main

import (
"github.com/sausheong/gwp/Chapter_2_Go_Chitchat/data"
"net/http"
)

// GET /login
// Show the login page
func login(writer http.ResponseWriter, request *http.Request) {
   t := parseTemplateFiles("login.layout", "public.navbar", "login")
   t.Execute(writer, nil)
}

// GET /signup
// Show the signup page
func signup(writer http.ResponseWriter, request *http.Request) {
   generateHTML(writer, nil, "login.layout", "public.navbar", "signup")
}
