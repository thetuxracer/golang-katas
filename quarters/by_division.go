package main

import "fmt"

// QuarterOfYear returns the quarter (1-4) for a given month (1-12)
func QuarterOfYear(month int) int {
	fmt.Println("Calculating quarter for month:", month)
	fmt.Println((month - 1) / 3 + 1)
	return (month - 1) / 3 + 1
}

func main() {
	// Test cases
	testCases := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	
	for _, month := range testCases {
		quarter := QuarterOfYear(month)
		fmt.Printf("Month %d -> Quarter %d\n", month, quarter)
	}
}

