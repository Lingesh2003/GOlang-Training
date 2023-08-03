package main
import "GOlang-Training/calculator"

func evaluateUsingStruct() {
	num := calculator.Num {}
	num.A = 1
	num.B = 4
	num.Action = []string {"add", "mul"}
	calculator.EvalUsingStruct(num)
}
