package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func syncprint(i int) {
	time.Sleep(time.Second * 5)
	fmt.Println(i)
}

func main() {
	c := make(chan os.Signal)
	stop := make(chan bool, 1)
	stop <- false
	var i int
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	var wg sync.WaitGroup

	go func() {
		<-c
		stop <- true
		fmt.Println("stop:", i)
		wg.Wait()
		os.Exit(1)
	}()

	for i = 0; i < 10; i++ {
		time.Sleep(time.Second * 1)
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			syncprint(i)
		}(i)
		if <-stop {
			break
		} else {
			stop <- false
		}
	}

	wg.Wait()
}
