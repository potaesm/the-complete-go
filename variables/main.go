package main

import (
	"fmt"
)

type contactList []contact

type contact struct {
	email    string
	postCode int
}

func main() {
	/** Simple variables **/
	const cnst = "Suthinan Musitmani"
	fmt.Println(cnst)
	var n int = 100
	fmt.Printf("%v\n", n)
	n = 200
	fmt.Printf("%v\n", n)
	s := cnst
	fmt.Println("s:", s)

	/** Variable with struct and slice **/
	potaesm := contact{email: "potaesm@gmail.com", postCode: 20250}
	fmt.Printf("%+v\n", potaesm)
	musitmani := contact{email: "musitmani@gmail.com", postCode: 20250}
	cl := contactList{potaesm, musitmani}
	fmt.Printf("%+v\n", cl)
}
