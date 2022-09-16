package main

import (
	"fmt"
	"github.com/vuhoangphuc11/vhp-golang-campaign-5/pkg"
	"time"
)

func main() {

	//VHP print channel value and stop if channel is closed
	pkg.PrintChannel()

	//VHP select channel
	pkg.SelectChannel()

	//VHP Fibonacci
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()

	pkg.Fibonacci(c, quit)

	//VHP Use 1 channel to listen to data from many places
	myChan1 := make(chan int)

	go pkg.Sender(myChan1, "S1")
	go pkg.Sender(myChan1, "S2")
	go pkg.Sender(myChan1, "S3")

	start := 0

	for {
		start += <-myChan1
		fmt.Println(start)

		if start >= 300 {
			break
		}
	}

	//VHP Load balancing with Channel
	myChan2 := pkg.Publisher()
	maxConsumer := 5

	for i := 1; i <= maxConsumer; i++ {
		go pkg.Consumer(myChan2, fmt.Sprintf("%d", i))
	}

	time.Sleep(time.Second * 10)

	//VHP merge channel using waitgroup and channel
	s := pkg.SumAllStreams(
		pkg.StreamNumbers(1, 2, 3, 4, 5),
		pkg.StreamNumbers(8, 8, 3, 3, 10, 12, 14),
		pkg.StreamNumbers(1, 1, 2, 2, 4, 4, 6),
	)

	fmt.Println(<-s)
}
