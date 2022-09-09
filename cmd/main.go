package main

import (
	"fmt"
	"github.com/vuhoangphuc11/vhp-golang-campaign-5/pkg"
)

func main() {

	//VHP print channel value and stop if channel is closed
	//pkg.PrintChannel()

	//pkg.SelectChannel()

	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	pkg.Fibonacci(c, quit)

}
