package main

import (
	"fmt"
	"slices"
)

func iterativeBinarySearch(arr []int, target int) int {
	left, right := 0, len(arr)-1

	for left <= right {
		mid := (left + right) / 2
		midValue := arr[mid]
		if midValue == target {
			return mid
		} else if midValue < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return -1
}

func recursiveBinarySearch(arr []int, target int, left int, right int) int {
	if left > right {
		// target is not found
		return -1
	}

	mid := left + (right-left)/2

	if arr[mid] == target {
		// target found, return the element
		return mid
	} else if arr[mid] < target {
		return recursiveBinarySearch(arr, target, mid+1, right)
	} else {
		return recursiveBinarySearch(arr, target, left, mid-1)
	}
}

func slice(arr []int, target int) int {
	result := 0
	if index, ok := slices.BinarySearch(arr, target); ok {
		result = arr[index]
		fmt.Printf("Element %d found at index %d\n", target, index)
	} else {
		fmt.Printf("Element %d not found in the list\n", target)
	}
	return result
}

func main() {
	sortedsearchList := []int{1, 3, 5, 7, 9, 11, 13, 14, 15}
	target := 7

	if index := iterativeBinarySearch(sortedsearchList, target); index != -1 {
		fmt.Printf("Element %d found at index %d\n", target, index)
	} else {
		fmt.Printf("Element %d not found in the list\n", target)
	}
}
