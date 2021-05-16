package services

import (
  "api/utils/sorting"
  "testing"
)

func TestSortOrderIncr(t *testing.T) {
  elements := sorting.GetElements(10)
  Sort(elements)
  if elements[0] != 0 {
    t.Error( "First element should be 0" )
  }
  if elements[len(elements) - 1] != 9 {
    t.Error( "Last element should be 9" )
  }
}

func TestSortMoreThan10K(t *testing.T) {
  elements := sorting.GetElements(10001)
  Sort(elements)
  if elements[0] != 0 {
    t.Error( "First element should be 0" )
  }
  if elements[len(elements) - 1] != 10000 {
    t.Error( "Last element should be 10000" )
  }
}
