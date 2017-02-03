package main

import "fmt"

// START OMIT
func ping(c chan string) {
	c <- "ping"
}

func pong(c chan string) {
	c <- "pong"
}

func main() {
	messages := make(chan string)
	go ping(messages)
	go pong(messages)

	fmt.Println(<-messages)
	fmt.Println(<-messages)
	close(messages)
}
// END OMIT