# Login Page with Session Handling

This is a simple login page with session handling enabled. Users will be asked to enter their email id and password which they used to create the account. If both password and email id is correct, user will be given access to the website. Once user enters correct credentials, a new session will be created and it be used next time when they again try to access the webpage. 

## Packages used

 ### 1.  `html\template`
For rendering HTML templates
 ### 2.  `log`
 For logging messages
  ### 3.  `net\http`
For serving HTTP requests
 ### 4.  `github.com\gorilla\sessions`
 For storing sessions data
  ### 6.  `tawesoft.co.uk/go/dialog`
For displaying alert messages

## How it works
A struct called Page is defined and an instance of it, P is created to store the session details
```go
type  Page  struct {
	Status bool
	Header1 interface{}
	Valid bool
}
```
A map called UserDB is defined to store predefined email and password values for validating user login

```go
var  userDB  =  map[string]string{
	"password": "amal@123xyz",
	"email": "amal@gmail.com",
}
```

### Functions

- `intit()` <br>
This function is called automatically when the code is run, and it is used to parse HTML files inside a directory called template.

- `login()` <br>
This function is called when the user accesses the `/login` URL, and it is used to render the login page.
- `loginHandler()` <br>
This function is called when the user submits the login form, and it is used to validate the user's email and password. If the login is successful, it redirects the user to the home page and stores their session data.
- `LogoutHandler()` <brx`>
This function is called when the user logs out, and it is used to clear the user's session data and redirect them to the login page.