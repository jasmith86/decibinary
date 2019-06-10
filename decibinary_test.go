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
	want := 5
	got := countDigits(12340)
	if got != want {
		t.Errorf("wanted %d got %d", want, got)
	}
}

func TestSliceDigits(t *testing.T) {
	want := []int{1, 2, 3, 4, 0}
	got := sliceDigits(12340)
	if notEqual(want, got) {
		t.Errorf("wanted %+v got %+v", want, got)
	}
}

func TestUnsliceDigits(t *testing.T) {
	want := 12340
	got := unsliceDigits([]int{1, 2, 3, 4, 0})
	if want != got {
		t.Errorf("wanted %+v got %+v", want, got)
	}
}
func TestSum(t *testing.T) {
	want := 112
	got := sum([]int{100,10,1,1})
	if want != got {
		t.Errorf("wanted %+v got %+v", want, got)
	}
}

func TestSolve(t *testing.T) {
	want := []int{111, 110}
	got := solve(221)
	if notEqual(want, got) {
		t.Errorf("wanted %+v got %+v", want, got)
	}
}
