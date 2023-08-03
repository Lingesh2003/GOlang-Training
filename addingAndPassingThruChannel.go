package main
import (
	"fmt"
)

/* Create a Slice in main 
Divide it in 2 parts
Send it to 2 different GOroutines from main
Get the individual sum from these
GOroutine and show the final addition in main */

func addingAndPassingThruChannel() {
	arr := []int {1,6,2,7,3,8,4,9,5,10}
	fmt.Println("Initial Array: ", arr)

	slice1 := arr[:5]
	slice2 := arr[5:]
	fmt.Println(slice1)
	fmt.Println(slice2)

	ch := make(chan int)

	go add(slice1, ch)
	x := <- ch
	go add(slice2, ch)
	y := <- ch

	fmt.Println(x + y)	
}

func add(arr []int, ch chan int) {
	sum := 0
	for _, v := range arr {
		sum += v
	}
	ch <- sum
}