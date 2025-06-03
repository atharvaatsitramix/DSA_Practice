package main

import "fmt"

// maxSumSubarray finds the maximum sum of any contiguous subarray of size K.
func maxSumSubarray(arr []int, k int) int {
	n := len(arr)
	if n < k {
		fmt.Println("Invalid input: array length is less than k")
		return -1
	}

	// Compute the sum of the first window
	maxSum := 0
	for i := 0; i < k; i++ {
		maxSum += arr[i]
	}

	// Slide the window over the array
	windowSum := maxSum
	for i := k; i < n; i++ {
		windowSum += arr[i] - arr[i-k]
		if windowSum > maxSum {
			maxSum = windowSum
		}
	}

	return maxSum
}

func runSlidingWindowExample() {
	arr := []int{1, 4, 2, 10, 23, 3, 1, 0, 20}
	k := 4
	fmt.Printf("Maximum sum of a subarray of size %d is %d\n", k, maxSumSubarray(arr, k))
}
