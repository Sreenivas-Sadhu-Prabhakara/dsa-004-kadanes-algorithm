package main

import "testing"

func TestMaxSubarraySum(t *testing.T) {
	cases := []struct {
		in   []int
		want int
	}{
		{[]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}, 6},
		{[]int{1}, 1},
		{[]int{5, 4, -1, 7, 8}, 23},
		{[]int{-3, -1, -2}, -1},
		{[]int{0, 0, 0}, 0},
		{[]int{}, 0},
	}
	for _, c := range cases {
		if got := MaxSubarraySum(c.in); got != c.want {
			t.Errorf("MaxSubarraySum(%v) = %d, want %d", c.in, got, c.want)
		}
	}
}
