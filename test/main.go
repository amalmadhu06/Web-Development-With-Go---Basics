package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/gorilla/sessions"
)

// declaring tpl for parsing and storing html files
var tpl *template.Template

// declaring Store for storing new cookie which gets created when new user logs in
var Store = sessions.NewCookieStore([]byte("admin"))

type Page struct {
	Status  bool
	Header1 interface{}
	Valid   bool
}

// creating an instance of Page and setting the default value of Status to false
var P = Page{
	Status: false,
}

// hardcoding login credentials to a map
var userDB = map[string]string{
	"email":    "amal@gmail.com",
	"password": "amal",
}

// func init() which works first , storing parsed html files to the variable
func init() {
	tpl = template.Must(template.ParseGlob("*.html"))
}



func index(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "index.html", r)

	Username := r.FormValue("username")
	Password := r.FormValue("password")
	if Username == userDB["email"] && Password == userDB["password"] {
		fmt.Println("you have logged in successfully")
	} else {
		fmt.Println("incorrect password")
	}
}

func login(w http.ResponseWriter, r *http.Request) {

}

func loginHandler(w http.ResponseWriter, r *http.Request) {

}
func logOutHandler(w http.ResponseWriter, r *http.Request) {

}

func main() {

	//handling routes
	http.HandleFunc("/", index)
	http.HandleFunc("/login", login)
	http.HandleFunc("/login-submit", loginHandler)
	http.HandleFunc("/logout", logOutHandler)

	//server on localhost:8997
	fmt.Println("server starts at port 8997")
	http.ListenAndServe("localhost:8997", nil)
}
