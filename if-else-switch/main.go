package main

import (
	"fmt"
	"time"
)

func main() {
	/** If-else **/
	playerName := "Suthinan Musitmani"
	age := 21
	// playerName = "Potaesm"
	if age >= 18 && playerName == "Suthinan Musitmani" {
		fmt.Println("You are authenticated")
	} else if age >= 18 && playerName == "Potaesm" {
		fmt.Println("You need permission")
	} else {
		fmt.Println("You are not allowed")
	}

	/** Switch **/
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the week end:", time.Now().Weekday())
		break
		// fallthrough
	default:
		fmt.Println("It's a week day:", time.Now().Weekday())
	}
}
