package main

//importing packages
import (
	"net/http"
	"strconv"
	"time"
)

// function CheckLastVisit
func CheckLastVisit(w http.ResponseWriter, r *http.Request) {

	c, err := r.Cookie("lastvisit")
	expiry := time.Now().AddDate(0, 0, 1) //(year, month, days) --> setting the expiry to one day

	//creating cookie variable which points to http.Cookie and sets the value there
	cookie := &http.Cookie{
		Name:    "lastvisit",
		Expires: expiry,
		Value:   strconv.FormatInt(time.Now().Unix(), 10),
	}

	//sending arguments to http.SetCookie
	http.SetCookie(w, cookie)

	//if no cookie is found, error will be thrown, then we will show "welcome to the website"
	//if cookie is present we will show "Welcome back! You last visited at: lastvisittime"
	if err != nil {
		w.Write(([]byte("Welcome to the website")))
	} else {
		lastTime, _ := strconv.ParseInt(c.Value, 10, 0)
		html := "Welcome back! You last visited at: "
		html = html + time.Unix(lastTime, 0).Format("15:04:05")
		w.Write([]byte(html))
	}
}

func main() {
	http.HandleFunc("/", CheckLastVisit)
	http.ListenAndServe(":8999", nil)
}
