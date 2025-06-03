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

These algorithms are efficient and reduce the time complexity of problems involving arrays, making them valuable tools in data structures and algorithms.
