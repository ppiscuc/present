package main

import (
	"time"
	"fmt"
	"sync"
	"math/rand"
)

// START OMIT
var wg sync.WaitGroup

func doWork(id int) {
	time.Sleep(time.Duration(rand.Intn(1000))*time.Millisecond)
	fmt.Printf("[%d] work\n", id)
}

func main() {
	wg.Add(3)
	for i:= 0; i<3; i++ {
		go func(id int) {
			doWork(id)
			wg.Done()
		}(i)
	}

	wg.Wait()
	fmt.Println("done")
}
// END OMIT