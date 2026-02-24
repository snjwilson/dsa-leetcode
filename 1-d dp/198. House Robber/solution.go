package main

func rob(nums []int) int {
	n := len(nums)
	if n == 1 {
		return nums[0]
	}
	if n == 2 {
		return max(nums[0], nums[1])
	}
	nums[2] = nums[0] + nums[2]
	if n == 3 {
		return max(nums[1], nums[2])
	}
	for i := 3; i < n; i++ {
		nums[i] += max(nums[i-2], nums[i-3])
	}
	return max(nums[n-1], nums[n-2])
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
