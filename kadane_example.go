package main

import "fmt"

// maxSubArray finds the maximum sum of a contiguous subarray using Kadane's Algorithm.
func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	maxSoFar := nums[0]
	maxEndingHere := nums[0]

	for i := 1; i < len(nums); i++ {
		maxEndingHere = max(nums[i], maxEndingHere+nums[i])
		maxSoFar = max(maxSoFar, maxEndingHere)
	}
	return maxSoFar
}

// max returns the maximum of two integers.
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func runKadaneExample() {
	arr := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	fmt.Printf("Maximum subarray sum is %d\n", maxSubArray(arr))
}
