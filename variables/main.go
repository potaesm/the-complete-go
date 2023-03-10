package main

import (
	"fmt"
)

type contactList []contact

type contact struct {
	email    string
	postCode int
}

var globalVariable string

const (
	simpleGlobalConstant  = "Suthinan Musitmani"
	anotherGlobalConstant = "Potaesm"
)

func main() {
	/** Simple variables **/
	fmt.Println("globalVariable:", globalVariable)
	globalVariable = "Musitmani"
	fmt.Println("globalVariable:", globalVariable)
	fmt.Println("simpleGlobalConstant:", simpleGlobalConstant)
	fmt.Println("anotherGlobalConstant:", anotherGlobalConstant)
	const localConstant = "Suthinan Musitmani"
	fmt.Println(localConstant)
	var n int = 100
	fmt.Printf("%v\n", n)
	n = 200
	fmt.Printf("%v\n", n)
	fmt.Print("Enter number: ")
	fmt.Scanln(&n)
	fmt.Printf("Your input: %v\n", n)
	s := localConstant
	fmt.Println("s:", s)

	/** Variable with struct and slice **/
	potaesm := contact{email: "potaesm@gmail.com", postCode: 20250}
	fmt.Printf("%+v\n", potaesm)
	musitmani := contact{email: "musitmani@gmail.com", postCode: 20250}
	cl := contactList{potaesm, musitmani}
	fmt.Printf("%+v\n", cl)
}
