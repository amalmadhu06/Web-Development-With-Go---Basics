package main

import (
	"html/template"
	"os"
)

func main() {
	// Define a simple HTML template
	const tmpl = `
    <p>Hello, {{.Name}}!</p>
    <p>Today is {{.Date}}</p>
  `

	// Parse the template
	t, err := template.New("hello").Parse(tmpl)
	if err != nil {
		panic(err)
	}

	value := map[string]string{
		"Name": "John Doe",
		"Date": "Tuesday",
	}

	// Execute the template and print the result to stdout
	err = t.Execute(os.Stdout, value)
	if err != nil {
		panic(err)
	}
}
