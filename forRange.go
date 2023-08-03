package main
import "fmt"

func forRange() {
	arr := []int {1,2,3,4}

	for _, v := range arr {
		fmt.Println(v)
	}
}