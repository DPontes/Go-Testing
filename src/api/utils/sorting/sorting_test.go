package sorting

// To install the "assert" library below, run:
// `go get github.com/stretchr/testify/assert`
import (
  "time"
  "testing"
  "github.com/stretchr/testify/assert"
)

func ExecuteBubbleSortWithTimeout(elements []int, t *testing.T) {
  timeoutChan := make(chan bool, 1)

  /*
    Defers the execution of a function until the surrounding function returns.
    The deferred call's arguments are evaluated immediately, but the function
    call is not executed until the surrounding function returns.
  */
  defer close(timeoutChan)

  go func() {
    BubbleSort(elements)
    timeoutChan <- false
  }()

  go func() {
    time.Sleep(500 * time.Millisecond)
    timeoutChan <- true
  }()

  // waits until the first coroutine returns a timeoutChan
  // If timeoutChan is "true" (returned by the time.Sleep coroutine)
  // the assert below will fail and the output will be printed out
  if <- timeoutChan {
    assert.Fail(t, "BubbleSort took more than 500ms")
    return
  }
}

func TestBubbleSortOrderIncre(t *testing.T) {
  elements := GetElements(10)

  assert.NotNil(t, elements)
  assert.EqualValues(t, 10, len(elements))
  assert.EqualValues(t, 9, elements[0], "First element should be 9")
  assert.EqualValues(t, 0, elements[len(elements) - 1], "Last element should be 0")

  ExecuteBubbleSortWithTimeout(elements, t)

  assert.NotNil(t, elements)
  assert.EqualValues(t, 10, len(elements))
  assert.EqualValues(t, 0, elements[0], "First element should be 0")
  assert.EqualValues(t, 9, elements[len(elements) - 1], "Last element should be 9")

}

func TestSortOrderIncr(t *testing.T) {
  elements := GetElements(10)
  Sort(elements)

  assert.EqualValues(t, 0, elements[0], "First element should be 0")
  assert.EqualValues(t, 9, elements[len(elements) - 1], "Last element should be 9")
}

// functions that start with "Benchmark*" are considered
// by the Go compiler as special "Benchmark" functions,
// much like "Test*" functions are functions for testing
//
// To run: go test -test.bench=.
//         go test -test.bench=BenchmarkBubbleSort
func BenchmarkBubbleSort(b *testing.B) {
  elements := GetElements(100000)

  // The function to be benchmarked will be inside the for-loop
  for i := 0; i <b.N; i++ {
    BubbleSort(elements)
  }
}

func BenchmarkSorter(b *testing.B) {
  elements := GetElements(100000)

  // The function to be benchmarked will be inside the for-loop
  for i := 0; i <b.N; i++ {
    Sort(elements)
  }
}
