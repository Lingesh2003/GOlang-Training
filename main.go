package main
import (
	"fmt"
	"time"
	"GOlang-Training/calculator"
)

func main() {
	go append()
	sum()
}

func append() {
	for i := 0; i < 100; i++ {
		calculator.InsertNumber(i)
		fmt.Println("Inserted : ", i)
		time.Sleep(1 * time.Second)
	}
}

func sum() {
	for i := 0; i < 100; i++ {
		sum := calculator.GetSum()
		fmt.Println("Sum : ", sum)
		time.Sleep(5 * time.Second)
	}
}