package main

import (
	"fmt"
	"time"
)

func saveToDatabase(s string, c chan<- bool) {
	time.Sleep(2 * time.Second)
	c <- true
}

func sendMessage(c chan string, s string) {
	for {
		time.Sleep(500 * time.Millisecond)
		c <- s
	}
}

func echo(done <-chan bool, interval time.Duration) {
	ticker := time.NewTicker(interval)
	for {
		select {
		case T := <-ticker.C:
			fmt.Println("Ticker ticked", T.Second())
		case <-done:
			ticker.Stop()
			return
		}
	}
}

func main() {
	/** Timeout **/
	// c := make(chan bool)
	// go saveToDatabase("Hello", c)
	// for {
	// 	select {
	// 	case <-c:
	// 		fmt.Println("Saved")
	// 		return
	// 	case <-time.After(3 * time.Second):
	// 		fmt.Println("Timeout!")
	// 		return
	// 	}
	// }

	/** Timer 1 **/
	// timer1 := time.NewTimer(2 * time.Second)
	// <-timer1.C
	// fmt.Println("Timer 1 end!")

	/** Timer 2 **/
	// 	chan1 := make(chan string)
	// 	chan2 := make(chan string)
	// 	go sendMessage(chan1, "Hello")
	// 	go sendMessage(chan2, "Go")
	// 	timer2 := time.NewTimer(2 * time.Second)
	// breakpoint:
	// 	for {
	// 		select {
	// 		case msg1 := <-chan1:
	// 			fmt.Println("msg1:", msg1)
	// 		case msg2 := <-chan2:
	// 			fmt.Println("msg2:", msg2)
	// 		case <-timer2.C:
	// 			fmt.Println("Timer 2 end!")
	// 			break breakpoint
	// 		}
	// 	}

	/** Timer 3 **/
	// timer3 := time.NewTimer(time.Second)
	// go func() {
	// 	fmt.Println("Start timer 3")
	// 	<-timer3.C
	// 	fmt.Println("Timer 3 end!")
	// }()
	// stopped := timer3.Stop() // Immidiately stop timer after call go routine
	// if stopped {
	// 	fmt.Println("Timer 3 stopped!")
	// }
	// time.Sleep(3 * time.Second)

	/** Ticker 1 **/
	// ticker := time.NewTicker(1 * time.Second)
	// timer := time.NewTimer(5 * time.Second)
	// for {
	// 	select {
	// 	case T := <-ticker.C:
	// 		fmt.Println("Ticker ticked", T.Second())
	// 	// Do something
	// 	case <-timer.C:
	// 		ticker.Stop()
	// 		return
	// 	}
	// }

	/** Ticker 2 **/
	done := make(chan bool)
	go echo(done, time.Second*1)
	time.Sleep(5 * time.Second)
	done <- true
	fmt.Println("Ticker stopped")
}
