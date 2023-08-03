package main
import (
	"fmt"
)

func main() {
	ch := make(chan int)
	go goRoutine(ch)
	ch <- 10
}

func goRoutine(ch chan int) {
	fmt.Println("Channel Value:", <-ch)
}