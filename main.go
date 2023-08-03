package main
import (
	"fmt"
	"time"
	"GOlang-Training/calculator"
)

func main() {
	for i := 0; i < 100; i++ {
		calculator.InsertNumber()
		time.Sleep(1 * time.Second)
		fmt.Println("Inserted : ", i)
	}
	
	for i := 0; i< 100; i++ {
		sum := calculator.GetSum()
		time.Sleep(5 * time.Second)
		fmt.Println("Sum : ", sum)
	}
}

