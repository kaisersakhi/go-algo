package sorts

import (
    "go-algo/algorithms/sorts"
    "go-algo/shared"
    "testing"
)

func TestBubbleSort(t *testing.T) {
    tests := []struct {
        input    []shared.Number
        expected []shared.Number
    }{
        {
            input:    shared.NumberArray([]int{5, 1, 4, 2, 8}),
            expected: shared.NumberArray([]int{1, 2, 4, 5, 8}),
        },
        {
            input:    shared.NumberArray([]int{3, 0, 2, 5, -1, 4, 1}),
            expected: shared.NumberArray([]int{-1, 0, 1, 2, 3, 4, 5}),
        },
        {
            input:    shared.NumberArray([]int{5, 4, 3, 2, 1}),
            expected: shared.NumberArray([]int{1, 2, 3, 4, 5}),
        },
        {
            input:    shared.NumberArray([]int{1, 2, 3, 4, 5}),
            expected: shared.NumberArray([]int{1, 2, 3, 4, 5}),
        },
        {
            input:    shared.NumberArray([]int{}),
            expected: shared.NumberArray([]int{}),
        },
    }

    for _, test := range tests {
        result := sorts.BubbleSort(test.input)
        for i := range result {
            if result[i].Value != test.expected[i].Value {
                t.Errorf("Expected %v, but got %v", test.expected, result)
            }
        }
    }
}
