package services

import (
  "testing"
)

func TestSortOrderDesc(t *testing.T) {
  elements := []int{9,7,5,3,1,2,4,6,8,0}
  Sort(elements)
  if elements[0] != 9 {
    t.Error( "First element should be 9" )
  }
  if elements[len(elements) - 1] != 0 {
    t.Error( "Last element should be 0" )
  }
}
