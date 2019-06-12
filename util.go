package main

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

// notEqual tells whether a and b contain the same elements.
// A nil argument is equivalent to an empty slice.
// Source for notEqual: https://yourbasic.org/golang/compare-slices/
func notEqual(a, b []int) bool {
	if len(a) != len(b) {
		return true
	}
	for i, v := range a {
		if v != b[i] {
			return true
		}
	}
	return false
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

//func main() {
//
//}
