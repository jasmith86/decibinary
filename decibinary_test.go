package main

import (
	"testing"
)

// notEqual tells whether a and b contain the same elements.
// A nil argument is equivalent to an empty slice.
// Source https://yourbasic.org/golang/compare-slices/
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

func TestSolve(t *testing.T) {
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
			response := solve(test.input)
			if notEqual(test.want, response) {
				t.Errorf("wanted %+v input %+v", test.want, response)
			}
		})
	}
}
