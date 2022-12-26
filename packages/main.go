package main

import (
	"fmt"

	"github.com/potaesm/the-complete-go/packages/packer/numbers"
	"github.com/potaesm/the-complete-go/packages/packer/strings"
	"github.com/potaesm/the-complete-go/packages/packer/strings/greeting"
)

func main() {
	fmt.Println(numbers.IsPrime(20))
	fmt.Println(strings.Reverse("demha"))
	fmt.Println(greeting.Welcome)
	fmt.Println(greeting.MorningText)
	fmt.Println(greeting.EveningText)
}
