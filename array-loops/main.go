package main

import (
	"fmt"
)

func main() {
	/** Array & Loops **/
	// var arr [6]string
	// arr = [6]string{"Jan", "Jack", "Hady", "Sara"}
	arr := [6]string{"Jan", "Jack", "Hady", "Sara"}
	fmt.Println("arr:", arr)
	fmt.Println("arrLen:", len(arr))
	// var num = [...]int{1, 2, 3, 4, 5, 6}
	num := [...]int{1, 2, 3, 4, 5, 6}
	fmt.Println("num:", num)
	fmt.Println("numLen:", len(num))

	XoBoard := [3][3]string{}
	/** Infinite loop **/
	for {
		fmt.Println("XoBoard:", XoBoard)
		break // return
	}
	XoBoard = [3][3]string{
		{"-", "-", "x"},
		{"-", "o", "-"},
		{"x", "o", "-"},
	}
	for i := 0; i < len(XoBoard); i++ {
		fmt.Println(i, XoBoard[i])
	}
	for i, item := range XoBoard {
		fmt.Println(i, item)
	}
	fmt.Println("Skip the round")
	for x := 0; x < 3; x++ {
		for y := 0; y < 3; y++ {
			if x == y && y == 1 {
				break
			}
			fmt.Println("x, y:", x, y)
		}
	}
	fmt.Println("Break entirely")
breakpoint:
	for x := 0; x < 3; x++ {
		for y := 0; y < 3; y++ {
			if x == y && y == 1 {
				break breakpoint
			}
			fmt.Println("x, y:", x, y)
		}
	}
}
