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
		time.Sleep(1 * time.Second)
		fmt.Println("Inserted : ", i)
	}
}

func sum() {
	for i := 0; i < 100; i++ {
		sum := calculator.GetSum()
		time.Sleep(5 * time.Second)
		fmt.Println("Sum : ", sum)
	}
}