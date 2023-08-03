package main
import "fmt"

type num struct {
	num1 int
	num2 int
	res  int
}

func sumUsingPointers() {
	n := num{1, 4, 0}
	n.sum()
	fmt.Println("Result: ", n.res)
}

func (n *num) sum() {
	 n.res = n.num1 + n.num2
}