package utils

import "fmt"

// PprintArray pretty prints an array of integers
func PprintArray(arr []int) {
	for i := 0; i < len(arr); i++ {
		fmt.Printf("%d ", arr[i])
	}
}
