package services

import (
  "api/utils/sorting"
)

// By calling an external package (sorting), the corresponding
// test is considered an "integration test"
func Sort(elements []int) {
  sorting.BubbleSort(elements)
}
