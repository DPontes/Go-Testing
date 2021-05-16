package services

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
