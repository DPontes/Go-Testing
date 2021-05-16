package services

import (
  "testing"
)

// returns a splice of ints in descending order
func getElements(n int) []int {
  result := make([]int, n)
  j := 0
  for i := n - 1; i > 0; i-- {
    result[j] = i
    j++
  }
  return result
}

func TestSortOrderIncr(t *testing.T) {
  elements := getElements(10)
  Sort(elements)
  if elements[0] != 0 {
    t.Error( "First element should be 0" )
  }
  if elements[len(elements) - 1] != 9 {
    t.Error( "Last element should be 9" )
  }
}

func TestSortMoreThan10K(t *testing.T) {
  elements := getElements(10001)
  Sort(elements)
  if elements[0] != 0 {
    t.Error( "First element should be 0" )
  }
  if elements[len(elements) - 1] != 10000 {
    t.Error( "Last element should be 10000" )
  }
}
