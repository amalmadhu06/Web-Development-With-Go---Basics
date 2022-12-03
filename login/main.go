package main

import (
	"fmt"

	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/sessions"
	"tawesoft.co.uk/go/dialog"
)

// creating a variable with type *template.Template for storing the parsed files
var tpl *template.Template

// for saving cookies
var Store = sessions.NewCookieStore([]byte("admin"))

// function init which parses html files inside the template folder
func init() {
	tpl = template.Must(template.ParseGlob("template/*.html"))
}

// struct for storing session details
type Page struct {
	Status  bool
	Header1 interface{}
	Valid   bool
}

// storing pre defined password and email in this map for validating
var userDB = map[string]string{
	"password": "amal@123xyz",
	"email":    "amal@gmail.com",
}

// creating a struct instance P which sets initial value of Status to false
var P = Page{
	Status: false,
}

// function login which gets called when we access /login
func login(w http.ResponseWriter, r *http.Request) {

	//setting the response header to the following values
	//response header - Cache-Control
	w.Header().Set("Cache-Control", "no-cache,no-store,must-revalidate")
	ok := Middleware(w, r)

	//redirecting to /login-submit is value in ok is true
	if ok {
		http.Redirect(w, r, "/login-submit", http.StatusSeeOther)
		return
	}
	P.Valid = Middleware(w, r)
	filename := "login.html"
	err := tpl.ExecuteTemplate(w, filename, P)

	//showing the given comment if there occurs and error
	if err != nil {
		fmt.Println("error while parsing file", err)
		return
	}

}

// function loginHandler which gets called when we access /login-submit
func loginHandler(w http.ResponseWriter, r *http.Request) {

	//when an error occurs
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "there is an error parsing %v", err)
		return
	}

	//receiving values from the user inputs
	emails := r.PostForm.Get("email")
	password := r.PostForm.Get("password")

	// if email and password is correct
	if userDB["email"] == emails && userDB["password"] == password && r.Method == "POST" {

		session, _ := Store.Get(r, "started")

		session.Values["id"] = "Amal"
		P.Header1 = session.Values["id"]
		fmt.Println(P.Header1)
		session.Save(r, w)

		fmt.Println(session)

		w.Header().Set("Cache-Control", "no-cache,no-store,must-revalidate")

		http.Redirect(w, r, "/", http.StatusSeeOther)

		//if email and password is not correct
	} else {

		//showing alert box
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		dialog.Alert("Incorrect Email or Password. Try again")
		return
	}

}

// function Logouthandler which get called when we access /logout

func Logouthandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Cache-Control", "no-cache,must-revalidate")

	if P.Status {
		session, _ := Store.Get(r, "started")
		session.Options.MaxAge = -1
		session.Save(r, w)
		http.Redirect(w, r, "/", http.StatusSeeOther)
		P.Status = false
	} else if !P.Status {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	}
}

// function Middleware for session handling
func Middleware(w http.ResponseWriter, r *http.Request) bool {
	session, _ := Store.Get(r, "started")

	if session.Values["id"] == nil {
		return false
	}
	P.Header1 = session.Values["id"]
	return true
}

// function index which gets called once the server is connected. ("/")
func index(w http.ResponseWriter, r *http.Request) {
	ok := Middleware(w, r)
	if ok {
		P.Status = true

	}
	filenamE := "index.html"

	//checking for error and handling it
	err := tpl.ExecuteTemplate(w, filenamE, P)
	if err != nil {
		fmt.Println("error while parsing file", err)
		return
	}

}

// function main ()
func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/login-submit", loginHandler)
	http.HandleFunc("/login", login)
	http.HandleFunc("/logout", Logouthandler)
	fmt.Println("server starts at port 8080")

	//error handling
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
