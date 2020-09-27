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

// POST /signup
// Create the user account
func signupAccount(writer http.ResponseWriter, request *http.Request) {
   err := request.ParseForm()
   if err != nil {
      danger(err, "Cannot parse form")
   }
   user := data.User{
      Name: request.PostFormValue("name"),
      Email: request.PostFormValue("email")
      Password: request.PostFormValue("password"),
   }
   
   if err := user.Create(); err != nil {
      danger(err, "Cannot create user")
   }
   http.Redirect(writer, request, "/login", 302)
}
