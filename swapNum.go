package main

import "fmt"

func SwapNum(num1, num2 int) (int, int) {
	temp := num1
	num1 = num2
	num2 = temp
	return num1, num2
}

func swapNumbers() {
	num1 := 1
	num2 := 4
	fmt.Println(num1, num2)

	fmt.Println("Swapping Values....")
	sNum1, sNum2 := SwapNum(num1, num2)
	fmt.Println(sNum1, sNum2)
}
