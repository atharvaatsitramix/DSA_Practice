# DSA_Practice

## Algorithm Implementations

### Two-Pointer Algorithm

The two-pointer technique is used to remove duplicates from a sorted array in-place, ensuring each unique element appears only once. The implementation involves:

- **Pointers**: Use two pointers, `i` and `j`. `i` tracks the position of the last unique element, while `j` iterates through the array.
- **Logic**: If `nums[j]` is different from `nums[i]`, increment `i` and set `nums[i]` to `nums[j]`.
- **Result**: The function returns the length of the array with unique elements.

**Example**:
```go
func removeDuplicates(nums []int) int {
    if len(nums) == 0 {
        return 0
    }

    i := 0
    for j := 1; j < len(nums); j++ {
        if nums[j] != nums[i] {
            i++
            nums[i] = nums[j]
        }
    }
    return i + 1
}
```

### Sliding Window Algorithm

The sliding window technique is used to find the maximum sum of any contiguous subarray of size `K`. The implementation involves:

- **Initial Window**: Calculate the sum of the first `K` elements.
- **Sliding the Window**: Move the window one element to the right, updating the sum by subtracting the element that is left out and adding the new element.
- **Max Sum Update**: Keep track of the maximum sum encountered during the sliding process.

**Example**:
```go
func maxSumSubarray(arr []int, k int) int {
    n := len(arr)
    if n < k {
        fmt.Println("Invalid input: array length is less than k")
        return -1
    }

    maxSum := 0
    for i := 0; i < k; i++ {
        maxSum += arr[i]
    }

    windowSum := maxSum
    for i := k; i < n; i++ {
        windowSum += arr[i] - arr[i-k]
        if windowSum > maxSum {
            maxSum = windowSum
        }
    }

    return maxSum
}
```

### Merge Intervals Algorithm

The merge intervals algorithm is used to merge overlapping intervals in a given collection. This algorithm is essential for solving scheduling problems and optimizing interval-based operations.

#### How the Merge Intervals Function Works:

**1. Problem Definition:**
- **Input**: A collection of intervals represented as `[][]int` where each sub-array represents an interval `[start, end]`
- **Output**: A collection of non-overlapping intervals after merging all overlapping ones
- **Goal**: Consolidate overlapping intervals to minimize the total number of intervals

**2. Algorithm Steps:**

```go
func mergeIntervals(intervals [][]int) [][]int {
    // Step 1: Handle edge cases
    if len(intervals) <= 1 {
        return intervals
    }

    // Step 2: Sort intervals by start time
    sort.Slice(intervals, func(i, j int) bool {
        return intervals[i][0] < intervals[j][0]
    })

    // Step 3: Initialize result with first interval
    result := [][]int{intervals[0]}

    // Step 4: Process each interval
    for i := 1; i < len(intervals); i++ {
        current := intervals[i]
        lastMerged := result[len(result)-1]

        // Step 5: Check for overlap and merge or add
        if current[0] <= lastMerged[1] {
            // Overlap detected - merge intervals
            lastMerged[1] = max(lastMerged[1], current[1])
        } else {
            // No overlap - add as new interval
            result = append(result, current)
        }
    }

    return result
}
```

**3. Key Concepts:**

- **Overlap Condition**: Two intervals `[a,b]` and `[c,d]` overlap if `c <= b` (the start of the second interval is less than or equal to the end of the first)
- **Merge Operation**: When overlapping, combine intervals by taking the minimum start time and maximum end time
- **Sorting Importance**: Sorting by start time ensures we process intervals in chronological order, making the merge process efficient

**4. Step-by-Step Example:**

**Input**: `[[1,3], [2,6], [8,10], [15,18]]`

1. **After Sorting**: `[[1,3], [2,6], [8,10], [15,18]]` (already sorted)
2. **Initialize**: `result = [[1,3]]`
3. **Process [2,6]**: Since `2 <= 3`, overlap detected. Merge to `[1,6]`. `result = [[1,6]]`
4. **Process [8,10]**: Since `8 > 6`, no overlap. Add new interval. `result = [[1,6], [8,10]]`
5. **Process [15,18]**: Since `15 > 10`, no overlap. Add new interval. `result = [[1,6], [8,10], [15,18]]`

**Output**: `[[1,6], [8,10], [15,18]]`

**5. Time and Space Complexity:**
- **Time Complexity**: O(n log n) due to sorting, where n is the number of intervals
- **Space Complexity**: O(n) for the result array in the worst case

**6. Common Use Cases:**
- Meeting room scheduling
- Calendar event consolidation
- Resource allocation optimization
- Timeline merging in data processing

These algorithms are efficient and reduce the time complexity of problems involving arrays, making them valuable tools in data structures and algorithms.
