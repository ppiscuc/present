package main

import (
"time"
"fmt"
"math/rand"
)
// START OMIT
func doWork(c chan bool, teardown chan bool) {
	for {
		select {
		case <- c:
			fmt.Println("teardown")
			time.Sleep(1 * time.Second)
			fmt.Println("done teardown")
			teardown <- true
			return
		default:
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
			fmt.Println("work")
		}
	}
}

func main() {
	quit := make(chan bool)
	teardown := make(chan bool)
	go doWork(quit, teardown)
	time.Sleep(5 * time.Second)
	close(quit) // quit <- true
	<-teardown
	fmt.Println("done")
}
// END OMIT
