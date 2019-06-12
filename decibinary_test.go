// Tests are a good thing. Use them.
package main

import (
	"testing"
)

func TestSolveDeciBinary(t *testing.T) {
	tests := []struct {
		name  string
		input int
		want  []int
	}{
		{name: "1", input: 221, want: []int{111, 110}},
		{name: "2", input: 214, want: []int{111, 101, 1, 1}},
		{name: "0", input: 0, want: []int{0}},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			response := SolveDeciBinary(test.input)
			if notEqual(test.want, response) {
				t.Errorf("wanted %+v inputA %+v", test.want, response)
			}
		})
	}
}

// Let's add some benchmarking
// Source: https://dave.cheney.net/2013/06/30/how-to-write-benchmarks-in-go

// Benchmark the SolveDeciBinary function
func benchmarkSolveDeciBinary(num int, b *testing.B) {
	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		SolveDeciBinary(num)
	}
}

func BenchmarkSolveDeciBinary1(b *testing.B)    { benchmarkSolveDeciBinary(1, b) }
func BenchmarkSolveDeciBinary2342(b *testing.B) { benchmarkSolveDeciBinary(2342, b) }
func BenchmarkSolveDeciBinaryBIG(b *testing.B)  { benchmarkSolveDeciBinary(999847827343453423, b) }
