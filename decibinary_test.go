// Tests are a good thing. Use them.
package main

import (
	"testing"
)

func TestSolveDeciBinary(t *testing.T) {
	tests := []struct {
		name  string
		input DeciBinary
		want  []int
	}{
		{name: "1", input: DeciBinary{number: 221, response: nil}, want: []int{111, 110}},
		{name: "2", input: DeciBinary{number: 214, response: nil}, want: []int{111, 101, 1, 1}},
		{name: "0", input: DeciBinary{number: 0, response: nil}, want: []int{0}},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			SolveDeciBinary(&test.input)
			if notEqual(test.want, test.input.response) {
				t.Errorf("wanted %+v inputA %+v", test.want, test.input.response)
			}
		})
	}
}

// TestFanOutWorker assures the worker puts the correct answer into
// the answers channel.
func TestFanOutWorker(t *testing.T) {
	tests := []struct {
		name  string
		input DeciBinary
		want  []int
	}{
		{name: "1", input: DeciBinary{number: 221}, want: []int{111, 110}},
		{name: "2", input: DeciBinary{number: 214}, want: []int{111, 101, 1, 1}},
		{name: "0", input: DeciBinary{number: 0}, want: []int{0}},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			jobs := make(chan DeciBinary, 10)
			jobs <- test.input
			answers := make(chan DeciBinary, 10)
			done := make(chan bool, 10)
			go fanOutWorker(jobs, answers, done)
			response := <-answers
			close(jobs)
			if <-done != true {
				t.Errorf("wanted done to be true")
			}
			close(answers)
			close(done)
			if notEqual(test.want, response.response) {
				t.Errorf("wanted %+v inputA %+v", test.want, response)
			}
		})
	}
}

// TestFanInWorker currently only tests that fanInWorker sends true to done,
// it does not test what is printed.
func TestFanInWorker(t *testing.T) {
	// @TODO DI for testing Println?
	tests := []struct {
		name  string
		input DeciBinary
		want  bool
	}{
		{name: "1", input: DeciBinary{number: 221}, want: true},
		{name: "2", input: DeciBinary{number: 214}, want: true},
		{name: "0", input: DeciBinary{number: 0}, want: true},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			jobs := make(chan DeciBinary, 10) // TODO simplify test procedure
			jobs <- test.input
			answers := make(chan DeciBinary, 10)
			done := make(chan bool, 10)
			go fanInPrinter(answers, done)
			go fanOutWorker(jobs, answers, done)
			close(jobs)
			response := <-done
			if response != true {
				t.Errorf("wanted done to be true for fanOut")
			}
			close(answers)
			response = <-done
			if response != true {
				t.Errorf("wanted done to be true for fanIn")
			}
			close(done)
		})
	}
}

// Let's add some benchmarking
// Source: https://dave.cheney.net/2013/06/30/how-to-write-benchmarks-in-go

// Benchmark the SolveDeciBinary function
func benchmarkSolveDeciBinary(num int, b *testing.B) {
	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		SolveDeciBinary(&DeciBinary{number: num})
	}
}

func BenchmarkSolveDeciBinary1(b *testing.B)    { benchmarkSolveDeciBinary(1, b) }
func BenchmarkSolveDeciBinary2342(b *testing.B) { benchmarkSolveDeciBinary(2342, b) }
func BenchmarkSolveDeciBinaryBIG(b *testing.B)  { benchmarkSolveDeciBinary(999847827343453423, b) }
