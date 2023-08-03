package main
import "fmt"

type Num struct {
	numArr []int
	result int
}

func arrayAdder() {
	num1 := Num{}
	num1.numArr = []int{1, 2, 3, 4, 5}
	num1.adder()
	fmt.Println(num1.result)
}

func (evalNum *Num) adder() {
	for i := 0; i < len(evalNum.numArr); i++ {
		evalNum.result += evalNum.numArr[i]
	}
}
