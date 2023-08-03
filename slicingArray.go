package main
import "fmt"

func slicingArray() {
	// Array
	arr := []int {1,2,3,4}
	fmt.Print(arr)
	
	// Slicing the Array
	slice := arr[0:3]
	fmt.Print(slice)

	fmt.Println(arr[0:2], cap(arr), len(arr))
	fmt.Println(slice, cap(slice), len(slice))

	//Appending a Number to the Alices Array
	slice = append(slice, 7)
	fmt.Println(slice)
}