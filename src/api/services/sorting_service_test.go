package services

// To install the "assert" library below, run:
// `go get github.com/stretchr/testify/assert`
import (
  "api/utils/sorting"
  "testing"
  "github.com/stretchr/testify/assert"
)

func TestSortOrderIncr(t *testing.T) {
  elements := sorting.GetElements(10)

  Sort(elements)

  assert.NotNil(t, elements)

  assert.EqualValues(t, 0, elements[0], "First element should be 0")
  assert.EqualValues(t, 9, elements[len(elements) - 1], "Last element should be 9")
}

func TestSortMoreThan10K(t *testing.T) {
  elements := sorting.GetElements(10001)

  Sort(elements)

  assert.EqualValues(t, 0, elements[0], "First element should be 0")
  assert.EqualValues(t, 10000, elements[len(elements) - 1], "Last element should be 10000")
}

func BenchmarkSort10K(b *testing.B) {
  elements := sorting.GetElements(10000)

  // The function to be benchmarked will be inside the for-loop
  for i := 0; i <b.N; i++ {
    Sort(elements)
  }
}

func BenchmarkSort100K(b *testing.B) {
  elements := sorting.GetElements(100000)

  // The function to be benchmarked will be inside the for-loop
  for i := 0; i <b.N; i++ {
    Sort(elements)
  }
}
