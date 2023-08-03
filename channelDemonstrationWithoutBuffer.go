package main
import (
	"fmt"
	"time"
)

func mchannelDemonstrationWithoutBuffer() {
	ch := make(chan int, 10)
	go a(ch)

	for v := range ch {
		fmt.Println(v)
	}
}

func a(ch chan int) {
	for i := 0; i < 10; i++ {
		time.Sleep(1 * time.Second)
		ch <- 1
	}
	close(ch)
}