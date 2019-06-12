package main

import "testing"

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
				t.Errorf("wanted %+v inputA %+v", test.want, response)
			}
		})
	}
}

func TestSliceDigits(t *testing.T) {
	want := []int{1, 2, 3, 4, 0}
	got := sliceDigits(12340)
	if notEqual(want, got) {
		t.Errorf("wanted %+v inputA %+v", want, got)
	}
}

func TestUnsliceDigits(t *testing.T) {
	want := 12340
	got := unsliceDigits([]int{1, 2, 3, 4, 0})
	if want != got {
		t.Errorf("wanted %+v inputA %+v", want, got)
	}
}

func TestSum(t *testing.T) {
	want := 112
	got := sum([]int{100, 10, 1, 1})
	if want != got {
		t.Errorf("wanted %+v inputA %+v", want, got)
	}
}

func TestNotEqual(t *testing.T) {
	tests := []struct {
		name   string
		inputA []int
		inputB []int
		want   bool
	}{
		{name: "are equal", inputA: []int{111, 110}, inputB: []int{111, 110}, want: false},
		{name: "diff length", inputA: []int{110}, inputB: []int{111, 110}, want: true},
		{name: "diff values", inputA: []int{111, 123}, inputB: []int{111, 110}, want: true},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			response := notEqual(test.inputA, test.inputB)
			if response != test.want {
				t.Errorf("wanted %+v inputA %+v", test.want, response)
			}
		})
	}
}

// Benchmarking

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
