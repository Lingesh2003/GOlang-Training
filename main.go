package main
import "fmt"

func main() {
	var num int
	for {
		fmt.Scanln(&num)
		InsertNumbers(num)
	}
}

func InsertNumbers(num int) {
	arr := []int{}
	arr = append(arr, num)
	theSum := GetAddition(arr)
	fmt.Println(theSum)
}

func GetAddition(arr []int) int {
	sum := 0
	for _, v := range arr {
		sum += v
	}
	return sum
}
