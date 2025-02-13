package main

import "fmt"

func findKthPositive(arr []int, k int) int {
	// arr = [1,3,5,7], k = 3
	// Mencari bilangan yang hilang ke-k

	j := 1 // j sebagai counter untuk mengecek bilangan 1,2,3,4,5...

	for i := 0; i < len(arr); j++ {
		cursor := arr[i]
		// i untuk iterasi array, j untuk mengecek bilangan
		// Iterasi akan berjalan sampai i < 4 (panjang array)

		if cursor == j {
			// Jika nilai arr[i] sama dengan j
			// Contoh: ketika i=0, j=1: arr[0]=1 == 1
			i++ // Lanjut ke elemen array berikutnya
		} else {
			// Jika nilai arr[i] tidak sama dengan j, berarti j adalah bilangan yang hilang
			// Contoh: ketika i=0, j=2: arr[0]=1 != 2, berarti 2 hilang
			k-- // Kurangi k karena menemukan 1 bilangan hilang
		}

		if k == 0 {
			// Jika sudah menemukan k bilangan yang hilang
			return j // Kembalikan bilangan hilang ke-k
		}
	}

	// Jika loop selesai tapi k belum 0,
	// berarti bilangan hilang ke-k ada setelah array
	return j + k - 1
}

func findKthPositive2(arr []int, k int) int {
	j := 1
	for i := 0; i < len(arr); j++ {
		cursor := arr[i]
		if cursor == j {
			i++
		} else {
			k--
		}

		if k == 0 {
			return j
		}
	}
	return j + k - 1
}

func main() {
	// Test cases
	testCases := []struct {
		arr []int
		k   int
	}{
		{[]int{1, 2, 3, 4}, 2},     // Expected: 6
		{[]int{2, 3, 4, 7, 11}, 5}, // Expected: 9
		{[]int{1, 3, 5, 7}, 3},     // Expected: 6
		{[]int{5, 6, 8, 9, 10}, 3}, // Expected: 3
	}

	// Run test cases
	for i, tc := range testCases {
		result := findKthPositive(tc.arr, tc.k)
		result2 := findKthPositive2(tc.arr, tc.k)
		fmt.Println(result, result2)

		fmt.Printf("Test Case %d:\n", i+1)
		fmt.Printf("Array: %v\n", tc.arr)
		fmt.Printf("k: %d\n", tc.k)
		fmt.Printf("Result: %d\n", result)
		fmt.Println("-------------------")
	}
}
