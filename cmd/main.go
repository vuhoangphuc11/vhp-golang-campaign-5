package main

import (
	"fmt"
	"github.com/vuhoangphuc11/vhp-golang-campaign-5/pkg"
)

func main() {

	//VHP print channel value and stop if channel is closed
	//pkg.PrintChannel()

	//VHP select channel
	//pkg.SelectChannel()

	//c := make(chan int)
	//quit := make(chan int)
	//go func() {
	//	for i := 0; i < 10; i++ {
	//		fmt.Println(<-c)
	//	}
	//	quit <- 0
	//}()
	//
	//pkg.Fibonacci(c, quit)

	myChan := make(chan int)

	go pkg.Sender(myChan, "S1")
	go pkg.Sender(myChan, "S2")
	go pkg.Sender(myChan, "S3")

	start := 0

	for {
		start += <-myChan
		fmt.Println(start)

		if start >= 300 {
			break
		}
	}

}
