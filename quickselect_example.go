package main

import (
	"fmt"
	"math/rand"
	"time"
)

// ================================
// QUICKSELECT ALGORITHM
// ================================

// QuickSelect finds the k-th smallest element in an array (0-indexed)
// Time Complexity: Average O(n), Worst O(n²)
// Space Complexity: O(log n) for recursion, O(1) for iterative
func QuickSelect(arr []int, k int) int {
	if k < 0 || k >= len(arr) {
		panic("k is out of bounds")
	}

	// Work on a copy to avoid modifying original array
	nums := make([]int, len(arr))
	copy(nums, arr)

	return quickSelectRecursive(nums, 0, len(nums)-1, k)
}

// Recursive implementation of QuickSelect
func quickSelectRecursive(arr []int, left, right, k int) int {
	if left == right {
		return arr[left]
	}

	// Choose pivot and partition
	pivotIndex := partition(arr, left, right)

	if k == pivotIndex {
		return arr[k]
	} else if k < pivotIndex {
		return quickSelectRecursive(arr, left, pivotIndex-1, k)
	} else {
		return quickSelectRecursive(arr, pivotIndex+1, right, k)
	}
}

// QuickSelectIterative finds the k-th smallest element using iterative approach
func QuickSelectIterative(arr []int, k int) int {
	if k < 0 || k >= len(arr) {
		panic("k is out of bounds")
	}

	// Work on a copy
	nums := make([]int, len(arr))
	copy(nums, arr)

	left, right := 0, len(nums)-1

	for left <= right {
		pivotIndex := partition(nums, left, right)

		if k == pivotIndex {
			return nums[k]
		} else if k < pivotIndex {
			right = pivotIndex - 1
		} else {
			left = pivotIndex + 1
		}
	}

	return nums[k]
}

// partition rearranges array so that elements smaller than pivot are on left,
// larger elements are on right, and returns the final position of pivot
func partition(arr []int, left, right int) int {
	// Choose rightmost element as pivot
	pivot := arr[right]
	i := left

	for j := left; j < right; j++ {
		if arr[j] <= pivot {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}

	// Place pivot in correct position
	arr[i], arr[right] = arr[right], arr[i]
	return i
}

// ================================
// OPTIMIZED VERSIONS
// ================================

// QuickSelectRandomized uses random pivot selection for better average performance
func QuickSelectRandomized(arr []int, k int) int {
	if k < 0 || k >= len(arr) {
		panic("k is out of bounds")
	}

	nums := make([]int, len(arr))
	copy(nums, arr)

	rand.Seed(time.Now().UnixNano())
	return quickSelectRandomizedHelper(nums, 0, len(nums)-1, k)
}

func quickSelectRandomizedHelper(arr []int, left, right, k int) int {
	if left == right {
		return arr[left]
	}

	// Randomly choose pivot and move to end
	randomIndex := left + rand.Intn(right-left+1)
	arr[randomIndex], arr[right] = arr[right], arr[randomIndex]

	pivotIndex := partition(arr, left, right)

	if k == pivotIndex {
		return arr[k]
	} else if k < pivotIndex {
		return quickSelectRandomizedHelper(arr, left, pivotIndex-1, k)
	} else {
		return quickSelectRandomizedHelper(arr, pivotIndex+1, right, k)
	}
}

// QuickSelectMedianOfMedians uses median-of-medians for guaranteed O(n) worst-case
func QuickSelectMedianOfMedians(arr []int, k int) int {
	if k < 0 || k >= len(arr) {
		panic("k is out of bounds")
	}

	nums := make([]int, len(arr))
	copy(nums, arr)

	return quickSelectMOM(nums, 0, len(nums)-1, k)
}

func quickSelectMOM(arr []int, left, right, k int) int {
	if left == right {
		return arr[left]
	}

	// Use median of medians as pivot
	pivotValue := medianOfMedians(arr, left, right)

	// Find pivot index and move to end
	pivotIndex := left
	for i := left; i <= right; i++ {
		if arr[i] == pivotValue {
			pivotIndex = i
			break
		}
	}
	arr[pivotIndex], arr[right] = arr[right], arr[pivotIndex]

	pivotIndex = partition(arr, left, right)

	if k == pivotIndex {
		return arr[k]
	} else if k < pivotIndex {
		return quickSelectMOM(arr, left, pivotIndex-1, k)
	} else {
		return quickSelectMOM(arr, pivotIndex+1, right, k)
	}
}

// medianOfMedians finds a good pivot using median-of-medians algorithm
func medianOfMedians(arr []int, left, right int) int {
	n := right - left + 1
	if n <= 5 {
		// Base case: use insertion sort and return median
		temp := make([]int, n)
		copy(temp, arr[left:right+1])
		insertionSort(temp)
		return temp[n/2]
	}

	// Divide into groups of 5
	medians := []int{}
	for i := left; i <= right; i += 5 {
		groupRight := i + 4
		if groupRight > right {
			groupRight = right
		}

		temp := make([]int, groupRight-i+1)
		copy(temp, arr[i:groupRight+1])
		insertionSort(temp)
		medians = append(medians, temp[len(temp)/2])
	}

	// Recursively find median of medians
	return QuickSelectMedianOfMedians(medians, len(medians)/2)
}

func insertionSort(arr []int) {
	for i := 1; i < len(arr); i++ {
		key := arr[i]
		j := i - 1
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
	}
}

// ================================
// UTILITY FUNCTIONS
// ================================

// FindKthSmallest finds the k-th smallest element (1-indexed)
func FindKthSmallest(arr []int, k int) int {
	return QuickSelect(arr, k-1) // Convert to 0-indexed
}

// FindKthLargest finds the k-th largest element (1-indexed)
func FindKthLargest(arr []int, k int) int {
	return QuickSelect(arr, len(arr)-k) // Convert to k-th smallest from end
}

// FindMedian finds the median of an array
func FindMedian(arr []int) float64 {
	n := len(arr)
	if n%2 == 1 {
		return float64(QuickSelect(arr, n/2))
	} else {
		smaller := QuickSelect(arr, n/2-1)
		larger := QuickSelect(arr, n/2)
		return float64(smaller+larger) / 2.0
	}
}

// TopK finds the k smallest elements (not necessarily sorted)
func TopKSmallest(arr []int, k int) []int {
	if k <= 0 || k > len(arr) {
		return []int{}
	}

	nums := make([]int, len(arr))
	copy(nums, arr)

	// Partition array so that first k elements are the smallest
	quickSelectPartial(nums, 0, len(nums)-1, k-1)

	result := make([]int, k)
	copy(result, nums[:k])
	return result
}

func quickSelectPartial(arr []int, left, right, k int) {
	if left >= right {
		return
	}

	pivotIndex := partition(arr, left, right)

	if k < pivotIndex {
		quickSelectPartial(arr, left, pivotIndex-1, k)
	} else if k > pivotIndex {
		quickSelectPartial(arr, pivotIndex+1, right, k)
	}
	// If k == pivotIndex, we're done
}

// ================================
// DEMONSTRATION FUNCTIONS
// ================================

func DemoQuickSelect() {
	fmt.Println("=== QUICKSELECT ALGORITHM EXPLANATION ===\n")

	fmt.Println("QuickSelect is a selection algorithm to find the k-th smallest element")
	fmt.Println("in an unordered list. It's related to QuickSort but only recurses into")
	fmt.Println("one partition, making it more efficient for selection problems.\n")

	// Example 1: Basic QuickSelect
	fmt.Println("=== EXAMPLE 1: Basic QuickSelect ===")
	arr1 := []int{3, 6, 8, 10, 1, 2, 1}
	fmt.Printf("Array: %v\n", arr1)

	for k := 1; k <= len(arr1); k++ {
		kthSmallest := FindKthSmallest(arr1, k)
		fmt.Printf("%d-th smallest element: %d\n", k, kthSmallest)
	}
	fmt.Println()

	// Example 2: Find k-th largest
	fmt.Println("=== EXAMPLE 2: Find k-th Largest ===")
	arr2 := []int{7, 10, 4, 3, 20, 15}
	fmt.Printf("Array: %v\n", arr2)

	for k := 1; k <= 3; k++ {
		kthLargest := FindKthLargest(arr2, k)
		fmt.Printf("%d-th largest element: %d\n", k, kthLargest)
	}
	fmt.Println()

	// Example 3: Find median
	fmt.Println("=== EXAMPLE 3: Find Median ===")
	arr3 := []int{1, 5, 2, 8, 3, 9, 4}
	fmt.Printf("Array: %v\n", arr3)
	median := FindMedian(arr3)
	fmt.Printf("Median: %.1f\n\n", median)

	arr4 := []int{1, 2, 3, 4, 5, 6}
	fmt.Printf("Array: %v\n", arr4)
	median2 := FindMedian(arr4)
	fmt.Printf("Median: %.1f\n\n", median2)

	// Example 4: Top K elements
	fmt.Println("=== EXAMPLE 4: Top K Smallest Elements ===")
	arr5 := []int{9, 4, 5, 6, 7, 3, 1, 2}
	fmt.Printf("Array: %v\n", arr5)

	for k := 1; k <= 4; k++ {
		topK := TopKSmallest(arr5, k)
		fmt.Printf("Top %d smallest elements: %v\n", k, topK)
	}
	fmt.Println()

	// Example 5: Algorithm comparison
	fmt.Println("=== EXAMPLE 5: Algorithm Comparison ===")
	arr6 := []int{64, 34, 25, 12, 22, 11, 90}
	k := 3
	fmt.Printf("Array: %v\n", arr6)
	fmt.Printf("Finding %d-th smallest element:\n", k)

	result1 := QuickSelect(arr6, k-1)
	fmt.Printf("QuickSelect (basic): %d\n", result1)

	result2 := QuickSelectIterative(arr6, k-1)
	fmt.Printf("QuickSelect (iterative): %d\n", result2)

	result3 := QuickSelectRandomized(arr6, k-1)
	fmt.Printf("QuickSelect (randomized): %d\n", result3)

	result4 := QuickSelectMedianOfMedians(arr6, k-1)
	fmt.Printf("QuickSelect (median-of-medians): %d\n\n", result4)

	// Example 6: Step-by-step trace
	fmt.Println("=== EXAMPLE 6: Step-by-Step Trace ===")
	arr7 := []int{3, 6, 8, 10, 1, 2, 1}
	k = 3
	fmt.Printf("Finding %d-th smallest in: %v\n", k, arr7)
	fmt.Println("Step-by-step execution:")
	result := quickSelectTrace(arr7, k-1)
	fmt.Printf("Result: %d\n\n", result)

	// Performance characteristics
	fmt.Println("=== ALGORITHM CHARACTERISTICS ===")
	fmt.Println("Time Complexity:")
	fmt.Println("- Average case: O(n)")
	fmt.Println("- Worst case: O(n²) for basic version")
	fmt.Println("- Worst case: O(n) for median-of-medians version")
	fmt.Println()
	fmt.Println("Space Complexity:")
	fmt.Println("- Recursive: O(log n) average, O(n) worst case")
	fmt.Println("- Iterative: O(1)")
	fmt.Println()
	fmt.Println("Advantages:")
	fmt.Println("- Faster than full sorting for selection problems")
	fmt.Println("- In-place algorithm (with minor modifications)")
	fmt.Println("- Good average performance")
	fmt.Println()
	fmt.Println("Use Cases:")
	fmt.Println("- Finding median in streaming data")
	fmt.Println("- k-th order statistics")
	fmt.Println("- Top-k problems in competitive programming")
	fmt.Println("- Database query optimization")
}

// quickSelectTrace provides step-by-step tracing of the algorithm
func quickSelectTrace(arr []int, k int) int {
	nums := make([]int, len(arr))
	copy(nums, arr)

	return quickSelectTraceHelper(nums, 0, len(nums)-1, k, 1)
}

func quickSelectTraceHelper(arr []int, left, right, k, step int) int {
	fmt.Printf("Step %d: Array[%d:%d] = %v, looking for index %d\n",
		step, left, right, arr[left:right+1], k)

	if left == right {
		fmt.Printf("  Base case reached: arr[%d] = %d\n", left, arr[left])
		return arr[left]
	}

	pivot := arr[right]
	fmt.Printf("  Pivot = %d (arr[%d])\n", pivot, right)

	pivotIndex := partitionTrace(arr, left, right)
	fmt.Printf("  After partition: %v, pivot at index %d\n", arr[left:right+1], pivotIndex)

	if k == pivotIndex {
		fmt.Printf("  Found! arr[%d] = %d\n", k, arr[k])
		return arr[k]
	} else if k < pivotIndex {
		fmt.Printf("  Search left partition\n")
		return quickSelectTraceHelper(arr, left, pivotIndex-1, k, step+1)
	} else {
		fmt.Printf("  Search right partition\n")
		return quickSelectTraceHelper(arr, pivotIndex+1, right, k, step+1)
	}
}

func partitionTrace(arr []int, left, right int) int {
	pivot := arr[right]
	i := left

	for j := left; j < right; j++ {
		if arr[j] <= pivot {
			if i != j {
				arr[i], arr[j] = arr[j], arr[i]
			}
			i++
		}
	}

	arr[i], arr[right] = arr[right], arr[i]
	return i
}

// ================================
// PRACTICAL APPLICATIONS
// ================================

func DemoApplications() {
	fmt.Println("\n=== PRACTICAL APPLICATIONS ===")

	// Application 1: Finding salary percentiles
	fmt.Println("1. SALARY ANALYSIS")
	salaries := []int{45000, 52000, 48000, 65000, 58000, 72000, 41000, 55000, 62000, 70000}
	fmt.Printf("Employee salaries: %v\n", salaries)

	median := FindMedian(salaries)
	fmt.Printf("Median salary: $%.0f\n", median)

	p25 := FindKthSmallest(salaries, len(salaries)/4)
	p75 := FindKthSmallest(salaries, 3*len(salaries)/4)
	fmt.Printf("25th percentile: $%d\n", p25)
	fmt.Printf("75th percentile: $%d\n", p75)
	fmt.Println()

	// Application 2: Top performers
	fmt.Println("2. TOP PERFORMERS")
	scores := []int{87, 92, 78, 96, 89, 84, 91, 85, 93, 88}
	fmt.Printf("Test scores: %v\n", scores)

	top3 := []int{}
	for i := 1; i <= 3; i++ {
		score := FindKthLargest(scores, i)
		top3 = append(top3, score)
	}
	fmt.Printf("Top 3 scores: %v\n", top3)
	fmt.Println()

	// Application 3: Load balancing
	fmt.Println("3. LOAD BALANCING")
	serverLoads := []int{23, 45, 12, 67, 34, 56, 78, 29, 41, 52}
	fmt.Printf("Server loads: %v\n", serverLoads)

	medianLoad := FindMedian(serverLoads)
	fmt.Printf("Median load: %.1f\n", medianLoad)
	fmt.Printf("Servers below median load: ")

	for i, load := range serverLoads {
		if float64(load) < medianLoad {
			fmt.Printf("Server%d(%d) ", i+1, load)
		}
	}
	fmt.Println()
}
