package main

import (
	"time"
	"fmt"
	"sync"
	"math/rand"
)

// START OMIT
var wg sync.WaitGroup // HL

func doWork(id int) {
	time.Sleep(time.Duration(rand.Intn(1000))*time.Millisecond)
	fmt.Printf("[%d] work\n", id)
}

func main() {
	wg.Add(3) // HL
	for i:= 0; i<3; i++ {
		go func(id int) {
			doWork(id)
			wg.Done() // HL
		}(i)
	}

	wg.Wait() // HL
	fmt.Println("done")
}
// END OMIT