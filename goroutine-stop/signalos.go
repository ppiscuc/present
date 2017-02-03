package main

import (
	"time"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
)

func doWork(c chan bool) {
	for {
		select {
		case <- c:
			return
		default:
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
			fmt.Println("work")
		}
	}
}
// START OMIT
func main() {
	quit := make(chan bool)
	c := make(chan os.Signal, 1) //buffered or risk missing the signal if not ready to receive when signal is sent
	signal.Notify(c, syscall.SIGTERM, syscall.SIGINT, syscall.SIGKILL)

	go doWork(quit)
	fmt.Println("waiting for signal")
	s := <- c
	quit <- true
	fmt.Println("got signal:", s)
	fmt.Println("done")

}
// END OMIT