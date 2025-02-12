package main

import (
	"fmt"
	"math"
)

// isPrime mengecek apakah sebuah angka adalah bilangan prima
func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	if n <= 3 {
		return true
	}
	if n%2 == 0 || n%3 == 0 {
		return false
	}

	// Mengecek faktor sampai akar kuadrat dari n
	for i := 5; float64(i) <= math.Sqrt(float64(n)); i += 6 {
		if n%i == 0 || n%(i+2) == 0 {
			return false
		}
	}
	return true
}

// printPrimes mencetak semua bilangan prima sampai batas tertentu
func printPrimes(limit int) {
	fmt.Printf("Bilangan prima dari 1 sampai %d adalah:\n", limit)
	for i := 2; i <= limit; i++ {
		if isPrime(i) {
			fmt.Printf("%d ", i)
		}
	}
	fmt.Println()
}

func main() {
	// Contoh penggunaan: mencetak bilangan prima sampai 50
	limit := 50
	printPrimes(limit)

	// Contoh penggunaan fungsi isPrime
	number := 17
	if isPrime(number) {
		fmt.Printf("\n%d adalah bilangan prima\n", number)
	} else {
		fmt.Printf("\n%d bukan bilangan prima\n", number)
	}
}
