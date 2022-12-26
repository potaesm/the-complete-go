package main

import (
	"fmt"

	"github.com/potaesm/the-complete-go/packages/packer/numbers"
	"github.com/potaesm/the-complete-go/packages/packer/strings"
	"github.com/potaesm/the-complete-go/packages/packer/strings/greeting"
)

func main() {
	fmt.Println("numbers.IsPrime(7):", numbers.IsPrime(7))
	fmt.Println("strings.Reverse(\"Suthinan\"):", strings.Reverse("Suthinan"))
	fmt.Println("greeting.Welcome:", greeting.Welcome)
	fmt.Println("greeting.MorningText:", greeting.MorningText)
	fmt.Println("greeting.EveningText:", greeting.EveningText)
}
