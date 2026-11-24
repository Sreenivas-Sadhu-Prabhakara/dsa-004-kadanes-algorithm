package main

import "fmt"

// MaxSubarraySum returns the largest sum over all non-empty contiguous
// subarrays of nums (Kadane's algorithm). Returns 0 for an empty slice.
func MaxSubarraySum(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	best, cur := nums[0], nums[0]
	for _, v := range nums[1:] {
		if cur < 0 {
			cur = 0
		}
		cur += v
		if cur > best {
			best = cur
		}
	}
	return best
}

func main() {
	fmt.Println(MaxSubarraySum([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
}
