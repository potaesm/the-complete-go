package main

import (
	"errors"
	"fmt"
)

/** No Return **/
func print(s string) {
	fmt.Println("s:", s)
}

/** No Return & Multiple Params **/
func printParams(params ...interface{}) {
	fmt.Println("params:", params, "lenParams", len(params))
}

/** 2 Integer Params & Return Integer **/
func add(x int, y int) int {
	return x + y
}
func add2(x, y int) int {
	return x + y
}

/** 2 String Params & Return 2 Results **/
func concat(a, b string) (string, int) {
	r := a + b
	return r, len(r)
}

/** 2 String Params & Return Named Variables **/
func concat2(a, b string) (r string, l int) {
	r = a + b
	l = len(r)
	return
}

/** Return OK or Error **/
func returnError(b bool) (string, error) {
	if b {
		return "", errors.New("Some error")
	}
	return "OK", nil
}

/** Pass By Value **/
func passByValue(s string) {
	fmt.Println("Original value:", s)
	s = "Potaesm"
}

/** Pass By Reference **/
func passByReference(s *string) {
	fmt.Println("Memeory address:", &s)
	fmt.Println("Original value:", *s)
	*s = "Potaesm"
}

/** Pass Nil As Pointer **/
func defaultValuePointer(s *string) {
	fmt.Println("s:", s)
}

/** Player & Game Struct **/
type player struct {
	firstname string
	lastname  string
	score     int
}
type game struct {
	players map[string]*player
}

func (g *game) getWinner() *player {
	var winner *player
	var maxScore int
	for _, value := range g.players {
		if value.score > maxScore {
			winner = value
			maxScore = value.score
		}
	}

	return winner
}

/** Chain & Base Struct **/
type chain struct {
	base
	value int
	next  *chain
}

func (c *chain) sayOK() {
	fmt.Println("Chain OK!")
}

type base struct{}

func (b *base) sayHi() {
	fmt.Println("Base Hi!")
}
func (c *base) sayOK() {
	fmt.Println("Base OK!")
}

func main() {
	/** Function **/
	print("Suthinan")
	printParams("Potaesm", 2, "Hi", true, 4.5)
	fmt.Println("add(2, 3):", add(2, 3))
	fmt.Println("add2(2, 4):", add2(2, 4))
	res, len := concat("Suthinan", "Musitmani")
	fmt.Println("concat:", res, len)
	res, len = concat2("Suthinan", "Potaesm")
	fmt.Println("concat2:", res, len)
	r, err := returnError(true)
	if err != nil {
		fmt.Println("Error:", err.Error())
	} else {
		fmt.Println("Result", r)
	}
	var sumClone func(int, int) int
	sumClone = add
	fmt.Println("sumClone(3, 5):", sumClone(3, 5))

	/** Pointer **/
	name := "Suthinan"
	passByValue(name)
	fmt.Println("passByValue Result:", name)
	passByReference(&name)
	fmt.Println("passByReference Result:", name)
	var namePointer *string
	namePointer = &name
	fmt.Println("namePointer:", namePointer)
	fmt.Println("namePointer value:", *namePointer)
	defaultValuePointer(nil)

	/** Player & Game Struct **/
	g := &game{
		players: make(map[string]*player),
	}
	fmt.Println("g:", g)
	p1 := player{}
	fmt.Println("p1:", p1)
	g.players["p1"] = &p1
	fmt.Println("g:", g)
	p2 := player{}
	p2.score = 4
	fmt.Println("p2:", p2)
	g.players["p2"] = &p2
	fmt.Println("g:", g)
	p3 := player{
		lastname:  "Musitmani",
		firstname: "Suthinan",
	}
	p3.score = 10
	fmt.Println("p3:", p3)
	g.players["p3"] = &p3
	fmt.Println("g:", g)

	fmt.Println("g.players[\"p2\"]", g.players["p2"])
	fmt.Println("*g.players[\"p2\"]", *g.players["p2"])

	winner := g.getWinner()
	fmt.Println("winner:", winner)
	fmt.Println("*winner:", *winner)

	/** Base & Chain Struct **/
	b := base{}
	b.sayHi()
	b.sayOK()

	c1 := chain{value: 100}
	fmt.Println("c1:", c1)
	// c1.base.sayHi()
	c1.sayHi()
	c2 := chain{value: 200}
	c1.next = &c2
	fmt.Println("c1:", c1)
	fmt.Println("*c1.next", *c1.next)
	// c1.next.base.sayHi()
	c1.next.sayHi()
	c1.sayOK()
}
