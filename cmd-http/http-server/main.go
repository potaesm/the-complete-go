package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Score int    `json:"score"`
}

var Users = []User{
	{ID: 0, Name: "Suthinan", Score: 10},
	{ID: 1, Name: "Musitmani", Score: 9},
	{ID: 2, Name: "Potaesm", Score: 6},
}

func main() {
	http.HandleFunc("/users", handleUsers)
	http.HandleFunc("/", handleIndex)
	port := 8090
	fmt.Println("Starting http server at port", port)
	http.ListenAndServe(":"+strconv.Itoa(port), nil)
}

func handleIndex(w http.ResponseWriter, req *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	fmt.Fprintf(w, `{"message":"Hello World!"}`)
}

func handleUsers(w http.ResponseWriter, req *http.Request) {
	if req.Method == "GET" {
		handleGet(w, req)
	} else if req.Method == "POST" {
		handlePost(w, req)
	} else {
		handleError(w, req, http.StatusMethodNotAllowed, fmt.Errorf("Invalid Method"))
	}
}

func handleGet(w http.ResponseWriter, req *http.Request) {
	query := req.URL.Query()
	id := query.Get("id")
	var result []byte
	var err error
	if id == "" {
		result, err = json.Marshal(Users)
	} else {
		idInt, err := strconv.Atoi(id)
		if err == nil {
			result, err = json.Marshal(Users[idInt])
		}
	}
	if err != nil {
		handleError(w, req, http.StatusInternalServerError, err)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	fmt.Fprintf(w, string(result))
}

func handlePost(w http.ResponseWriter, req *http.Request) {
	var u User
	defer req.Body.Close()
	b, err := ioutil.ReadAll(req.Body)
	if err != nil {
		handleError(w, req, http.StatusInternalServerError, err)
		return
	}
	err = json.Unmarshal(b, &u)
	if err != nil {
		handleError(w, req, http.StatusInternalServerError, err)
		return
	}
	Users = append(Users, u)
	w.WriteHeader(http.StatusCreated)
}

func handleError(w http.ResponseWriter, req *http.Request, status int, err error) {
	w.WriteHeader(status)
	w.Header().Add("Content-type", "application/json")
	fmt.Fprintf(w, `{error:%v}`, err.Error())
}
