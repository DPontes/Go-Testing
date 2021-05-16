package sorting

import (
  "testing"
)

func TestBubbleSortOrderIncre(t *testing.T) {
  // Init
  elements := []int{9,7,5,3,1,2,4,6,8,0}

  // Execution
  BubbleSort(elements)

  // Validation
  if elements[0] != 0 {
    t.Error( "First element should be 0" )
  }
  if elements[len(elements) - 1] != 9 {
    t.Error( "Last element should be 9" )
  }
}

func TestSortOrderIncr(t *testing.T) {
  elements := []int{9,7,5,3,1,2,4,6,8,0}
  Sort(elements)
  if elements[0] != 0 {
    t.Error( "First element should be 0" )
  }
  if elements[len(elements) - 1] != 9 {
    t.Error( "Last element should be 9" )
  }
}

// functions that start with "Benchmark*" are considered
// by the Go compiler as special "Benchmark" functions,
// much like "Test*" functions are functions for testing
//
// To run: go test -test.bench=.
//         go test -test.bench=BenchmarkBubbleSort
func BenchmarkBubbleSort(b *testing.B) {
  elements := []int{9,7,5,3,1,2,4,6,8,0}

  // The function to be benchmarked will be inside the for-loop
  for i := 0; i <b.N; i++ {
    BubbleSort(elements)
  }
}

func BenchmarkSort(b *testing.B) {
  elements := []int{9,7,5,3,1,2,4,6,8,0}

  // The function to be benchmarked will be inside the for-loop
  for i := 0; i <b.N; i++ {
    Sort(elements)
  }
}
