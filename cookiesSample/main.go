package main

//importing packages
import (
	"net/http"
	"strconv"
	"time"
)

// function CheckLastVisit
func CheckLastVisit(w http.ResponseWriter, r *http.Request) {

	//nameed cookie is received to 'c'. If multiple cookies found, only one cookie will be returned
	//if no cookie is found, errNoCookie is returned and stored in  'err'
	c, err := r.Cookie("lastvisit")
	expiry := time.Now().AddDate(0, 0, 1) //(year, month, days) --> setting the expiry to one day

	//creating cookie variable which points to http.Cookie and sets the value there
	//Cookie is a struct . It represents an HTTP cookie as sent in the Set-Cookie header of an
	// HTTP response or the Cookie header of an HTTP request
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
