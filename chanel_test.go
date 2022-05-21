package golang_goroutine

import (
	"fmt"
	"testing"
)

func TestCreateChanel(t *testing.T) {
	ch := make(chan int)
	go func() {
		ch <- 1
	}()
	fmt.Println(<-ch)
	defer close(ch)
}
