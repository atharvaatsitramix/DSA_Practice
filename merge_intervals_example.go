package main

import (
	"fmt"
	"sort"
)

// mergeIntervals merges all overlapping intervals in a given collection.
// Input: A collection of intervals represented as [][]int
// Output: A collection of non-overlapping intervals
func mergeIntervals(intervals [][]int) [][]int {
	if len(intervals) <= 1 {
		return intervals
	}

	// Sort intervals by start time
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	result := [][]int{intervals[0]}

	for i := 1; i < len(intervals); i++ {
		current := intervals[i]
		lastMerged := result[len(result)-1]

		// Check if current interval overlaps with the last merged interval
		if current[0] <= lastMerged[1] {
			// Merge intervals by updating the end time
			lastMerged[1] = maxInt(lastMerged[1], current[1])
		} else {
			// No overlap, add current interval to result
			result = append(result, current)
		}
	}

	return result
}

// maxInt returns the maximum of two integers.
func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func runMergeIntervalsExample() {
	fmt.Println("=== Merge Intervals Algorithm Example ===")

	// Example 1
	intervals1 := [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}
	fmt.Printf("Input: %v\n", intervals1)
	merged1 := mergeIntervals(intervals1)
	fmt.Printf("Output: %v\n", merged1)
	fmt.Println("Explanation: [1,3] and [2,6] overlap, so they merge to [1,6]")
	fmt.Println()

	// Example 2
	intervals2 := [][]int{{1, 4}, {4, 5}}
	fmt.Printf("Input: %v\n", intervals2)
	merged2 := mergeIntervals(intervals2)
	fmt.Printf("Output: %v\n", merged2)
	fmt.Println("Explanation: [1,4] and [4,5] overlap at point 4, so they merge to [1,5]")
	fmt.Println()

	// Example 3
	intervals3 := [][]int{{1, 4}, {0, 4}}
	fmt.Printf("Input: %v\n", intervals3)
	merged3 := mergeIntervals(intervals3)
	fmt.Printf("Output: %v\n", merged3)
	fmt.Println("Explanation: After sorting: [0,4] and [1,4], they overlap and merge to [0,4]")
	fmt.Println()

	fmt.Println("Algorithm Steps:")
	fmt.Println("1. Sort intervals by start time")
	fmt.Println("2. Initialize result with first interval")
	fmt.Println("3. For each subsequent interval:")
	fmt.Println("   - If it overlaps with last interval in result, merge them")
	fmt.Println("   - Otherwise, add it as a new interval")
	fmt.Println("4. Return the merged intervals")
}
