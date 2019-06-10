package main

import (
	"fmt"
	"math"
)

func countDigits(n int) int {
	// report 0 for negative numbers
	count := 0
	if n >= 0 {
		for n != 0 {
			n /= 10
			count += 1
		}
	}
	return count
}

func sliceDigits(n int) []int {
	count := countDigits(n)
	var num []int
	for i := count - 1; i >= 0; i-- {
		digit := math.Mod(float64(n)/math.Pow10(i), 10)
		num = append(num, int(digit))
	}
	return num
}

func unsliceDigits(n []int) int {
	rv := 0
	for i, v := range n {
		rv += int(float64(v) * math.Pow10(len(n)-1-i))
	}
	return rv
}

func sum(n []int) int {
	total := 0
	for _, val := range n {
		total += val
	}
	return total
}

func solve(n int) []int {
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

func main() {
	n := 432
	fmt.Println("slice", sliceDigits(n))
	fmt.Println("unslice", unsliceDigits(sliceDigits(n)))
	solution := solve(n)
	fmt.Println("solve", solution)
	fmt.Println("steps", len(solution))
	fmt.Println("sum", sum(solution))

}
