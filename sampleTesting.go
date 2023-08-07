package main

import "fmt"

func addSum(num1, num2 int) int {
	return num1 + num2
}

func sampleTesting() {
	var num1 ,num2 int
	fmt.Scanln(&num1)
	fmt.Scanln(&num2)
	var sum int = addSum(num1, num2)
	fmt.Println(sum)
}