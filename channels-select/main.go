package main

import (
	"fmt"
	"sync"
	"time"
)

func sendMessage(c chan string, s string) {
	time.Sleep(500 * time.Millisecond)
	c <- "Message " + s + " sent"
}

func main() {
	/** Receive One Message From Channel Directly **/
	// c1 := make(chan string)
	// go sendMessage(c1, "Hello")
	// // result := <-c1
	// // fmt.Println("result:", result)
	// fmt.Println("result:", <-c1)

	/** Receive Multiple Messages From Channel Directly **/
	// c2 := make(chan string, 2)
	// c2 <- "Hello from channel 2"
	// c2 <- "Golang channel 2"
	// fmt.Println("Do something")
	// fmt.Println(<-c2)
	// fmt.Println(<-c2)

	/** Receive Multiple Messages From Channel By Range Looping **/
	// c3 := make(chan string, 2)
	// c3 <- "Hi from channel 3"
	// c3 <- "Golang channel 3"
	// // close(c3) // Close channel 3 immediately
	// go func() {
	// 	time.Sleep(time.Second * 2) // Delay 2 seconds before closing channel 3
	// 	close(c3)
	// }()
	// fmt.Println("Do something")
	// for c3msg := range c3 {
	// 	fmt.Println("c3msg:", c3msg)
	// }

	/** Receive Multiple Messages From Channels By Select **/
	// 	c4 := make(chan string, 2)
	// 	c5 := make(chan string, 2)
	// 	c4 <- "Golang channel 4"
	// 	c4 <- "Hello from channel 4"
	// 	c5 <- "Golang channel 5"
	// 	c5 <- "Hello from channel 5"
	// breakpoint:
	// 	for {
	// 		select {
	// 		case msg4 := <-c4:
	// 			fmt.Println("received", msg4)
	// 			break
	// 		case msg5 := <-c5:
	// 			fmt.Println("received", msg5)
	// 			break
	// 		default:
	// 			break breakpoint
	// 		}
	// 	}

	// 	c4 := make(chan string)
	// 	c5 := make(chan string)
	// 	go sendMessage(c4, "Golang channel 4")
	// 	go sendMessage(c5, "Golang channel 5")
	// breakpoint:
	// 	for {
	// 		select {
	// 		case msg4 := <-c4:
	// 			fmt.Println("received", msg4)
	// 			break
	// 		case msg5 := <-c5:
	// 			fmt.Println("received", msg5)
	// 			break
	// 		case <-time.After(3 * time.Second):
	// 			break breakpoint
	// 		}
	// 	}

	/** Play Ping-Pong Using Wait Group **/
	var wg sync.WaitGroup
	pingCh := make(chan bool)
	pongCh := make(chan bool)
	go playPing(pingCh, pongCh, &wg)
	go playPong(pingCh, pongCh, &wg)
	pongCh <- true
	wg.Wait()
}

func playPing(pingCh chan<- bool, pongChan <-chan bool, wg *sync.WaitGroup) {
	for {
		wg.Add(1)
		defer wg.Done()
		time.Sleep(500 * time.Millisecond)
		<-pongChan // Wait for pong
		fmt.Println("Ping!")
		pingCh <- true
	}
}

func playPong(pingCh <-chan bool, pongChan chan<- bool, wg *sync.WaitGroup) {
	for {
		wg.Add(1)
		defer wg.Done()
		time.Sleep(500 * time.Millisecond)
		<-pingCh // Wait for ping
		fmt.Println("Pong!")
		pongChan <- true
	}
}
