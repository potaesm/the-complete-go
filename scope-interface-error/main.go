package main

import (
	"errors"
	"fmt"
)

/** Interface **/
type walker interface {
	walk(p point) error
	getPosition() point
}

type talker interface {
	talk(s string)
}

type point struct {
	x int
	y int
}

type human struct {
	position point
}

func (h *human) walk(p point) error {
	if p.x < 0 || p.y < 0 {
		return errors.New("Invalid point")
	}
	h.position = p
	fmt.Println("Human walked to", h.position)
	return nil
}

func (h *human) getPosition() point {
	return h.position
}

func (h *human) talk(s string) {
	fmt.Println("Human talking", s)
}

type animal struct {
	position point
}

func (a *animal) walk(p point) error {
	if p.x < 0 || p.y < 0 {
		return errors.New("Invalid point")
	}
	a.position = p
	fmt.Println("Animal walked to", a.position)
	return nil
}

func (a *animal) getPosition() point {
	return a.position
}

func move(w walker, points []point) error {
	for _, p := range points {
		err := w.walk(p)
		if err != nil {
			return err
		}
	}
	return nil
}

func moveHuman(w *human, points []point) error {
	for _, p := range points {
		err := w.walk(p)
		if err != nil {
			return err
		}
	}
	return nil
}

func moveAnimal(w *animal, points []point) error {
	for _, p := range points {
		err := w.walk(p)
		if err != nil {
			return err
		}
	}
	return nil
}

/** Error **/
// error is an interface that have Error()
// type error interface {
// 	Error() string
// }

func getTestError(b bool) (err error) {
	if b == true {
		// err = fmt.Errorf("this is an error :(")
		err = errors.New("this is an error :(")
	}
	return
}

/** Validator **/
var errInvalidlength = errors.New("Invalid length")

func inValidLength(s string) error {
	if len(s) < 1 {
		return errInvalidlength
	} else if s == "error" {
		return errors.New("Invalid input")
	}
	return nil
}

type invalidInput struct {
	filedName string
}

func (i invalidInput) Error() string {
	return fmt.Sprintf("Invalid input on field '%v'", i.filedName)
}

func validateInput(field, value string) (err error) {
	if value == "" {
		err = invalidInput{filedName: field}
	}
	return
}

/** Struct Scope **/
type PublicType struct{}

func simpleFunc() {
	type PrivateType struct{}
	t := PrivateType{}
	_ = t
}

/** File Scope Variable **/
var name = "Suthinan"

func main() {
	fmt.Println("name:", name)

	/** Function Scope Variable **/
	a := 1
	{ // Scope
		a = 2 // Reassign
		fmt.Println("a inner:", a)
	}
	fmt.Println("a outer:", a)
	b := 1
	{ // Scope
		b := 2 // Redefine
		fmt.Println("b inner:", b)
	}
	fmt.Println("b outer:", b)

	/** Struct Scope **/
	simpleFunc()
	pubt := PublicType{}
	fmt.Printf("%+v\n", pubt)
	// prit := PrivateType{}
	// fmt.Printf("%+v\n", prit)

	/** Interface **/
	potaesm := &human{}
	dog := &animal{}
	steps := []point{
		{x: 1, y: 1},
		{x: 2, y: 2},
		{x: 3, y: 3},
	}
	moveHuman(potaesm, steps)   // Can be used only for human
	moveAnimal(dog, steps)      // Can be used only for animal
	err := move(potaesm, steps) // Can be used both human and animal that implement walk
	if err != nil {
		fmt.Println(err)
	}
	err = move(dog, steps)
	if err != nil {
		fmt.Println(err)
	}
	err = move(dog, []point{{x: -1, y: 2}})
	if err != nil {
		fmt.Println(err)
	}

	/** Error **/
	{
		err := getTestError(true)
		if err != nil {
			fmt.Println(err.Error())
		}
	}
	{
		err := validateInput("firstname", "Suthinan")
		if err != nil {
			fmt.Println(err.Error())
		}
	}
	{
		err := validateInput("lastname", "")
		if err != nil {
			fmt.Println(err.Error(), "as &invalidInput{} :", errors.As(err, &invalidInput{}))
		}
	}
	{
		err := inValidLength("")
		fmt.Println(err.Error(), "is", errInvalidlength.Error(), ":", errors.Is(err, errInvalidlength))
	}
}
