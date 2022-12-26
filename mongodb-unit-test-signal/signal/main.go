package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	/** Exit **/
	// {
	// 	defer fmt.Println("bye")
	// 	// os.Exit(0)
	// 	os.Exit(9) // Read https://tldp.org/LDP/abs/html/exitcodes.html
	// }

	/** Signal **/
	{
		sigs := make(chan os.Signal)                         // Use to get signals on it
		done := make(chan bool)                              // Use to shutdown the app
		signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM) // Relay incoming/provided signals to "sigs"
		go func() {
			<-sigs // Blocking until os.Signal is received
			fmt.Println("\nDoing graceful shutdown")
			done <- true // Send message to shutdown
		}()
		fmt.Println("Waiting for os.Signal")
		<-done // Blocking until internal shutdown message is received
		fmt.Println("Exit!")
	}
}
