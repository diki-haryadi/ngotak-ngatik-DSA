package main

import (
	"fmt"
	"strconv"
)

func fizzBuzz(n int) []string {
	var out = make([]string, n)
	for index := range out {
		switch (index + 1) % 15 {
		case 0:
			out[index] = "FizzBuzz"
		case 3, 6, 9, 12:
			out[index] = "Fizz"
		case 5, 10:
			out[index] = "Buzz"
		default:
			out[index] = strconv.Itoa(index + 1)
		}
	}
	return out
}

func main() {
	// Test cases
	testCases := []int{15, 20, 30}

	for _, n := range testCases {
		fmt.Printf("\nFizzBuzz for n = %d:\n", n)
		result := fizzBuzz(n)

		// Print results in a formatted way
		for i, val := range result {
			// Add padding to align output
			padding := ""
			if len(val) < 8 {
				padding = "\t"
			}

			fmt.Printf("%d: %s%s", i+1, val, padding)

			// New line every 5 items for readability
			if (i+1)%5 == 0 {
				fmt.Println()
			}
		}
		fmt.Println("\n-------------------")
	}
}
