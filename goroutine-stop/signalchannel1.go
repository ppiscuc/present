package main

import (
	"fmt"
	"math/rand"
	"time"
)

// START OMIT
func doWork(id int, c <-chan bool) {
	for {
		select {
		case <-c:
			return
		default:
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
			fmt.Printf("[%d] work\n", id)
		}
	}
}

func main() {
	quit := make(chan bool)
	for i := 0; i < 3; i++ {
		go doWork(i, quit)
	}
	time.Sleep(3 * time.Second)
	close(quit) // quit <- true
	fmt.Println("done")
}

// END OMIT
