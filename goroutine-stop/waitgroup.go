package main

import (
	"time"
	"fmt"
	"sync"
	"math/rand"
)

// START OMIT
func doWork() {
	time.Sleep(time.Duration(rand.Intn(1000))*time.Millisecond)
	fmt.Println("work")
}

func main() {
	var wg sync.WaitGroup // HL
	wg.Add(1) // HL
		go func() {
			doWork()
			wg.Done() // HL
		}()
	wg.Wait() // HL
	fmt.Println("done")
}
// END OMIT