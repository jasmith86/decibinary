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

type DeciBinary struct {
	number   int
	response []int
}

// Compute the minimal combination of deci-binary numbers that sum to n.
func SolveDeciBinary(dbnum *DeciBinary) {
	//n := *dbnum
	n := dbnum.number
	if n == 0 { // Make sure 0 can be handled.
		dbnum.response = []int{0}
		//return []int{0}
	} else { // TODO eliminate this else
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
		dbnum.response = allFactors
		//fmt.Println(allFactors, dbnum.number, dbnum.response)
		//return allFactors
	}
}

// fanOutWorker gets job and puts answer in answers channel
func fanOutWorker(jobs <-chan DeciBinary, answers chan<- DeciBinary, done chan<- bool) { // removed unused workerID int
	for j := range jobs {
		//fmt.Println("starting  job", j)
		SolveDeciBinary(&j)
		answers <- j
		//fmt.Println("finishing job", j)
	}
	//fmt.Println(workerID, "done")
	done <- true
}

// fanInPrinter collects finished jobs (answers) and prints them
func fanInPrinter(answers <-chan DeciBinary, done chan<- bool) {
	// @TODO DI for testing Println?
	for solution := range answers {
		fmt.Println(solution.number, "\t", len(solution.response), "steps:", solution.response)
	}
	//fmt.Println("done printing")
	done <- true
}

// Driver
func main() {
	const numWorkers = 4
	args := os.Args[1:]

	jobs := make(chan DeciBinary) // create in/out channels
	answers := make(chan DeciBinary)
	done := make(chan bool)

	for w := 1; w <= numWorkers; w++ { // create workers
		go fanOutWorker(jobs, answers, done)
	}

	go fanInPrinter(answers, done)

	for _, arg := range args {
		n, err := strconv.Atoi(arg)
		if err != nil {
			fmt.Println("ignoring non-integer", err)
			continue
		}
		jobs <- DeciBinary{number: n} // add each valid input as job
		//jobs <- n // add each valid input as job
	}
	close(jobs)
	for i := 1; i <= numWorkers; i++ { // wait for fanOutWorkers to all finish
		<-done
	}
	close(answers)
	<-done // wait til printing is finished
}
