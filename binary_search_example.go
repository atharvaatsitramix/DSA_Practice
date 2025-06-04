package main

import "fmt"

// binarySearch performs binary search on a sorted array
// Returns the index of target if found, -1 otherwise
func binarySearch(arr []int, target int) int {
	left := 0
	right := len(arr) - 1

	for left <= right {
		mid := left + (right-left)/2

		if arr[mid] == target {
			return mid // Target found
		} else if arr[mid] < target {
			left = mid + 1 // Search right half
		} else {
			right = mid - 1 // Search left half
		}
	}

	return -1 // Target not found
}

// binarySearchVerbose performs binary search with step-by-step output
func binarySearchVerbose(arr []int, target int) int {
	left := 0
	right := len(arr) - 1
	step := 1

	fmt.Printf("Searching for target: %d in array: %v\n", target, arr)
	fmt.Printf("Initial: left=%d, right=%d\n", left, right)

	for left <= right {
		mid := left + (right-left)/2
		fmt.Printf("\nStep %d:\n", step)
		fmt.Printf("  left=%d, right=%d, mid=%d\n", left, right, mid)
		fmt.Printf("  arr[%d] = %d\n", mid, arr[mid])

		if arr[mid] == target {
			fmt.Printf("  Target found at index %d!\n", mid)
			return mid
		} else if arr[mid] < target {
			fmt.Printf("  %d < %d, search right half\n", arr[mid], target)
			left = mid + 1
		} else {
			fmt.Printf("  %d > %d, search left half\n", arr[mid], target)
			right = mid - 1
		}
		step++
	}

	fmt.Printf("\nTarget %d not found in the array\n", target)
	return -1
}

func runBinarySearchExample() {
	fmt.Println("=== Binary Search Algorithm Example ===")
	fmt.Println()

	arr := []int{1, 3, 5, 7, 9, 11, 13, 15}
	fmt.Printf("Sorted Array: %v\n", arr)
	fmt.Printf("Array Indices: [0, 1, 2, 3, 4, 5, 6, 7]\n")
	fmt.Println()

	// Example 1: Target found quickly
	fmt.Println("--- Example 1: Target found at middle ---")
	binarySearchVerbose(arr, 7)
	fmt.Println()

	// Example 2: Target found after multiple steps
	fmt.Println("--- Example 2: Target found after multiple steps ---")
	binarySearchVerbose(arr, 11)
	fmt.Println()

	// Example 3: Target not found
	fmt.Println("--- Example 3: Target not found ---")
	binarySearchVerbose(arr, 6)
	fmt.Println()

	// Test multiple targets
	fmt.Println("--- Multiple Test Cases ---")
	targets := []int{1, 3, 5, 7, 9, 11, 13, 15, 0, 2, 16}

	for _, target := range targets {
		index := binarySearch(arr, target)
		if index != -1 {
			fmt.Printf("Target %2d: Found at index %d\n", target, index)
		} else {
			fmt.Printf("Target %2d: Not found\n", target)
		}
	}

	fmt.Println()
	fmt.Println("=== Algorithm Analysis ===")
	fmt.Println("Time Complexity:")
	fmt.Println("  - Best Case: O(1) - target found at middle")
	fmt.Println("  - Average Case: O(log n)")
	fmt.Println("  - Worst Case: O(log n)")
	fmt.Println("Space Complexity: O(1) - iterative approach")
	fmt.Println()
	fmt.Println("Key Requirements:")
	fmt.Println("  - Array must be sorted")
	fmt.Println("  - Random access to elements (arrays work, linked lists don't)")
	fmt.Println()
	fmt.Println("Algorithm Steps:")
	fmt.Println("  1. Initialize left=0, right=n-1")
	fmt.Println("  2. Calculate mid = left + (right-left)/2")
	fmt.Println("  3. Compare arr[mid] with target:")
	fmt.Println("     - If equal: target found")
	fmt.Println("     - If arr[mid] < target: search right half (left = mid+1)")
	fmt.Println("     - If arr[mid] > target: search left half (right = mid-1)")
	fmt.Println("  4. Repeat until found or left > right")
}
