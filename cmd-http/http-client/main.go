package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strings"
)

func main() {
	/** HTTP Get **/
	{
		resp, err := http.Get("http://localhost:8090/users")
		panicOnError(err)
		defer resp.Body.Close()
		fmt.Println("HTTP Get Status Code:", resp.StatusCode)
		reader := bufio.NewReader(resp.Body)
		s := bufio.NewScanner(reader)
		for s.Scan() {
			fmt.Println(s.Text())
		}
		err = s.Err()
		panicOnError(err)
	}
	/** HTTP Post **/
	{
		json := `{"id":6,"name":"Mongoose","score":200}`
		b := strings.NewReader(json)
		resp, err := http.Post("http://localhost:8090/users", "application/json", b)
		panicOnError(err)
		defer resp.Body.Close()
		fmt.Println("HTTP Post Status Code :", resp.StatusCode)
	}
}

func panicOnError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
