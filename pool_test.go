package golang_goroutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestPool(t *testing.T) {
	pool := sync.Pool{}
	pool.New = func() interface{} {
		return "Hello"
	}
	fmt.Println(pool.Get())
	pool.Put("World")
	fmt.Println(pool.Get())
}

//Check Ignored Goroutine
func TestPoolIngored(t *testing.T) {
	pool := sync.Pool{}
	pool.Put("Hello")
	pool.Put("World")
	pool.Put("Asnur")

	for i := 0; i < 10; i++ {
		go func() {
			data := pool.Get()
			fmt.Println(data)
			time.Sleep(time.Second * 1)
			pool.Put(data)
		}()
	}

	time.Sleep(time.Second * 11)
	fmt.Print("Done")
}
