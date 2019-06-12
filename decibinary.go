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

// fanOutWorker gets job and puts answer in answers channel
func fanOutWorker(workerID int, jobs chan int, answers chan []int, done chan bool) {
	for j := range jobs {
		fmt.Println(workerID, "starting  job", j)
		answers <- SolveDeciBinary(j)
		fmt.Println(workerID, "finishing job", j)
	}
	fmt.Println(workerID, "done")
	done <- true
}

//
func fanInPrinter(answers chan []int, done chan bool) {
	for solution := range answers {
		fmt.Println("\t", len(solution), "steps:", solution)
	}
	fmt.Println("done printing")
	done <- true
}

// Driver
func main() {
	// @TODO handle more inputs right now must be less than buffer size of jobs & answers
	const numWorkers = 4
	args := os.Args[1:]

	jobs := make(chan int) // create in/out channels
	answers := make(chan []int)
	done := make(chan bool)

	for w := 1; w <= numWorkers; w++ { // create workers
		go fanOutWorker(w, jobs, answers, done)
	}

	go fanInPrinter(answers, done)

	for _, arg := range args {
		n, err := strconv.Atoi(arg)
		if err != nil {
			fmt.Println("ignoring non-integer", err)
			//os.Exit(1)
		} else {
			jobs <- n // add each input as job
			//time.Sleep(10 * time.Millisecond)
		}
	}
	close(jobs)
	for i := 1; i <= numWorkers; i++ { // wait for fanOutWorkers to all finish
		<-done
	}
	close(answers)
	<-done // wait til printing is finished
}
