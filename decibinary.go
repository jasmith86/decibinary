// Given an input integer, find the smallest combination of deci-binary/pseudo-binary numbers
// that sum to that input. A deci-binary/pseudo-binary number is defined as any base-10
// number whose digits are composed of only 1 or 0. For example, 10, 11, 110, 101 are all
// deci-binary/pseudo-binary numbers.
// Working as of go 1.12.5
package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
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
	for i := count - 1; i >= 0; i-- { // get least significant digit first, then next...
		num = append(num, n%10)
		n /= 10
	}
	for i, j := 0, count-1; i < j; i, j = i+1, j-1 { // reverse the slice
		num[i], num[j] = num[j], num[i]
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
func SolveDeciBinary(n int) []int {
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
	args := os.Args[1:]
	for _, arg := range args {
		n, err := strconv.Atoi(arg)
		if err != nil {
			fmt.Println("not an integer", err)
			os.Exit(1)
		} else {
			fmt.Println("Input: ", n)
			solution := SolveDeciBinary(n)
			fmt.Println("\t", len(solution), "steps:", solution)
		}
	}
}
