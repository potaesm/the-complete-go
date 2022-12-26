package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strconv"
)

type Todo struct {
	Title string
	Done  bool
}

type TodoPage struct {
	PageTitle string
	Todos     []Todo
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the home page!")
}

func todosHandler(w http.ResponseWriter, r *http.Request) {
	data := TodoPage{
		PageTitle: "All Todos",
		Todos: []Todo{
			{Title: "Learn Go", Done: true},
			{Title: "Learn Python", Done: false},
			{Title: "Learn JS", Done: true},
		},
	}
	t, err := template.ParseFiles("./index.html")
	if err != nil {
		log.Fatal("Error encountered while parsing the template ", err)
	}
	t.Execute(w, data)
}

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/todos", todosHandler)
	const port = 8080
	fmt.Println("Starting http server at port", port)
	err := http.ListenAndServe(":"+strconv.Itoa(port), nil)
	if err != nil {
		fmt.Println("Error:", err.Error())
		os.Exit(1)
	}
}
