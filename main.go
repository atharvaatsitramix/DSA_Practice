package main

import "fmt"

func main() {
	fmt.Println("Welcome to DSA Practice!")
	// sayHello()
	nums2 := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	k2 := removeDuplicates(nums2)
	fmt.Printf("Output: k = %d, nums = %v\n", k2, nums2[:k2])
}
