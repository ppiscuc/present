package main

import (
	"time"
	"fmt"
	"math/rand"
	"context"
)
// START OMIT
func doWork(ctx context.Context) {
	for {
		select {
		case <- ctx.Done(): // HL
			return
		default:
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
			fmt.Println("work")
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background()) // HL
	go doWork(ctx)
	time.Sleep(5 * time.Second)
	cancel() // HL
	fmt.Println("done")
}
// END OMIT