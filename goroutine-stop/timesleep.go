package main

import (
	"time"
	"fmt"
	"math/rand"
)
// START OMIT
func doWork() {
	for {
		time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		fmt.Println("work")
	}
}

func main() {
	go doWork()
	time.Sleep(3 * time.Second)
	fmt.Println("done")
}
// END OMIT