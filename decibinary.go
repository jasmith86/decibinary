// Given an input integer, find the smallest combination of deci-binary/pseudo-binary numbers
// that sum to that input. A deci-binary/pseudo-binary number is defined as any base-10
// number whose digits are composed of only 1 or 0. For example, 10, 11, 110, 101 are all
// deci-binary/pseudo-binary numbers.
// Working as of go 1.12.5
package main

import (
	"fmt"
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

// Reverse a slice. Takes point to slice as arg.
func reverse(n *[]int) {
	num := *n
	for i, j := 0, len(num)-1; i < j; i, j = i+1, j-1 { // reverse the slice
		num[i], num[j] = num[j], num[i]
	}
}

// Get a slice that corresponds to the digits of integer n.
// For example, sliceDigits(1234) will return []int{1, 2, 3, 4}
func sliceDigits(n int) []int {
	var num []int
	for n > 0 {
		num = append(num, n%10)
		n /= 10
	}
	reverse(&num)
	return num //reverse(num)
}

// Get integer from slice of digits. Eg. []int{1,2,3} -> 123
func unsliceDigits(n []int) int {
	rv := 0
	for _, v := range n {
		rv = (rv * 10) + v // shift << 1 and add digit
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
	var allFactors []int
	for n > 0 {
		temp := n
		curFactor := 0
		digit := 1
		for temp > 0 { // build current factor
			rem := temp % 10
			temp /= 10
			if rem != 0 {
				curFactor += digit
			}
			digit *= 10
		}
		allFactors = append(allFactors, curFactor)
		n -= curFactor
	}
	return allFactors
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
