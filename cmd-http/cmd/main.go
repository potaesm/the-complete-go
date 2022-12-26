package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	/** Read CMD Arguments **/
	{
		args := os.Args
		path := os.Args[0]
		fmt.Printf("path: %v\n", path)
		args = os.Args[1:]
		for i, v := range args {
			fmt.Printf("Index: %v, Value: %v \n", i, v)
		}
	}
	/** Read CMD Flags **/
	{
		fname := flag.String("fname", "", "Firstname as string")
		score := flag.Int("score", 0, "Score as int")
		start := flag.Bool("start", false, "Start as boolean")
		var lname string
		flag.StringVar(&lname, "lname", "", "Lastname as string")
		flag.Parse()
		fmt.Printf("Firstname: %v, Score: %v, Lastname: %v, Start: %v\n", *fname, *score, lname, *start)
		// fmt.Println("Run of the args", flag.Args())
	}

	/** Read ENV **/
	{
		fmt.Println("GAME_SCORE", os.Getenv("GAME_SCORE"))
		os.Setenv("GAME_SCORE", "100")
		fmt.Println("GAME_SCORE", os.Getenv("GAME_SCORE"))
		fmt.Println("ENV List")
		for _, e := range os.Environ() {
			envRow := strings.Split(e, "=")
			fmt.Printf("Key: %v, Value: %v\n", envRow[0], envRow[1])
		}
	}

	// TEST_ENV=Suthinan go run main.go -fname Suthinan -score 123 -start arg1 arg2
	// TEST_ENV=Suthinan ./main -fname Suthinan -score 123 -start arg1 arg2
}
