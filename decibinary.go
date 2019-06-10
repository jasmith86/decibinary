// Given an input integer, find the smallest combination of deci-binary/pseudo-binary numbers
// that sum to that input. A deci-binary/pseudo-binary number is defined as any base-10
// number whose digits are composed of only 1 or 0. For example, 10, 11, 110, 101 are all
// deci-binary/pseudo-binary numbers.
package main

import (
	"fmt"
	"math"
)

// Count the number of digits in integer n.
// 		countDigits(123) -> 3
// 		countDigits(-12) -> 2
// 		countDigits(0)   -> 1
func countDigits(n int) int {
	count := 0
	if n == 0 {
		return 1
	} else if n < 0 {
		n *= -1
	}
	for n != 0 {
		n /= 10
		count += 1
	}

	return count
}

// Get a slice that corresponds to the digits of integer n.
// For example, sliceDigits(1234) will return []int{1, 2, 3, 4}
func sliceDigits(n int) []int {
	count := countDigits(n)
	var num []int
	for i := count - 1; i >= 0; i-- { // get most significant digit first, then next...
		digit := math.Mod(float64(n)/math.Pow10(i), 10)
		num = append(num, int(digit))
	}
	return num
}

// Get integer from slice of digits. Eg. []int{1,2,3} -> 123
func unsliceDigits(n []int) int {
	rv := 0
	for i, v := range n {
		rv += int(float64(v) * math.Pow10(len(n)-1-i))
	}
	return rv
}

// Simple sum for int slice.
func sum(n []int) int {
	total := 0
	for _, val := range n {
		total += val
	}
	return total
}

// Compute the minimal combination of deci-binary numbers that sum to n.
func solve(n int) []int {
	if n == 0 { // Make sure 0 can be handled.
		return []int{0}
	}
	var dbnums []int
	for n > 0 {
		nslice := sliceDigits(n)
		var dbnum []int
		for ind, digit := range nslice {
			if digit > 0 {
				dbnum = append(dbnum, 1)
				nslice[ind] -= 1
			} else {
				dbnum = append(dbnum, 0)
			}
		}
		dbnums = append(dbnums, unsliceDigits(dbnum))
		n = unsliceDigits(nslice)
	}
	return dbnums
}

// Driver
func main() {
	n := 432
	fmt.Println("slice", sliceDigits(n))
	fmt.Println("unslice", unsliceDigits(sliceDigits(n)))
	solution := solve(n)
	fmt.Println("solve", solution)
	fmt.Println("steps", len(solution))
	fmt.Println("sum", sum(solution))

}
