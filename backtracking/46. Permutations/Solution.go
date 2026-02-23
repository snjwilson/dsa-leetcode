package Solution

func permute(nums []int) [][]int {
	result := [][]int{}
	backtrack(&result, nums, []int{}, make(map[int]bool))
	return result
}

func backtrack(result *[][]int, nums, current []int, used map[int]bool) {
	if len(current) == len(nums) {
		tmp := make([]int, len(current))
		copy(tmp, current)
		*result = append(*result, tmp)
		return
	}

	for i := 0; i < len(nums); i++ {
		if used[nums[i]] {
			continue
		}

		used[nums[i]] = true
		current = append(current, nums[i])
		backtrack(result, nums, current, used)
		used[nums[i]] = false
		current = current[:len(current)-1]
	}
}
