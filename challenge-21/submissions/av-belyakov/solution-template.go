package main

import (
	"fmt"
	"slices"
	"math"
)

func main() {
	// Example sorted array for testing
	arr := []int{1, 3, 5, 7, 9, 11, 13, 15, 17, 19}

	// Test binary search
	target := 7
	index := BinarySearch(arr, target)
	fmt.Printf("BinarySearch: %d found at index %d\n", target, index)

	// Test recursive binary search
	recursiveIndex := BinarySearchRecursive(arr, target, 0, len(arr)-1)
	// Test find insert position
	fmt.Printf("BinarySearchRecursive: %d found at index %d\n", target, recursiveIndex)

	insertTarget := 8
	insertPos := FindInsertPosition(arr, insertTarget)
	fmt.Printf("FindInsertPosition: %d should be inserted at index %d\n", insertTarget, insertPos)
}

// BinarySearch performs a standard binary search to find the target in the sorted array.
// Returns the index of the target if found, or -1 if not found.
func BinarySearch(arr []int, target int) int {
	count := len(arr)
	if count == 0 {
		return -1
	}

	if count == 1 {
		if arr[0] == target {
			return 0
		} else {
			return -1
		}
	}

	slices.Sort(arr)

	num := int(math.Ceil(float64(count) / 2))
	arrOne := arr[:num]
	arrTwo := arr[num:]

    //fmt.Println("arr:", arr)
    //fmt.Println("arrOne:", arrOne)
    //fmt.Println("arrTwo:", arrTwo)

	if arrOne[0] <= target && arrOne[len(arrOne)-1] >= target {
		return searchInSlice(target, arrOne)
	}

	if arrTwo[0] <= target && arrTwo[len(arrTwo)-1] >= target {
	    index := searchInSlice(target, arrTwo)
	    if index == -1 {
	        return index
	    }
	    
	    return len(arrOne) + index
	}

	return -1
}

// BinarySearchRecursive performs binary search using recursion.
// Returns the index of the target if found, or -1 if not found.
func BinarySearchRecursive(arr []int, target int, left int, right int) int {
	if left > right {
		return -1
	}

	count := len(arr)
	if count == 0 {
		return -1
	}

	if count == 1 {
		if arr[0] == target {
			return 0
		} else {
			return -1
		}
	}

    mid := left + (right - left) / 2
	if arr[mid] == target {
		return mid
	}

	if arr[mid] > target {
		return BinarySearchRecursive(arr, target, left, mid-1)
	}

	if arr[mid] < target {
		return BinarySearchRecursive(arr, target, mid+1, right)
	}

	return -1
}

// FindInsertPosition returns the index where the target should be inserted
// to maintain the sorted order of the array.
func FindInsertPosition(arr []int, target int) int {
    left, right := 0, len(arr)
    
    for left < right {
        mid := left + (right-left)/2
        if arr[mid] < target {
            left = mid + 1
        } else {
            right = mid
        }
    }
    
    return left
}

func searchInSlice[T comparable](elem T, l []T) int {
	for k, v := range l {
		if elem == v {
            return k
		}
	}

    return -1
}
