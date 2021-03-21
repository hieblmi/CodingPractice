package main

import (
	"fmt"
)

// Find lowest available positive slot in an array of n integers ranging from 0 to n-1
// might contain negative numbers, duplicates
// Goal: Runtime complexity O(n), Space complexity O(1)

// Linear runtime -> we can fully traverse the array multiple times for a constant number of times

// Idea:
// 1. traverse - move all negatives and values > size of array(pigeon hole)to the left side of the array by swapping them with positives
// and mark them positive, store the beginning of the subarray of positives as offset from the negatives
// 2. for the positives in the second traversal, take the value as index in the array and mark the entry negative
// 3. traverse the result array a third time, the first positive occurrence is the lowest missing integer

func swap(arr []int, i int, j int) {
	tmp := arr[i]
	arr[i] = arr[j]
	arr[j] = tmp
}

func main() {
	arr := []int{2, 3, -7, 6, 8, 1, -10, 15, 20, -20}
	offset := 0

	// 1.
	for j, i := range arr {
		if i < 0 || i > len(arr) {
			swap(arr, j, offset)
			if i < 0 {
				arr[offset] = -arr[offset]
			}
			offset++
		}
	}

	//2.
	fmt.Println("Complete array: ", arr)
	fmt.Println("Positive subarray: ", arr[offset:])
	for i := offset; i < len(arr); i++ {
		if arr[i] > 0 {
			arr[arr[i]-1] = -abs(arr[arr[i]-1])
		} else {
			arr[abs(arr[i])-1] = -abs(abs(arr[i]) - 1)
		}
	}
	fmt.Println("After applying negatives: ", arr)

	// 3.
	for j, i := range arr {
		if i > 0 {
			fmt.Printf("Lowest available slow is: %d\n", j+1)
			return
		}
	}
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}
