package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
)

// initializing store to hold the session values	
var store = sessions.NewCookieStore(
	[]byte("secret-session"))

func handler(w http.ResponseWriter, r *http.Request) {

	//retreving session called session-name using .Get function
	//if session is already there, we will get access to session
	//if not present, new session will be created
	session, err := store.Get(r, "session-name")

	if err != nil {
		http.Error(w, err.Error(),
			http.StatusInternalServerError)
		return
	}

	// once we have a session,
	//settng some values in the session
	session.Values["abc"] = "cba"
	session.Values[111] = 222

	//saving the session in response
	session.Save(r, w)
}


func main() {
	router := mux.NewRouter()
	http.Handle("/", router)
	router.HandleFunc("/", handler)
	http.ListenAndServe(":8998", nil)
}
