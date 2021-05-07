package services

import (
  "api/utils/sort"
)

// By calling an external package (sort), the corresponding
// test is considered an "integration test"
func Sort(elements []int) {
  sort.BubbleSort(elements)
}
