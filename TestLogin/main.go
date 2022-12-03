package main

import (
	"fmt"

	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/sessions"
	"tawesoft.co.uk/go/dialog"
)

var tpl *template.Template
var Store = sessions.NewCookieStore([]byte("francis"))

func init() {
	tpl = template.Must(template.ParseGlob("Template/*.html"))
}

type Page struct {
	Status  bool
	Header1 interface{}
	Valid   bool
}

var userDB = map[string]string{
	"email":    "amal@gmail.com",
	"password": "amal123",
}
var P = Page{
	Status: false,
}

func login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Cache-Control", "no-cache,no-store,must-revalidate")

	ok := Middleware(w, r)

	if ok {

		http.Redirect(w, r, "/login-submit", http.StatusSeeOther)
		return
	}
	P.Valid = Middleware(w, r)
	filename := "login.html"
	err := tpl.ExecuteTemplate(w, filename, P)
	if err != nil {
		fmt.Println("error while parsing file", err)
		return
	}

}

func loginHandler(w http.ResponseWriter, r *http.Request) {

	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "there is an error parsing %v", err)
		return
	}
	emails := r.PostForm.Get("email")
	fmt.Println(emails)

	password := r.PostForm.Get("password")
	fmt.Println(password)

	if userDB["email"] == emails && userDB["password"] == password && r.Method == "POST" {

		session, _ := Store.Get(r, "started")

		session.Values["id"] = "amal"
		P.Header1 = session.Values["id"]
		fmt.Println(P.Header1)
		session.Save(r, w)

		fmt.Println(session)

		w.Header().Set("Cache-Control", "no-cache,no-store,must-revalidate")

		http.Redirect(w, r, "/", http.StatusSeeOther)

	} else {
		dialog.Alert("Action Not allowed")
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return

	}

}
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

func Middleware(w http.ResponseWriter, r *http.Request) bool {
	session, _ := Store.Get(r, "started")

	if session.Values["id"] == nil {
		return false
	}
	P.Header1 = session.Values["id"]
	return true

}

func index(w http.ResponseWriter, r *http.Request) {
	ok := Middleware(w, r)
	if ok {
		P.Status = true

	}
	filenamE := "login.html"
	err := tpl.ExecuteTemplate(w, filenamE, P)
	if err != nil {
		fmt.Println("error while parsing file", err)
		return
	}

}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/login-submit", loginHandler)
	http.HandleFunc("/login", login)
	http.HandleFunc("/logout", Logouthandler)
	fmt.Println("server starts at port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}

}
