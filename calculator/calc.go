package calculator
import "fmt"

func Sum(a, b int) int {
	return (a + b)
}

type Num struct {
	A int
	B int
	Action []string
}

func EvalUsingStruct(val Num) int {
	for i := 0; i < len(val.Action); i++ {
		if (val.Action[i] == "add") {
			fmt.Println(val.A + val.B)
		} else {
			fmt.Println(val.A * val.B)
		}
	}
	return 0
}


var numList = []int{}

func InsertNumber(n int) {
	numList = append(numList, n)
}

func GetSum() int {
	sum := 0
	for _, v := range numList {
		sum += v
	}
	return sum
}