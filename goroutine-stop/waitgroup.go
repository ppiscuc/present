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
	var wg sync.WaitGroup
	wg.Add(1)
		go func() {
			doWork()
			wg.Done()
		}()
	wg.Wait()
	fmt.Println("done")
}
// END OMIT