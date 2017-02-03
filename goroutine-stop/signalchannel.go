package main

import (
	"time"
	"fmt"
	"math/rand"
)
// START OMIT
func doWork(c chan bool) {
	for {
		select {
		case <- c: // HL
		return // HL
		default:
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
			fmt.Println("work")
		}
	}
}

func main() {
	quit := make(chan bool)
	go doWork(quit)
	time.Sleep(5 * time.Second)
	close(quit) // quit <- true
	fmt.Println("done")
}
// END OMIT
