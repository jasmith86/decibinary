// Tests are a good thing. Use them.

package main

import (
	"testing"
)

// notEqual tells whether a and b contain the same elements.
// A nil argument is equivalent to an empty slice.
// Source for notEqual: https://yourbasic.org/golang/compare-slices/
func notEqual(a, b []int) bool {
	if len(a) != len(b) {
		//fmt.Printf("diff len %+v %+v\n", len(a), len(b))
		return true
	}
	for i, v := range a {
		if v != b[i] {
			//fmt.Printf("diff val %+v %+v\n", v, b[i])
			return true
		}
	}
	//fmt.Println("all ok")
	return false
}

func TestCountDigits(t *testing.T) {
	tests := []struct {
		name  string
		input int
		want  int
	}{
		{name: "positive", input: 221, want: 3},
		{name: "negative", input: -4414, want: 4},
		{name: "zero", input: 0, want: 1},
		{name: "single", input: 2, want: 1},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			response := countDigits(test.input)
			if test.want != response {
				t.Errorf("wanted %+v input %+v", test.want, response)
			}
		})
	}
}

func TestSliceDigits(t *testing.T) {
	want := []int{1, 2, 3, 4, 0}
	got := sliceDigits(12340)
	if notEqual(want, got) {
		t.Errorf("wanted %+v input %+v", want, got)
	}
}

func TestUnsliceDigits(t *testing.T) {
	want := 12340
	got := unsliceDigits([]int{1, 2, 3, 4, 0})
	if want != got {
		t.Errorf("wanted %+v input %+v", want, got)
	}
}

func TestSum(t *testing.T) {
	want := 112
	got := sum([]int{100, 10, 1, 1})
	if want != got {
		t.Errorf("wanted %+v input %+v", want, got)
	}
}

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
				t.Errorf("wanted %+v input %+v", test.want, response)
			}
		})
	}
}

// Let's add some benchmarking
// Source: https://dave.cheney.net/2013/06/30/how-to-write-benchmarks-in-go

// Benchmark the notEqual function
func BenchmarkNotEqual(b *testing.B) {
	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		notEqual(
			[]int{9, 9, 9, 8, 4, 7, 8, 2, 7, 3, 4, 3, 4, 5, 3, 4, 2, 3},
			[]int{9, 9, 9, 8, 4, 7, 8, 2, 7, 3, 4, 3, 4, 5, 3, 4, 2, 3},
		)
	}
}

// Benchmark the countDigits function
func benchmarkTestCountDigits(num int, b *testing.B) {
	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		countDigits(num)
	}
}

//func BenchmarkTestCountDigits1(b *testing.B)    { benchmarkTestCountDigits(1, b) }
//func BenchmarkTestCountDigits12342(b *testing.B) { benchmarkTestCountDigits(2342, b) }
func BenchmarkTestCountDigits1BIG(b *testing.B) { benchmarkTestCountDigits(999847827343453423, b) }

// Benchmark the slice function
func BenchmarkSliceDigits(b *testing.B) {
	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		sliceDigits(999847827343453423)
	}
}

// Benchmark the unslice function
func BenchmarkUnsliceDigits(b *testing.B) {
	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		unsliceDigits([]int{9, 9, 9, 8, 4, 7, 8, 2, 7, 3, 4, 3, 4, 5, 3, 4, 2, 3})
	}
}

// Benchmark the sum function
func BenchmarkSum(b *testing.B) {
	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		sum([]int{9, 9, 9, 8, 4, 7, 8, 2, 7, 3, 4, 3, 4, 5, 3, 4, 2, 3})
	}
}

// Benchmark the SolveDeciBinary function
func benchmarkSolveDeciBinary(num int, b *testing.B) {
	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		SolveDeciBinary(num)
	}
}

//func BenchmarkSolveDeciBinary1(b *testing.B)    { benchmarkSolveDeciBinary(1, b) }
//func BenchmarkSolveDeciBinary2342(b *testing.B) { benchmarkSolveDeciBinary(2342, b) }
func BenchmarkSolveDeciBinaryBIG(b *testing.B) { benchmarkSolveDeciBinary(999847827343453423, b) }
