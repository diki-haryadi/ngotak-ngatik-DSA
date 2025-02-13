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

// Binary search untuk mencari target dalam array yang sudah terurut
func IterativeBinarySearch2(arr []int, target int) int {
	// arr [1,3,5,7,9,11,13,14,15]
	// target 7
	left, right := 0, len(arr)-1 // Inisialisasi pointer kiri di indeks 0 dan kanan di indeks terakhir

	for left <= right { // Loop selama area pencarian masih valid
		mid := (left + right) / 2 // Hitung indeks tengah
		midValue := arr[mid]      // Ambil nilai di indeks tengah

		if midValue == target {
			return mid // Target ditemukan, kembalikan indeksnya
		} else if midValue < target {
			left = mid + 1 // BUG: Seharusnya left = mid + 1 untuk mencari di bagian kanan
			// karena target lebih besar dari nilai tengah
		} else {
			right = mid - 1
		}
		// BUG: Tidak ada kondisi untuk menangani midValue > target
	}
	return -1 // Target tidak ditemukan
}

func main() {
	sortedsearchList := []int{1, 3, 5, 7, 9, 11, 13, 14, 15}
	target := 7

	if index := IterativeBinarySearch2(sortedsearchList, target); index != -1 {
		fmt.Printf("Element %d found at index %d\n", target, index)
	} else {
		fmt.Printf("Element %d not found in the list\n", target)
	}
}
