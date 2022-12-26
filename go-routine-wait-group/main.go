package main

import (
	"fmt"
	"sync"
	"time"
)

func syncFunc(n int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Start synchronous task %d\n", n)
	time.Sleep(3 * time.Second)
	fmt.Printf("Finish task %d\n", n)
}

func main() {
	/** Use Of Literal Function For Go Routine **/
	// for i := 0; i < 5; i++ {
	// 	go func(n int) {
	// 		fmt.Printf("Start synchronous task %d\n", n)
	// 		time.Sleep(time.Second)
	// 		fmt.Printf("Finish task %d\n", n)
	// 	}(i)
	// }
	// time.Sleep(time.Second * 2)

	/** Wait Group **/
	var wg sync.WaitGroup
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go syncFunc(i, &wg)
	}
	wg.Wait()

	fmt.Println("Main routine will exit now!")
}
