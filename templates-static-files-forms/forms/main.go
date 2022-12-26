package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"strconv"
)

type ContactDetails struct {
	Email   string
	Subject string
	Message string
}

func main() {
	tem1 := template.Must(template.ParseFiles("form.html"))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			tem1.Execute(w, nil)
		}
		details := ContactDetails{
			Email:   r.FormValue("email"),
			Subject: r.FormValue("subject"),
			Message: r.FormValue("message"),
		}
		tem1.Execute(w, details)
	})
	const port = 8080
	fmt.Println("Starting http server at port", port)
	err := http.ListenAndServe(":"+strconv.Itoa(port), nil)
	if err != nil {
		fmt.Println("Error:", err.Error())
		os.Exit(1)
	}
}
