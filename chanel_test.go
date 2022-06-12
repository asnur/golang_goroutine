package golang_goroutine

import (
	"fmt"
	"testing"
	"time"
)

func TestCreateChanel(t *testing.T) {
	ch := make(chan int)
	go func() {
		ch <- 1
	}()
	fmt.Println(<-ch)
	defer close(ch)
}

func GiveMeResponse(ch chan string) {
	time.Sleep(time.Second * 1)
	ch <- "Hello"
}

func TestChanelAsParameter(t *testing.T) {
	ch := make(chan string)
	defer close(ch)
	go GiveMeResponse(ch)
	fmt.Println(<-ch)
}

func OnlyIn(ch chan<- int) {
	time.Sleep(time.Second * 2)
	ch <- 1
}

func OnlyOut(ch <-chan int) {
	fmt.Println(<-ch)
}

func TestInOutChanel(t *testing.T) {
	ch := make(chan int)
	go OnlyIn(ch)
	go OnlyOut(ch)

	time.Sleep(time.Second * 3)
	close(ch)
}

func TestBufferedChanel(t *testing.T) {
	ch := make(chan int, 4)
	go func() {
		ch <- 1
		ch <- 2
		ch <- 3
		ch <- 4
	}()
	go func() {
		fmt.Println(<-ch)
		fmt.Println(<-ch)
		fmt.Println(<-ch)
		fmt.Println(<-ch)
	}()
	defer close(ch)
	time.Sleep(time.Second * 2)
}

func TestRangeChanel(t *testing.T) {
	ch := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
		close(ch)
	}()

	for i := range ch {
		fmt.Println(i)
	}
}

func TestRaceCondition(t *testing.T) {
	x := 0

	for i := 0; i < 1000; i++ {
		go func() {
			for j := 0; j < 100; j++ {
				x++
			}
		}()
	}
	time.Sleep(time.Second * 5)
	fmt.Println("Counter = ", x)
}
