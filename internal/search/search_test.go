package search

import "testing"

func TestBinary(t *testing.T) {
	type TestResult struct {
		value int
		error bool
	}
	type Test struct {
		items  []int
		item   int
		result TestResult
	}

	tests := []Test{
		{[]int{1, 3, 7, 8, 9}, 7, TestResult{2, false}},
		{[]int{1, 3, 7, 8, 9}, 1, TestResult{0, false}},
		{[]int{1, 3, 7, 8, 9}, 4, TestResult{0, true}},
		{[]int{}, 4, TestResult{0, true}},
		{[]int{1, 3}, 1, TestResult{0, false}},
		{[]int{1, 3}, 2, TestResult{0, true}},
		{[]int{1, 3}, 3, TestResult{1, false}},
		{[]int{1, 3}, -3, TestResult{0, true}},
		{[]int{1, 3}, 100, TestResult{0, true}},
	}

	for _, test := range tests {
		value, ok := Binary(test.items, test.item)
		if ok == nil && test.result.error {
			t.Errorf("Binary(%v, %v) failed.", test.items, test.item)
			continue
		}
		if value != test.result.value {
			t.Errorf("Binary(%v, %v) failed.", test.items, test.item)
		}
	}
}
