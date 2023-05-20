package main

import (
	"fmt"
	"os"
	"sync"
	"time"
)

type Car struct {
	Model string
}

func Select(receive chan int, quites chan struct{}) {

	for {
		time.Sleep(time.Second)

		select {

		case <-receive:
			fmt.Println("received")

		case <-quites:
			fmt.Println("Quit")
			os.Exit(0)
			break
		}

	}
}

func main() {

	receive := make(chan int)
	quites := make(chan struct{})

	go func() {
		receive <- 1
		quites <- struct{}{}
	}()

	go Select(receive, quites)

	select {}

}

// The pinger print a ping and wait for ponger channel
func pinger(pinger <-chan int, ponger chan<- int) {

	for {

		<-pinger
		fmt.Println("ping")
		time.Sleep(time.Second)
		ponger <- 1

	}
}

// The ponger prints a pong and wait for pinger channel
func ponger(pinger chan<- int, ponger <-chan int) {

	for {
		<-ponger
		fmt.Println("pong")
		time.Sleep(time.Second)
		pinger <- 1
	}
}

func pingPongGame() {

	ping := make(chan int)
	pong := make(chan int)

	go pinger(ping, pong)
	go ponger(ping, pong)

	ping <- 1

	select {}
}

func heavy() {

	for {
		time.Sleep(time.Second * 1)
		fmt.Println("Heavy")
	}

}

func superHeavy() {
	time.Sleep(time.Second * 2)
	fmt.Println("Super Heavy")
}

func impWaitGroup() {

	//Wait groups
	wg := &sync.WaitGroup{}
	wg.Add(2)
	go func() {
		fmt.Println("Database")

		wg.Done()
	}()

	go func() {
		fmt.Println("Second Anonyouns function")
		wg.Done()
	}()

	fmt.Println("Finish after 5 min")

	wg.Wait()
}

func implMutex() {

	//Mutex
	//mut := &sync.Mutex{}
	//mut.Lock()
	//mut.Unlock()
}

func unBufferedChannel() {

	c := make(chan int)

	//Sending in goroutine
	go func() {
		c <- 1
	}()

	//sniffing out value from channel
	val := <-c

	//print Value
	fmt.Println(val)

	//Sleep for 2 sec
	time.Sleep(time.Second * 2)

	//Sending in goroutine
	go func() {
		c <- 2
	}()

	//Sniffing from chanel. variable val as already been initialize, just assign a value to it
	val = <-c

	fmt.Println(val)
}

func bufferedChannel() {

	c := make(chan *Car, 3)

	go func() {
		c <- &Car{"1"}
		c <- &Car{"2"}
		c <- &Car{"3"}
		c <- &Car{"4"}
		close(c)
	}()

	for i := range c {

		fmt.Println(i.Model)
	}
}
