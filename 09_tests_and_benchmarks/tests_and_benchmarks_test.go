package main

import (
	"testing"
)

// function needs to start with Test and gets a pointer of type
// testing.T to run functions on it
func TestAdd(t *testing.T) {
	result := add(1, 2)

	// yep, as straight forward as error handling
	if result != 3 {
		t.Errorf("Expected add(1,2) to be 3, got %d instead", result)
	}
}

// Run benchmarks with `go test -bench=.`

// benchmarks are also included out of the box
// the naming indicates what it is (benchmark or test)
// and the parameter is different
// b.N is a constant which magically is provided by the goruntime
// and runs the test as often as is needed to get good average results
func BenchmarkAdd(b *testing.B) {
	for n := 0; n < b.N; n++ {
		add(10, 20)
	}
}

func BenchmarkAdd2(b *testing.B) {
	for n := 0; n < b.N; n++ {
		add2(10, 20)
	}
}
