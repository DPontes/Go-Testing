package sort

import (
  "testing"
)

func TestBubbleSortOrderDesc(t *testing.T) {
  // Init
  elements := []int{9,7,5,3,1,2,4,6,8,0}

  // Execution
  BubbleSort(elements)

  // Validation
  if elements[0] != 9 {
    t.Error( "First element should be 9" )
  }
  if elements[len(elements) - 1] != 0 {
    t.Error( "Last element should be 0" )
  }
}
