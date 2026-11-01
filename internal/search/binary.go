package search

import (
	"errors"
)

func Binary(items []int, item int) (int, error) {
	if len(items) == 0 {
		return 0, errors.New("search: item not found")
	}

	low := 0
	high := len(items) - 1

	for {
		mid := (low + high) / 2
		if items[mid] == item {
			return mid, nil
		}
		if low >= high {
			return 0, errors.New("search: item not found")
		}
		if items[mid] < item {
			low = mid + 1
		}
		if items[mid] > item {
			high = mid - 1
		}
	}
}
