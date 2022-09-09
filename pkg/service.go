package pkg

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

func PrintChannel() {
	r := make(chan int)

	go func() {
		for i := 1; i <= 10; i++ {
			r <- i
		}
		defer close(r)
	}()

	for v := range r {
		fmt.Println("value of channel", v)
	}
}

func SelectChannel() {

	r := rand.New(rand.NewSource(time.Now().Unix()))

	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(time.Second * time.Duration(r.Intn(5)))
		ch1 <- "Hello Golang"
	}()

	go func() {
		time.Sleep(time.Second * time.Duration(r.Intn(5)))
		ch2 <- "Golang Hello"
	}()

	select {
	case v1 := <-ch1:
		fmt.Println(v1)
	case v2 := <-ch2:
		fmt.Println(v2)
	default:
		fmt.Println("Not found")
	}

}

func Fibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, y+x
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func Sender(c chan<- int, name string) {
	for i := 1; i <= 100; i++ {
		c <- 1
		fmt.Printf("%s has sent 1 to channel\n", name)
		runtime.Gosched()
	}
}

