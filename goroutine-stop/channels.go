package main

import "fmt"

// START OMIT
func worker1(c chan<- string) {
	c <- "me first!"
}

func worker2(c chan<- string) {
	c <- "me me!"
}

func main() {
	messages := make(chan string)
	go worker1(messages)
	go worker2(messages)

	fmt.Println(<-messages)
	fmt.Println(<-messages)
	close(messages)
}

// END OMIT
