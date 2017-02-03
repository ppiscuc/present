package main



import (
"time"
"fmt"
"math/rand"
)

func expensiveTeardown() {
	time.Sleep(1 * time.Second)
	fmt.Println("done teardown")
}
// START OMIT
func doWork(c <-chan bool, teardown chan<- bool) {
	for {
		select {
		case <- c:
			expensiveTeardown()
			teardown <- true // HL
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
	time.Sleep(3 * time.Second)
	close(quit) // quit <- true
	<-teardown // HL
	fmt.Println("done")
}
// END OMIT
